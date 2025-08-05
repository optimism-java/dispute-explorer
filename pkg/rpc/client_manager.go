package rpc

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/optimism-java/dispute-explorer/pkg/log"
)

// ClientManager manages multiple RPC clients, provides polling, failover and health check functionality
type ClientManager struct {
	clients        []*ethclient.Client
	urls           []string
	healthStatus   []bool
	currentIndex   int
	maxRetries     int
	retryDelay     time.Duration
	mutex          sync.RWMutex
	healthCheckMux sync.Mutex
	httpManager    *HTTPClientManager // Built-in HTTP manager for simple JSON-RPC calls
}

// NewClientManager creates a new client manager
func NewClientManager(rpcUrls string, maxRetries int, retryDelay time.Duration) (*ClientManager, error) {
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

	clients := make([]*ethclient.Client, len(urls))
	healthStatus := make([]bool, len(urls))

	cm := &ClientManager{
		clients:      clients,
		urls:         urls,
		healthStatus: healthStatus,
		currentIndex: 0,
		maxRetries:   maxRetries,
		retryDelay:   retryDelay,
	}

	// Create built-in HTTP manager
	httpManager, err := NewHTTPClientManager(rpcUrls, maxRetries, retryDelay)
	if err != nil {
		log.Errorf("[RPC.ClientManager] Failed to create HTTP manager: %v", err)
	} else {
		cm.httpManager = httpManager
	}

	// Initialize all clients
	for i, url := range urls {
		client, err := ethclient.Dial(url)
		if err != nil {
			log.Errorf("[RPC.ClientManager] Failed to connect to %s: %v", url, err)
			healthStatus[i] = false
		} else {
			clients[i] = client
			healthStatus[i] = true
			log.Infof("[RPC.ClientManager] Successfully connected to %s", url)
		}
	}

	// Start health check
	go cm.startHealthCheck()

	return cm, nil
}

// GetClient gets an available client with failover functionality
func (cm *ClientManager) GetClient() *ethclient.Client {
	cm.mutex.RLock()
	defer cm.mutex.RUnlock()

	// Try current index client first
	if cm.healthStatus[cm.currentIndex] && cm.clients[cm.currentIndex] != nil {
		return cm.clients[cm.currentIndex]
	}

	// If current client is unavailable, find next available client
	for i := 0; i < len(cm.clients); i++ {
		index := (cm.currentIndex + i) % len(cm.clients)
		if cm.healthStatus[index] && cm.clients[index] != nil {
			cm.currentIndex = index
			return cm.clients[index]
		}
	}

	// If all clients are unavailable, return first client (may be nil)
	log.Errorf("[RPC.ClientManager] All RPC clients are unavailable")
	return cm.clients[0]
}

// GetNextClient gets next available client (round robin)
func (cm *ClientManager) GetNextClient() *ethclient.Client {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	// Move to next client
	cm.currentIndex = (cm.currentIndex + 1) % len(cm.clients)

	// Find next available client
	for i := 0; i < len(cm.clients); i++ {
		index := (cm.currentIndex + i) % len(cm.clients)
		if cm.healthStatus[index] && cm.clients[index] != nil {
			cm.currentIndex = index
			return cm.clients[index]
		}
	}

	// If all clients are unavailable, return current index client
	return cm.clients[cm.currentIndex]
}

// GetRandomClient gets a random available client
func (cm *ClientManager) GetRandomClient() *ethclient.Client {
	cm.mutex.RLock()
	defer cm.mutex.RUnlock()

	availableIndices := make([]int, 0)
	for i, isHealthy := range cm.healthStatus {
		if isHealthy && cm.clients[i] != nil {
			availableIndices = append(availableIndices, i)
		}
	}

	if len(availableIndices) == 0 {
		log.Errorf("[RPC.ClientManager] No healthy RPC clients available")
		return cm.clients[0]
	}

	randomIndex := availableIndices[rand.Intn(len(availableIndices))]
	return cm.clients[randomIndex]
}

