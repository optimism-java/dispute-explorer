package rpc

import (
	"context"
	"fmt"
	"math/big"
	"net/http"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/time/rate"
)

const (
	// RPC type constants
	nodeRPCType = "Node"
)

// Manager unified RPC resource manager
type Manager struct {
	// Configuration
	l1RPCUrls   []string
	l2RPCUrls   []string
	nodeRPCUrls []string
	proxyURL    string

	// Round-robin indices
	l1Index   int
	l2Index   int
	nodeIndex int

	// Rate limiters
	l1Limiter   *rate.Limiter
	l2Limiter   *rate.Limiter
	nodeLimiter *rate.Limiter

	// Client pools
	l1Clients   []*ethclient.Client
	l2Clients   []*ethclient.Client
	nodeClients []*ethclient.Client

	// HTTP client
	httpClient *http.Client

	// Statistics
	stats *Stats
	mu    sync.RWMutex
}

// Config RPC manager configuration
type Config struct {
	L1RPCUrls   []string
	L2RPCUrls   []string
	NodeRPCUrls []string
	ProxyURL    string
	RateLimit   int
	RateBurst   int
	HTTPTimeout time.Duration
}

// Stats RPC call statistics
type Stats struct {
	L1RequestCount       int64
	L2RequestCount       int64
	NodeRequestCount     int64
	L1RateLimitedCount   int64
	L2RateLimitedCount   int64
	NodeRateLimitedCount int64
	HTTPRequestCount     int64
	LastRequestTime      time.Time
	mu                   sync.RWMutex
}

// NewManager creates a new RPC manager
func NewManager(config Config) (*Manager, error) {
	// Validate configuration
	if len(config.L1RPCUrls) == 0 {
		return nil, fmt.Errorf("l1 RPC URLs cannot be empty")
	}
	if len(config.L2RPCUrls) == 0 {
		return nil, fmt.Errorf("l2 RPC URLs cannot be empty")
	}
	if len(config.NodeRPCUrls) == 0 {
		return nil, fmt.Errorf("node RPC URLs cannot be empty")
	}

	// Create L1 client pool
	l1Clients := make([]*ethclient.Client, 0, len(config.L1RPCUrls))
	for _, url := range config.L1RPCUrls {
		client, err := ethclient.Dial(url)
		if err != nil {
			return nil, fmt.Errorf("failed to connect to L1 RPC %s: %w", url, err)
		}
		l1Clients = append(l1Clients, client)
	}

	// Create L2 client pool
	l2Clients := make([]*ethclient.Client, 0, len(config.L2RPCUrls))
	for _, url := range config.L2RPCUrls {
		client, err := ethclient.Dial(url)
		if err != nil {
			return nil, fmt.Errorf("failed to connect to L2 RPC %s: %w", url, err)
		}
		l2Clients = append(l2Clients, client)
	}

	// Create Node client pool
	nodeClients := make([]*ethclient.Client, 0, len(config.NodeRPCUrls))
	for _, url := range config.NodeRPCUrls {
		client, err := ethclient.Dial(url)
		if err != nil {
			return nil, fmt.Errorf("failed to connect to %s RPC %s: %w", nodeRPCType, url, err)
		}
		nodeClients = append(nodeClients, client)
	}

	// Create rate limiters
	l1Limiter := rate.NewLimiter(rate.Limit(config.RateLimit), config.RateBurst)
	l2Limiter := rate.NewLimiter(rate.Limit(config.RateLimit), config.RateBurst)
	nodeLimiter := rate.NewLimiter(rate.Limit(config.RateLimit), config.RateBurst)

	// Set HTTP timeout
	timeout := config.HTTPTimeout
	if timeout == 0 {
		timeout = 10 * time.Second
	}

	httpClient := &http.Client{
		Timeout: timeout,
	}

	return &Manager{
		l1RPCUrls:   config.L1RPCUrls,
		l2RPCUrls:   config.L2RPCUrls,
		nodeRPCUrls: config.NodeRPCUrls,
		proxyURL:    config.ProxyURL,
		l1Index:     0,
		l2Index:     0,
		nodeIndex:   0,
		l1Limiter:   l1Limiter,
		l2Limiter:   l2Limiter,
		nodeLimiter: nodeLimiter,
		l1Clients:   l1Clients,
		l2Clients:   l2Clients,
		nodeClients: nodeClients,
		httpClient:  httpClient,
		stats:       &Stats{},
	}, nil
}

