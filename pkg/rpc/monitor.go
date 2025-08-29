package rpc

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

// PerMinuteStats 每分钟统计数据
type PerMinuteStats struct {
	L1RequestsPerMin      float64
	L2RequestsPerMin      float64
	NodeRequestsPerMin    float64
	L1RateLimitedPerMin   float64
	L2RateLimitedPerMin   float64
	NodeRateLimitedPerMin float64
	HTTPRequestsPerMin    float64
	TotalGoroutines       int
	ActiveGoroutines      int
}

// Monitor RPC manager monitor
type Monitor struct {
	manager   *Manager
	interval  time.Duration
	logger    Logger
	lastStats StatsSnapshot
	lastTime  time.Time
	mu        sync.RWMutex
}

// Logger logging interface
type Logger interface {
	Printf(format string, v ...interface{})
}

// DefaultLogger default logging implementation
type DefaultLogger struct{}

func (dl *DefaultLogger) Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

// NewMonitor creates monitor
func NewMonitor(manager *Manager, interval time.Duration) *Monitor {
	return &Monitor{
		manager:   manager,
		interval:  interval,
		logger:    &DefaultLogger{},
		lastStats: manager.GetStats(), // 初始化统计数据
		lastTime:  time.Now(),
	}
}

// SetLogger sets custom logger
func (m *Monitor) SetLogger(logger Logger) {
	m.logger = logger
}

// Start starts monitoring
func (m *Monitor) Start(ctx context.Context) {
	ticker := time.NewTicker(m.interval)
	defer ticker.Stop()

	m.logger.Printf("[RPC Monitor] Started with interval %v", m.interval)

	for {
		select {
		case <-ticker.C:
			m.logStats()
		case <-ctx.Done():
			m.logger.Printf("[RPC Monitor] Stopped")
			return
		}
	}
}

// logStats logs statistics information
func (m *Monitor) logStats() {
	m.mu.Lock()
	defer m.mu.Unlock()

	currentStats := m.manager.GetStats()
	currentTime := time.Now()

	// 计算时间差（分钟）
	timeDiff := currentTime.Sub(m.lastTime).Minutes()
	if timeDiff == 0 {
		return // 避免除零
	}

	// 计算每分钟的增量
	perMinStats := PerMinuteStats{
		L1RequestsPerMin:      float64(currentStats.L1RequestCount-m.lastStats.L1RequestCount) / timeDiff,
		L2RequestsPerMin:      float64(currentStats.L2RequestCount-m.lastStats.L2RequestCount) / timeDiff,
		NodeRequestsPerMin:    float64(currentStats.NodeRequestCount-m.lastStats.NodeRequestCount) / timeDiff,
		L1RateLimitedPerMin:   float64(currentStats.L1RateLimitedCount-m.lastStats.L1RateLimitedCount) / timeDiff,
		L2RateLimitedPerMin:   float64(currentStats.L2RateLimitedCount-m.lastStats.L2RateLimitedCount) / timeDiff,
		NodeRateLimitedPerMin: float64(currentStats.NodeRateLimitedCount-m.lastStats.NodeRateLimitedCount) / timeDiff,
		HTTPRequestsPerMin:    float64(currentStats.HTTPRequestCount-m.lastStats.HTTPRequestCount) / timeDiff,
	}

	// 获取限流信息
	l1Rate, l1Burst := m.manager.GetRateLimit("L1")
	l2Rate, l2Burst := m.manager.GetRateLimit("L2")
	nodeRate, nodeBurst := m.manager.GetRateLimit("Node")

	l1Tokens := m.manager.GetTokens("L1")
	l2Tokens := m.manager.GetTokens("L2")
	nodeTokens := m.manager.GetTokens("Node")

	// 输出每分钟平均统计
	m.logger.Printf(
		"[RPC PerMin] L1: %.1f req/min (%.1f limited/min), L2: %.1f req/min (%.1f limited/min), Node: %.1f req/min (%.1f limited/min), HTTP: %.1f req/min",
		perMinStats.L1RequestsPerMin, perMinStats.L1RateLimitedPerMin,
		perMinStats.L2RequestsPerMin, perMinStats.L2RateLimitedPerMin,
		perMinStats.NodeRequestsPerMin, perMinStats.NodeRateLimitedPerMin,
		perMinStats.HTTPRequestsPerMin,
	)

	m.logger.Printf(
		"[RPC Limits] L1: %.1f/s (burst %d, tokens %.2f), L2: %.1f/s (burst %d, tokens %.2f), Node: %.1f/s (burst %d, tokens %.2f)",
		l1Rate, l1Burst, l1Tokens,
		l2Rate, l2Burst, l2Tokens,
		nodeRate, nodeBurst, nodeTokens,
	)

	// 输出累计统计（可选）
	m.logger.Printf(
		"[RPC Total] L1: %d total (%d limited), L2: %d total (%d limited), Node: %d total (%d limited), HTTP: %d total",
		currentStats.L1RequestCount, currentStats.L1RateLimitedCount,
		currentStats.L2RequestCount, currentStats.L2RateLimitedCount,
		currentStats.NodeRequestCount, currentStats.NodeRateLimitedCount,
		currentStats.HTTPRequestCount,
	)

	// warning messages
	if perMinStats.L1RateLimitedPerMin > 0 || perMinStats.L2RateLimitedPerMin > 0 || perMinStats.NodeRateLimitedPerMin > 0 {
		m.logger.Printf("[RPC Warning] Rate limiting is active! L1: %.1f/min, L2: %.1f/min, Node: %.1f/min",
			perMinStats.L1RateLimitedPerMin, perMinStats.L2RateLimitedPerMin, perMinStats.NodeRateLimitedPerMin)
	}

	if l1Tokens < 1.0 {
		m.logger.Printf("[RPC Warning] L1 tokens running low: %.2f", l1Tokens)
	}

	if l2Tokens < 1.0 {
		m.logger.Printf("[RPC Warning] L2 tokens running low: %.2f", l2Tokens)
	}

	if nodeTokens < 1.0 {
		m.logger.Printf("[RPC Warning] Node tokens running low: %.2f", nodeTokens)
	}

	// 更新记录
	m.lastStats = currentStats
	m.lastTime = currentTime
}