// ExecuteWithRetry executes operation with retry and failover
func (cm *ClientManager) ExecuteWithRetry(operation func(*ethclient.Client) error) error {
	var lastErr error

	for attempt := 0; attempt < cm.maxRetries; attempt++ {
		client := cm.GetClient()
		if client == nil {
			lastErr = fmt.Errorf("no available RPC client")
			time.Sleep(cm.retryDelay)
			continue
		}

		err := operation(client)
		if err == nil {
			return nil
		}

		lastErr = err
		log.Warnf("[RPC.ClientManager] Operation failed on attempt %d: %v", attempt+1, err)

		// Mark current client as unhealthy
		cm.markUnhealthy(cm.currentIndex)

		// If not the last retry, wait and retry
		if attempt < cm.maxRetries-1 {
			time.Sleep(cm.retryDelay)
		}
	}

	return fmt.Errorf("operation failed after %d attempts: %v", cm.maxRetries, lastErr)
}

// markUnhealthy marks client at specified index as unhealthy
func (cm *ClientManager) markUnhealthy(index int) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	if index >= 0 && index < len(cm.healthStatus) {
		cm.healthStatus[index] = false
		log.Warnf("[RPC.ClientManager] Marked RPC client %s as unhealthy", cm.urls[index])
	}
}

// startHealthCheck starts health check goroutine
func (cm *ClientManager) startHealthCheck() {
	ticker := time.NewTicker(30 * time.Second) // Check every 30 seconds
	defer ticker.Stop()

	for range ticker.C {
		cm.performHealthCheck()
	}
}

// performHealthCheck performs health check
func (cm *ClientManager) performHealthCheck() {
	cm.healthCheckMux.Lock()
	defer cm.healthCheckMux.Unlock()

	for i, client := range cm.clients {
		if client == nil {
			// Try to reconnect
			newClient, err := ethclient.Dial(cm.urls[i])
			if err != nil {
				log.Debugf("[RPC.ClientManager] Health check failed for %s: %v", cm.urls[i], err)
				cm.mutex.Lock()
				cm.healthStatus[i] = false
				cm.mutex.Unlock()
				continue
			}

			// Update client
			cm.mutex.Lock()
			cm.clients[i] = newClient
			cm.mutex.Unlock()
			client = newClient
		}

		// Check if client is healthy
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		_, err := client.BlockNumber(ctx)
		cancel()

		cm.mutex.Lock()
		if err != nil {
			cm.healthStatus[i] = false
			log.Debugf("[RPC.ClientManager] Health check failed for %s: %v", cm.urls[i], err)

			// Close unhealthy connection
			if cm.clients[i] != nil {
				cm.clients[i].Close()
				cm.clients[i] = nil
			}
		} else {
			if !cm.healthStatus[i] {
				log.Infof("[RPC.ClientManager] RPC client %s is back online", cm.urls[i])
			}
			cm.healthStatus[i] = true
		}
		cm.mutex.Unlock()
	}
}

// GetHealthyClientCount gets count of healthy clients
func (cm *ClientManager) GetHealthyClientCount() int {
	cm.mutex.RLock()
	defer cm.mutex.RUnlock()

	count := 0
	for _, isHealthy := range cm.healthStatus {
		if isHealthy {
			count++
		}
	}
	return count
}

// GetStatus gets status of all clients
func (cm *ClientManager) GetStatus() map[string]bool {
	cm.mutex.RLock()
	defer cm.mutex.RUnlock()

	status := make(map[string]bool)
	for i, url := range cm.urls {
		status[url] = cm.healthStatus[i]
	}
	return status
}

// HTTPPostJSONWithFailover sends JSON-RPC request through built-in HTTP manager
func (cm *ClientManager) HTTPPostJSONWithFailover(proxyURL, bodyJSON string) ([]byte, error) {
	if cm.httpManager != nil {
		return cm.httpManager.HTTPPostJSONWithFailover(proxyURL, bodyJSON)
	}

	// If no HTTP manager, fallback to first URL
	if len(cm.urls) > 0 {
		return HTTPPostJSON(proxyURL, cm.urls[0], bodyJSON)
	}

	return nil, fmt.Errorf("no available RPC URLs")
}

// Close closes all client connections
func (cm *ClientManager) Close() {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	for i, client := range cm.clients {
		if client != nil {
			client.Close()
			cm.clients[i] = nil
		}
	}
}
