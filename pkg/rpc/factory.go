package rpc

import (
	"time"

	"github.com/optimism-java/dispute-explorer/internal/types"
)

// CreateManagerFromConfig creates RPC manager from configuration
func CreateManagerFromConfig(config *types.Config) (*Manager, error) {
	return NewManager(Config{
		L1RPCUrls:   config.GetL1RPCUrls(),
		L2RPCUrls:   config.GetL2RPCUrls(),
		NodeRPCUrls: config.GetNodeRPCUrls(),
		ProxyURL:    "",
		RateLimit:   config.RPCRateLimit,
		RateBurst:   config.RPCRateBurst,
		HTTPTimeout: 10 * time.Second,
	})
}

// CreateManagerWithSeparateLimits creates manager with different L1/L2 limits
func CreateManagerWithSeparateLimits(
	l1URL, l2URL string,
	l1Rate, l1Burst, _, _ int, // l2Rate, l2Burst unused for now
) (*Manager, error) {
	// Note: current implementation uses same limits for L1 and L2
	// if different limits are needed, Manager structure needs to be modified
	return NewManager(Config{
		L1RPCUrls:   []string{l1URL},
		L2RPCUrls:   []string{l2URL},
		RateLimit:   l1Rate, // 使用L1的限制作为默认
		RateBurst:   l1Burst,
		HTTPTimeout: 10 * time.Second,
	})
}

// WrapExistingClient wraps existing ethclient.Client (for backward compatibility)
func WrapExistingClient(config *types.Config, _, _ interface{}) (*Manager, error) {
	// create new manager but maintain backward compatibility
	manager, err := CreateManagerFromConfig(config)
	if err != nil {
		return nil, err
	}

	// logic can be added here to integrate existing clients
	// for now, return newly created manager
	return manager, nil
}
