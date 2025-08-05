package rpc

import (
	"time"

	"github.com/optimism-java/dispute-explorer/internal/types"
)

// CreateManagerFromConfig creates RPC manager from configuration
func CreateManagerFromConfig(config *types.Config) (*Manager, error) {
	return NewManager(Config{
		L1RPCUrl:    config.L1RPCUrl,
		L2RPCUrl:    config.L2RPCUrl,
		ProxyURL:    "", // if proxy is needed, can be added from configuration
		RateLimit:   config.RPCRateLimit,
		RateBurst:   config.RPCRateBurst,
		HTTPTimeout: 10 * time.Second,
	})
}

// CreateManagerWithSeparateLimits creates manager with different L1/L2 limits
func CreateManagerWithSeparateLimits(
	l1URL, l2URL string,
	l1Rate, l1Burst, l2Rate, l2Burst int,
) (*Manager, error) {
	// Note: current implementation uses same limits for L1 and L2
	// if different limits are needed, Manager structure needs to be modified
	return NewManager(Config{
		L1RPCUrl:    l1URL,
		L2RPCUrl:    l2URL,
		RateLimit:   l1Rate, // 使用L1的限制作为默认
		RateBurst:   l1Burst,
		HTTPTimeout: 10 * time.Second,
	})
}

// WrapExistingClient wraps existing ethclient.Client (for backward compatibility)
func WrapExistingClient(config *types.Config, existingL1, existingL2 interface{}) (*Manager, error) {
	// create new manager but maintain backward compatibility
	manager, err := CreateManagerFromConfig(config)
	if err != nil {
		return nil, err
	}

	// logic can be added here to integrate existing clients
	// for now, return newly created manager
	return manager, nil
}