// GetPerMinuteStats 获取每分钟统计信息
func (m *Monitor) GetPerMinuteStats() PerMinuteStats {
	m.mu.RLock()
	defer m.mu.RUnlock()

	currentStats := m.manager.GetStats()
	currentTime := time.Now()

	timeDiff := currentTime.Sub(m.lastTime).Minutes()
	if timeDiff == 0 {
		return PerMinuteStats{}
	}

	return PerMinuteStats{
		L1RequestsPerMin:      float64(currentStats.L1RequestCount-m.lastStats.L1RequestCount) / timeDiff,
		L2RequestsPerMin:      float64(currentStats.L2RequestCount-m.lastStats.L2RequestCount) / timeDiff,
		NodeRequestsPerMin:    float64(currentStats.NodeRequestCount-m.lastStats.NodeRequestCount) / timeDiff,
		L1RateLimitedPerMin:   float64(currentStats.L1RateLimitedCount-m.lastStats.L1RateLimitedCount) / timeDiff,
		L2RateLimitedPerMin:   float64(currentStats.L2RateLimitedCount-m.lastStats.L2RateLimitedCount) / timeDiff,
		NodeRateLimitedPerMin: float64(currentStats.NodeRateLimitedCount-m.lastStats.NodeRateLimitedCount) / timeDiff,
		HTTPRequestsPerMin:    float64(currentStats.HTTPRequestCount-m.lastStats.HTTPRequestCount) / timeDiff,
	}
}

// GetHealthCheck gets health check information
func (m *Monitor) GetHealthCheck() HealthCheckResult {
	stats := m.manager.GetStats()
	l1Tokens := m.manager.GetTokens("L1")
	l2Tokens := m.manager.GetTokens("L2")

	isHealthy := true
	var issues []string

	// check token count
	if l1Tokens < 1 {
		isHealthy = false
		issues = append(issues, "L1 rate limit exhausted")
	}

	if l2Tokens < 1 {
		isHealthy = false
		issues = append(issues, "L2 rate limit exhausted")
	}

	// check if there are recent rate limits
	if stats.L1RateLimitedCount > 0 {
		issues = append(issues, fmt.Sprintf("L1 has %d rate limited requests", stats.L1RateLimitedCount))
	}

	if stats.L2RateLimitedCount > 0 {
		issues = append(issues, fmt.Sprintf("L2 has %d rate limited requests", stats.L2RateLimitedCount))
	}

	// check if there are recent requests
	if time.Since(stats.LastRequestTime) > 5*time.Minute {
		issues = append(issues, "No recent RPC requests detected")
	}

	return HealthCheckResult{
		Healthy:           isHealthy,
		Issues:            issues,
		L1Tokens:          l1Tokens,
		L2Tokens:          l2Tokens,
		LastRequest:       stats.LastRequestTime,
		TotalL1Requests:   stats.L1RequestCount,
		TotalL2Requests:   stats.L2RequestCount,
		TotalHTTPRequests: stats.HTTPRequestCount,
	}
}

// HealthCheckResult health check result
type HealthCheckResult struct {
	Healthy           bool
	Issues            []string
	L1Tokens          float64
	L2Tokens          float64
	LastRequest       time.Time
	TotalL1Requests   int64
	TotalL2Requests   int64
	TotalHTTPRequests int64
}

// LogHealthCheck logs health check result
func (m *Monitor) LogHealthCheck() {
	health := m.GetHealthCheck()

	if health.Healthy {
		m.logger.Printf("[RPC Health] System is healthy")
	} else {
		m.logger.Printf("[RPC Health] System has issues: %v", health.Issues)
	}

	m.logger.Printf("[RPC Health] Total requests - L1: %d, L2: %d, HTTP: %d",
		health.TotalL1Requests, health.TotalL2Requests, health.TotalHTTPRequests)
}
