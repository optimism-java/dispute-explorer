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

// Manager unified RPC resource manager
type Manager struct {
	// Configuration
	l1RPCUrl string
	l2RPCUrl string
	proxyURL string

	// Rate limiters
	l1Limiter *rate.Limiter
	l2Limiter *rate.Limiter

	// Native Ethereum clients
	l1Client *ethclient.Client
	l2Client *ethclient.Client

	// HTTP client
	httpClient *http.Client

	// Statistics
	stats *Stats
	mu    sync.RWMutex
}

// Config RPC manager configuration
type Config struct {
	L1RPCUrl    string
	L2RPCUrl    string
	ProxyURL    string
	RateLimit   int
	RateBurst   int
	HTTPTimeout time.Duration
}

// Stats RPC call statistics
type Stats struct {
	L1RequestCount     int64
	L2RequestCount     int64
	L1RateLimitedCount int64
	L2RateLimitedCount int64
	HTTPRequestCount   int64
	LastRequestTime    time.Time
	mu                 sync.RWMutex
}

// NewManager creates a new RPC manager
func NewManager(config Config) (*Manager, error) {
	// Create Ethereum clients
	l1Client, err := ethclient.Dial(config.L1RPCUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to L1 RPC: %w", err)
	}

	l2Client, err := ethclient.Dial(config.L2RPCUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to L2 RPC: %w", err)
	}

	// Create rate limiters
	l1Limiter := rate.NewLimiter(rate.Limit(config.RateLimit), config.RateBurst)
	l2Limiter := rate.NewLimiter(rate.Limit(config.RateLimit), config.RateBurst)

	// Set HTTP timeout
	timeout := config.HTTPTimeout
	if timeout == 0 {
		timeout = 10 * time.Second
	}

	httpClient := &http.Client{
		Timeout: timeout,
	}

	return &Manager{
		l1RPCUrl:   config.L1RPCUrl,
		l2RPCUrl:   config.L2RPCUrl,
		proxyURL:   config.ProxyURL,
		l1Limiter:  l1Limiter,
		l2Limiter:  l2Limiter,
		l1Client:   l1Client,
		l2Client:   l2Client,
		httpClient: httpClient,
		stats:      &Stats{},
	}, nil
}

// GetLatestBlockNumber gets the latest block number (with rate limiting)
func (m *Manager) GetLatestBlockNumber(ctx context.Context, isL1 bool) (uint64, error) {
	if isL1 {
		if err := m.l1Limiter.Wait(ctx); err != nil {
			m.updateRateLimitedStats(true)
			return 0, fmt.Errorf("L1 rate limit exceeded: %w", err)
		}
		m.updateRequestStats(true)
		return m.l1Client.BlockNumber(ctx)
	}
	if err := m.l2Limiter.Wait(ctx); err != nil {
		m.updateRateLimitedStats(false)
		return 0, fmt.Errorf("L2 rate limit exceeded: %w", err)
	}
	m.updateRequestStats(false)
	return m.l2Client.BlockNumber(ctx)
}

// GetBlockByNumber gets a block by number (with rate limiting)
func (m *Manager) GetBlockByNumber(ctx context.Context, number *big.Int, isL1 bool) (*types.Block, error) {
	if isL1 {
		if err := m.l1Limiter.Wait(ctx); err != nil {
			m.updateRateLimitedStats(true)
			return nil, fmt.Errorf("L1 rate limit exceeded: %w", err)
		}
		m.updateRequestStats(true)
		return m.l1Client.BlockByNumber(ctx, number)
	}
	if err := m.l2Limiter.Wait(ctx); err != nil {
		m.updateRateLimitedStats(false)
		return nil, fmt.Errorf("L2 rate limit exceeded: %w", err)
	}
	m.updateRequestStats(false)
	return m.l2Client.BlockByNumber(ctx, number)
}

// GetBlockByHash gets a block by hash (with rate limiting)
func (m *Manager) GetBlockByHash(ctx context.Context, hash common.Hash, isL1 bool) (*types.Block, error) {
	if isL1 {
		if err := m.l1Limiter.Wait(ctx); err != nil {
			m.updateRateLimitedStats(true)
			return nil, fmt.Errorf("L1 rate limit exceeded: %w", err)
		}
		m.updateRequestStats(true)
		return m.l1Client.BlockByHash(ctx, hash)
	}
	if err := m.l2Limiter.Wait(ctx); err != nil {
		m.updateRateLimitedStats(false)
		return nil, fmt.Errorf("L2 rate limit exceeded: %w", err)
	}
	m.updateRequestStats(false)
	return m.l2Client.BlockByHash(ctx, hash)
}

// FilterLogs filters logs (with rate limiting)
func (m *Manager) FilterLogs(ctx context.Context, query ethereum.FilterQuery, isL1 bool) ([]types.Log, error) {
	if isL1 {
		if err := m.l1Limiter.Wait(ctx); err != nil {
			m.updateRateLimitedStats(true)
			return nil, fmt.Errorf("L1 rate limit exceeded: %w", err)
		}
		m.updateRequestStats(true)
		return m.l1Client.FilterLogs(ctx, query)
	}
	if err := m.l2Limiter.Wait(ctx); err != nil {
		m.updateRateLimitedStats(false)
		return nil, fmt.Errorf("L2 rate limit exceeded: %w", err)
	}
	m.updateRequestStats(false)
	return m.l2Client.FilterLogs(ctx, query)
}