// getNextL1 gets next L1 RPC URL and client using round-robin
func (m *Manager) getNextL1() (string, *ethclient.Client) {
	m.mu.Lock()
	defer m.mu.Unlock()

	url := m.l1RPCUrls[m.l1Index]
	client := m.l1Clients[m.l1Index]
	m.l1Index = (m.l1Index + 1) % len(m.l1RPCUrls)

	return url, client
}

// getNextL2 gets next L2 RPC URL and client using round-robin
func (m *Manager) getNextL2() (string, *ethclient.Client) {
	m.mu.Lock()
	defer m.mu.Unlock()

	url := m.l2RPCUrls[m.l2Index]
	client := m.l2Clients[m.l2Index]
	m.l2Index = (m.l2Index + 1) % len(m.l2RPCUrls)

	return url, client
}

// getNextNode gets next Node RPC URL and client using round-robin
func (m *Manager) getNextNode() (string, *ethclient.Client) {
	m.mu.Lock()
	defer m.mu.Unlock()

	url := m.nodeRPCUrls[m.nodeIndex]
	client := m.nodeClients[m.nodeIndex]
	m.nodeIndex = (m.nodeIndex + 1) % len(m.nodeRPCUrls)

	return url, client
}

// GetLatestBlockNumber gets the latest block number (with rate limiting)
func (m *Manager) GetLatestBlockNumber(ctx context.Context, isL1 bool) (uint64, error) {
	if isL1 {
		if err := m.l1Limiter.Wait(ctx); err != nil {
			m.updateRateLimitedStats("L1")
			return 0, fmt.Errorf("L1 rate limit exceeded: %w", err)
		}
		m.updateRequestStats("L1")
		_, client := m.getNextL1()
		return client.BlockNumber(ctx)
	}
	if err := m.l2Limiter.Wait(ctx); err != nil {
		m.updateRateLimitedStats("L2")
		return 0, fmt.Errorf("L2 rate limit exceeded: %w", err)
	}
	m.updateRequestStats("L2")
	_, client := m.getNextL2()
	return client.BlockNumber(ctx)
}

// GetBlockByNumber gets a block by number (with rate limiting)
func (m *Manager) GetBlockByNumber(ctx context.Context, number *big.Int, isL1 bool) (*types.Block, error) {
	if isL1 {
		if err := m.l1Limiter.Wait(ctx); err != nil {
			m.updateRateLimitedStats("L1")
			return nil, fmt.Errorf("L1 rate limit exceeded: %w", err)
		}
		m.updateRequestStats("L1")
		_, client := m.getNextL1()
		return client.BlockByNumber(ctx, number)
	}
	if err := m.l2Limiter.Wait(ctx); err != nil {
		m.updateRateLimitedStats("L2")
		return nil, fmt.Errorf("L2 rate limit exceeded: %w", err)
	}
	m.updateRequestStats("L2")
	_, client := m.getNextL2()
	return client.BlockByNumber(ctx, number)
}

// GetBlockByHash gets a block by hash (with rate limiting)
func (m *Manager) GetBlockByHash(ctx context.Context, hash common.Hash, isL1 bool) (*types.Block, error) {
	if isL1 {
		if err := m.l1Limiter.Wait(ctx); err != nil {
			m.updateRateLimitedStats("L1")
			return nil, fmt.Errorf("L1 rate limit exceeded: %w", err)
		}
		m.updateRequestStats("L1")
		_, client := m.getNextL1()
		return client.BlockByHash(ctx, hash)
	}
	if err := m.l2Limiter.Wait(ctx); err != nil {
		m.updateRateLimitedStats("L2")
		return nil, fmt.Errorf("L2 rate limit exceeded: %w", err)
	}
	m.updateRequestStats("L2")
	_, client := m.getNextL2()
	return client.BlockByHash(ctx, hash)
}

