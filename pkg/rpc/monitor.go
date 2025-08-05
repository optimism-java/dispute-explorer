package rpc

import (
	"context"
	"fmt"
	"log"
	"time"
)

// Monitor RPC manager monitor
type Monitor struct {
	manager  *Manager
	interval time.Duration
	logger   Logger
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
		manager:  manager,
		interval: interval,
		logger:   &DefaultLogger{},
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
	stats := m.manager.GetStats()

	l1Rate, l1Burst := m.manager.GetRateLimit(true)
	l2Rate, l2Burst := m.manager.GetRateLimit(false)

	l1Tokens := m.manager.GetTokens(true)
	l2Tokens := m.manager.GetTokens(false)

	m.logger.Printf(
		"[RPC Stats] L1: %d requests (%d limited), L2: %d requests (%d limited), HTTP: %d",
		stats.L1RequestCount, stats.L1RateLimitedCount,
		stats.L2RequestCount, stats.L2RateLimitedCount,
		stats.HTTPRequestCount,
	)

	m.logger.Printf(
		"[RPC Limits] L1: %.1f/s (burst %d, tokens %.2f), L2: %.1f/s (burst %d, tokens %.2f)",
		l1Rate, l1Burst, l1Tokens,
		l2Rate, l2Burst, l2Tokens,
	)

	// warning messages
	if stats.L1RateLimitedCount > 0 || stats.L2RateLimitedCount > 0 {
		m.logger.Printf("[RPC Warning] Rate limiting is active!")
	}

	if l1Tokens < 1.0 {
		m.logger.Printf("[RPC Warning] L1 tokens running low: %.2f", l1Tokens)
	}

	if l2Tokens < 1.0 {
		m.logger.Printf("[RPC Warning] L2 tokens running low: %.2f", l2Tokens)
	}
}

// GetHealthCheck gets health check information
func (m *Monitor) GetHealthCheck() HealthCheckResult {
	stats := m.manager.GetStats()
	l1Tokens := m.manager.GetTokens(true)
	l2Tokens := m.manager.GetTokens(false)

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
