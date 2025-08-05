package rpc

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/optimism-java/dispute-explorer/pkg/log"
)

// HTTPClientManager manages multiple HTTP RPC endpoints, provides failover functionality
type HTTPClientManager struct {
	urls           []string
	healthStatus   []bool
	currentIndex   int
	maxRetries     int
	retryDelay     time.Duration
	mutex          sync.RWMutex
	healthCheckMux sync.Mutex
}

// NewHTTPClientManager creates a new HTTP client manager
func NewHTTPClientManager(rpcUrls string, maxRetries int, retryDelay time.Duration) (*HTTPClientManager, error) {
	if rpcUrls == "" {
		return nil, fmt.Errorf("rpc urls cannot be empty")
	}

	urls := strings.Split(rpcUrls, ",")
	if len(urls) == 0 {
		return nil, fmt.Errorf("no valid rpc urls found")
	}

	// Remove whitespace
	for i, url := range urls {
		urls[i] = strings.TrimSpace(url)
	}

	healthStatus := make([]bool, len(urls))
	for i := range healthStatus {
		healthStatus[i] = true // Assume all URLs are healthy initially
	}

	hcm := &HTTPClientManager{
		urls:         urls,
		healthStatus: healthStatus,
		currentIndex: 0,
		maxRetries:   maxRetries,
		retryDelay:   retryDelay,
	}

	// Start health check
	go hcm.startHealthCheck()

	return hcm, nil
}

// HTTPPostJSONWithFailover sends JSON POST request with failover
func (hcm *HTTPClientManager) HTTPPostJSONWithFailover(proxyURL, bodyJSON string) ([]byte, error) {
	var lastErr error

	for attempt := 0; attempt < hcm.maxRetries; attempt++ {
		url := hcm.getNextHealthyURL()
		if url == "" {
			lastErr = fmt.Errorf("no healthy RPC URLs available")
			time.Sleep(hcm.retryDelay)
			continue
		}

		result, err := HTTPPostJSON(proxyURL, url, bodyJSON)
		if err == nil {
			return result, nil
		}

		lastErr = err
		log.Warnf("[RPC.HTTPClientManager] Request failed to %s on attempt %d: %v", url, attempt+1, err)

		// Mark current URL as unhealthy
		hcm.markUnhealthy(url)

		// If not the last retry, wait and retry
		if attempt < hcm.maxRetries-1 {
			time.Sleep(hcm.retryDelay)
		}
	}

	return nil, fmt.Errorf("all RPC requests failed after %d attempts: %v", hcm.maxRetries, lastErr)
}

// getNextHealthyURL gets next healthy URL
func (hcm *HTTPClientManager) getNextHealthyURL() string {
	hcm.mutex.Lock()
	defer hcm.mutex.Unlock()

	// Try current index URL first
	if hcm.healthStatus[hcm.currentIndex] {
		return hcm.urls[hcm.currentIndex]
	}

	// Find next healthy URL
	for i := 0; i < len(hcm.urls); i++ {
		index := (hcm.currentIndex + i) % len(hcm.urls)
		if hcm.healthStatus[index] {
			hcm.currentIndex = index
			return hcm.urls[index]
		}
	}

	return ""
}

// markUnhealthy marks specified URL as unhealthy
func (hcm *HTTPClientManager) markUnhealthy(targetURL string) {
	hcm.mutex.Lock()
	defer hcm.mutex.Unlock()

	for i, url := range hcm.urls {
		if url == targetURL {
			hcm.healthStatus[i] = false
			log.Warnf("[RPC.HTTPClientManager] Marked RPC URL %s as unhealthy", url)
			break
		}
	}
}

// startHealthCheck starts health check goroutine
func (hcm *HTTPClientManager) startHealthCheck() {
	ticker := time.NewTicker(30 * time.Second) // Check every 30 seconds
	defer ticker.Stop()

	for range ticker.C {
		hcm.performHealthCheck()
	}
}

// performHealthCheck performs health check
func (hcm *HTTPClientManager) performHealthCheck() {
	hcm.healthCheckMux.Lock()
	defer hcm.healthCheckMux.Unlock()

	for i, url := range hcm.urls {
		// Send simple health check request
		healthCheckBody := `{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}`
		_, err := HTTPPostJSON("", url, healthCheckBody)

		hcm.mutex.Lock()
		if err != nil {
			hcm.healthStatus[i] = false
			log.Debugf("[RPC.HTTPClientManager] Health check failed for %s: %v", url, err)
		} else {
			if !hcm.healthStatus[i] {
				log.Infof("[RPC.HTTPClientManager] RPC URL %s is back online", url)
			}
			hcm.healthStatus[i] = true
		}
		hcm.mutex.Unlock()
	}
}

// GetHealthyURLCount gets count of healthy URLs
func (hcm *HTTPClientManager) GetHealthyURLCount() int {
	hcm.mutex.RLock()
	defer hcm.mutex.RUnlock()

	count := 0
	for _, isHealthy := range hcm.healthStatus {
		if isHealthy {
			count++
		}
	}
	return count
}

// GetStatus gets status of all URLs
func (hcm *HTTPClientManager) GetStatus() map[string]bool {
	hcm.mutex.RLock()
	defer hcm.mutex.RUnlock()

	status := make(map[string]bool)
	for i, url := range hcm.urls {
		status[url] = hcm.healthStatus[i]
	}
	return status
}

func HTTPPostJSON(proxyURL, httpURL, bodyJSON string) ([]byte, error) {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}
	if proxyURL != "" {
		proxy, err := url.Parse(proxyURL)
		if err != nil {
			return nil, err
		}
		netTransport := &http.Transport{
			Proxy:                 http.ProxyURL(proxy),
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second * time.Duration(10),
		}
		httpClient.Transport = netTransport
	}
	b := strings.NewReader(bodyJSON)
	res, err := httpClient.Post(httpURL, "application/json", b)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("StatusCode: %d", res.StatusCode)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