// FilterLogs filters logs (with rate limiting)
func (m *Manager) FilterLogs(ctx context.Context, query ethereum.FilterQuery, isL1 bool) ([]types.Log, error) {
	if isL1 {
		if err := m.l1Limiter.Wait(ctx); err != nil {
			m.updateRateLimitedStats("L1")
			return nil, fmt.Errorf("L1 rate limit exceeded: %w", err)
		}
		m.updateRequestStats("L1")
		_, client := m.getNextL1()
		return client.FilterLogs(ctx, query)
	}
	if err := m.l2Limiter.Wait(ctx); err != nil {
		m.updateRateLimitedStats("L2")
		return nil, fmt.Errorf("L2 rate limit exceeded: %w", err)
	}
	m.updateRequestStats("L2")
	_, client := m.getNextL2()
	return client.FilterLogs(ctx, query)
}

// HeaderByNumber gets a block header by number (with rate limiting)
func (m *Manager) HeaderByNumber(ctx context.Context, number *big.Int, isL1 bool) (*types.Header, error) {
	if isL1 {
		if err := m.l1Limiter.Wait(ctx); err != nil {
			m.updateRateLimitedStats("L1")
			return nil, fmt.Errorf("L1 rate limit exceeded: %w", err)
		}
		m.updateRequestStats("L1")
		_, client := m.getNextL1()
		return client.HeaderByNumber(ctx, number)
	}
	if err := m.l2Limiter.Wait(ctx); err != nil {
		m.updateRateLimitedStats("L2")
		return nil, fmt.Errorf("L2 rate limit exceeded: %w", err)
	}
	m.updateRequestStats("L2")
	_, client := m.getNextL2()
	return client.HeaderByNumber(ctx, number)
}

// CallContract calls a smart contract (with rate limiting)
func (m *Manager) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int, isL1 bool) ([]byte, error) {
	if isL1 {
		if err := m.l1Limiter.Wait(ctx); err != nil {
			m.updateRateLimitedStats("L1")
			return nil, fmt.Errorf("L1 rate limit exceeded: %w", err)
		}
		m.updateRequestStats("L1")
		_, client := m.getNextL1()
		return client.CallContract(ctx, call, blockNumber)
	}
	if err := m.l2Limiter.Wait(ctx); err != nil {
		m.updateRateLimitedStats("L2")
		return nil, fmt.Errorf("L2 rate limit exceeded: %w", err)
	}
	m.updateRequestStats("L2")
	_, client := m.getNextL2()
	return client.CallContract(ctx, call, blockNumber)
}

// HTTPPostJSON HTTP POST JSON request (with rate limiting)
func (m *Manager) HTTPPostJSON(ctx context.Context, bodyJSON string, isL1 bool) ([]byte, error) {
	var rpcURL string

	if isL1 {
		if err := m.l1Limiter.Wait(ctx); err != nil {
			m.updateRateLimitedStats("L1")
			return nil, fmt.Errorf("L1 rate limit exceeded: %w", err)
		}
		m.updateRequestStats("L1")
		rpcURL, _ = m.getNextL1()
	} else {
		if err := m.l2Limiter.Wait(ctx); err != nil {
			m.updateRateLimitedStats("L2")
			return nil, fmt.Errorf("L2 rate limit exceeded: %w", err)
		}
		m.updateRequestStats("L2")
		rpcURL, _ = m.getNextL2()
	}

	m.updateHTTPRequestStats()
	return HTTPPostJSON(m.proxyURL, rpcURL, bodyJSON)
}

// GetRawClient gets the raw client (for backward compatibility)
func (m *Manager) GetRawClient(isL1 bool) *ethclient.Client {
	if isL1 {
		_, client := m.getNextL1()
		return client
	}
	_, client := m.getNextL2()
	return client
}

// GetNodeRawClient gets the Node RPC raw client using round-robin
func (m *Manager) GetNodeRawClient() *ethclient.Client {
	_, client := m.getNextNode()
	return client
}

// GetNodeRPCURL gets the next Node RPC URL using round-robin
func (m *Manager) GetNodeRPCURL() string {
	url, _ := m.getNextNode()
	return url
}