// HeaderByNumber gets a block header by number (with rate limiting)
func (m *Manager) HeaderByNumber(ctx context.Context, number *big.Int, isL1 bool) (*types.Header, error) {
	if isL1 {
		if err := m.l1Limiter.Wait(ctx); err != nil {
			m.updateRateLimitedStats(true)
			return nil, fmt.Errorf("L1 rate limit exceeded: %w", err)
		}
		m.updateRequestStats(true)
		return m.l1Client.HeaderByNumber(ctx, number)
	}
	if err := m.l2Limiter.Wait(ctx); err != nil {
		m.updateRateLimitedStats(false)
		return nil, fmt.Errorf("L2 rate limit exceeded: %w", err)
	}
	m.updateRequestStats(false)
	return m.l2Client.HeaderByNumber(ctx, number)
}

// CallContract calls a smart contract (with rate limiting)
func (m *Manager) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int, isL1 bool) ([]byte, error) {
	if isL1 {
		if err := m.l1Limiter.Wait(ctx); err != nil {
			m.updateRateLimitedStats(true)
			return nil, fmt.Errorf("L1 rate limit exceeded: %w", err)
		}
		m.updateRequestStats(true)
		return m.l1Client.CallContract(ctx, call, blockNumber)
	}
	if err := m.l2Limiter.Wait(ctx); err != nil {
		m.updateRateLimitedStats(false)
		return nil, fmt.Errorf("L2 rate limit exceeded: %w", err)
	}
	m.updateRequestStats(false)
	return m.l2Client.CallContract(ctx, call, blockNumber)
}

// HTTPPostJSON HTTP POST JSON request (with rate limiting)
func (m *Manager) HTTPPostJSON(ctx context.Context, bodyJSON string, isL1 bool) ([]byte, error) {
	var rpcURL string

	if isL1 {
		if err := m.l1Limiter.Wait(ctx); err != nil {
			m.updateRateLimitedStats(true)
			return nil, fmt.Errorf("L1 rate limit exceeded: %w", err)
		}
		m.updateRequestStats(true)
		rpcURL = m.l1RPCUrl
	} else {
		if err := m.l2Limiter.Wait(ctx); err != nil {
			m.updateRateLimitedStats(false)
			return nil, fmt.Errorf("L2 rate limit exceeded: %w", err)
		}
		m.updateRequestStats(false)
		rpcURL = m.l2RPCUrl
	}

	m.updateHTTPRequestStats()
	return HTTPPostJSON(m.proxyURL, rpcURL, bodyJSON)
}

// GetRawClient gets the raw client (for backward compatibility)
func (m *Manager) GetRawClient(isL1 bool) *ethclient.Client {
	if isL1 {
		return m.l1Client
	}
	return m.l2Client
}

// UpdateRateLimit dynamically updates rate limit
func (m *Manager) UpdateRateLimit(rateLimit int, rateBurst int, isL1 bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if isL1 {
		m.l1Limiter.SetLimit(rate.Limit(rateLimit))
		m.l1Limiter.SetBurst(rateBurst)
	} else {
		m.l2Limiter.SetLimit(rate.Limit(rateLimit))
		m.l2Limiter.SetBurst(rateBurst)
	}
}

// GetRateLimit gets current rate limit settings
func (m *Manager) GetRateLimit(isL1 bool) (rateLimit float64, rateBurst int) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if isL1 {
		return float64(m.l1Limiter.Limit()), m.l1Limiter.Burst()
	}
	return float64(m.l2Limiter.Limit()), m.l2Limiter.Burst()
}

// GetTokens 返回当前可用的令牌数
func (m *Manager) GetTokens(isL1 bool) float64 {
	if isL1 {
		return m.l1Limiter.Tokens()
	}
	return m.l2Limiter.Tokens()
}

// GetStats gets statistics information
func (m *Manager) GetStats() StatsSnapshot {
	m.stats.mu.RLock()
	defer m.stats.mu.RUnlock()

	return StatsSnapshot{
		L1RequestCount:     m.stats.L1RequestCount,
		L2RequestCount:     m.stats.L2RequestCount,
		L1RateLimitedCount: m.stats.L1RateLimitedCount,
		L2RateLimitedCount: m.stats.L2RateLimitedCount,
		HTTPRequestCount:   m.stats.HTTPRequestCount,
		LastRequestTime:    m.stats.LastRequestTime,
	}
}

// StatsSnapshot statistics snapshot
type StatsSnapshot struct {
	L1RequestCount     int64
	L2RequestCount     int64
	L1RateLimitedCount int64
	L2RateLimitedCount int64
	HTTPRequestCount   int64
	LastRequestTime    time.Time
}

// updateRequestStats updates request statistics
func (m *Manager) updateRequestStats(isL1 bool) {
	m.stats.mu.Lock()
	defer m.stats.mu.Unlock()

	m.stats.LastRequestTime = time.Now()
	if isL1 {
		m.stats.L1RequestCount++
	} else {
		m.stats.L2RequestCount++
	}
}

// updateRateLimitedStats updates rate-limited request statistics
func (m *Manager) updateRateLimitedStats(isL1 bool) {
	m.stats.mu.Lock()
	defer m.stats.mu.Unlock()

	if isL1 {
		m.stats.L1RateLimitedCount++
	} else {
		m.stats.L2RateLimitedCount++
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
	if m.l1Client != nil {
		m.l1Client.Close()
	}
	if m.l2Client != nil {
		m.l2Client.Close()
	}
}