// UpdateRateLimit dynamically updates rate limit
func (m *Manager) UpdateRateLimit(rateLimit int, rateBurst int, rpcType string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	switch rpcType {
	case "L1":
		m.l1Limiter.SetLimit(rate.Limit(rateLimit))
		m.l1Limiter.SetBurst(rateBurst)
	case "L2":
		m.l2Limiter.SetLimit(rate.Limit(rateLimit))
		m.l2Limiter.SetBurst(rateBurst)
	case nodeRPCType:
		m.nodeLimiter.SetLimit(rate.Limit(rateLimit))
		m.nodeLimiter.SetBurst(rateBurst)
	}
}

// GetRateLimit gets current rate limit settings
func (m *Manager) GetRateLimit(rpcType string) (rateLimit float64, rateBurst int) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	switch rpcType {
	case "L1":
		return float64(m.l1Limiter.Limit()), m.l1Limiter.Burst()
	case "L2":
		return float64(m.l2Limiter.Limit()), m.l2Limiter.Burst()
	case nodeRPCType:
		return float64(m.nodeLimiter.Limit()), m.nodeLimiter.Burst()
	default:
		return 0, 0
	}
}

// GetTokens 返回当前可用的令牌数
func (m *Manager) GetTokens(rpcType string) float64 {
	switch rpcType {
	case "L1":
		return m.l1Limiter.Tokens()
	case "L2":
		return m.l2Limiter.Tokens()
	case nodeRPCType:
		return m.nodeLimiter.Tokens()
	default:
		return 0
	}
}

// GetStats gets statistics information
func (m *Manager) GetStats() StatsSnapshot {
	m.stats.mu.RLock()
	defer m.stats.mu.RUnlock()

	return StatsSnapshot{
		L1RequestCount:       m.stats.L1RequestCount,
		L2RequestCount:       m.stats.L2RequestCount,
		NodeRequestCount:     m.stats.NodeRequestCount,
		L1RateLimitedCount:   m.stats.L1RateLimitedCount,
		L2RateLimitedCount:   m.stats.L2RateLimitedCount,
		NodeRateLimitedCount: m.stats.NodeRateLimitedCount,
		HTTPRequestCount:     m.stats.HTTPRequestCount,
		LastRequestTime:      m.stats.LastRequestTime,
	}
}

// StatsSnapshot statistics snapshot
type StatsSnapshot struct {
	L1RequestCount       int64
	L2RequestCount       int64
	NodeRequestCount     int64
	L1RateLimitedCount   int64
	L2RateLimitedCount   int64
	NodeRateLimitedCount int64
	HTTPRequestCount     int64
	LastRequestTime      time.Time
}

// updateRequestStats updates request statistics
func (m *Manager) updateRequestStats(rpcType string) {
	m.stats.mu.Lock()
	defer m.stats.mu.Unlock()

	m.stats.LastRequestTime = time.Now()
	switch rpcType {
	case "L1":
		m.stats.L1RequestCount++
	case "L2":
		m.stats.L2RequestCount++
	case nodeRPCType:
		m.stats.NodeRequestCount++
	}
}

// updateRateLimitedStats updates rate-limited request statistics
func (m *Manager) updateRateLimitedStats(rpcType string) {
	m.stats.mu.Lock()
	defer m.stats.mu.Unlock()

	switch rpcType {
	case "L1":
		m.stats.L1RateLimitedCount++
	case "L2":
		m.stats.L2RateLimitedCount++
	case nodeRPCType:
		m.stats.NodeRateLimitedCount++
	}
}

// updateHTTPRequestStats 更新HTTP请求统计
func (m *Manager) updateHTTPRequestStats() {
	m.stats.mu.Lock()
	defer m.stats.mu.Unlock()

	m.stats.HTTPRequestCount++
}

// Close closes all connections
func (m *Manager) Close() {
	// Close all L1 clients
	for _, client := range m.l1Clients {
		if client != nil {
			client.Close()
		}
	}
	// Close all L2 clients
	for _, client := range m.l2Clients {
		if client != nil {
			client.Close()
		}
	}
	// Close all Node clients
	for _, client := range m.nodeClients {
		if client != nil {
			client.Close()
		}
	}
}
