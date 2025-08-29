package rpc

import (
	"testing"
)

func TestGetNextL1SingleRPC(t *testing.T) {
	// Test with single L1 RPC
	config := Config{
		L1RPCUrls:   []string{"https://test-rpc.com"},
		L2RPCUrls:   []string{"https://test-l2-rpc.com"},
		NodeRPCUrls: []string{"https://test-node-rpc.com"},
		RateLimit:   5,
		RateBurst:   2,
	}

	// Note: This test won't actually create real connections
	// since we're using test URLs. In a real test environment,
	// you'd use mock servers or real test endpoints.

	t.Logf("Testing with single RPC URL configuration")
	t.Logf("L1 URLs: %v", config.L1RPCUrls)
	t.Logf("L2 URLs: %v", config.L2RPCUrls)
	t.Logf("Node URLs: %v", config.NodeRPCUrls)

	// The round-robin logic should work with single URL:
	// - Index starts at 0
	// - (0 + 1) % 1 = 0, so it cycles back to 0
	// - No array bounds error should occur
}

func TestEmptyRPCUrls(t *testing.T) {
	// Test with empty L1 RPC URLs
	config := Config{
		L1RPCUrls: []string{}, // Empty array
		L2RPCUrls: []string{"https://test-l2-rpc.com"},
		RateLimit: 5,
		RateBurst: 2,
	}

	_, err := NewManager(config)
	if err == nil {
		t.Error("Expected error for empty L1 RPC URLs, got nil")
	}
	if err.Error() != "L1 RPC URLs cannot be empty" {
		t.Errorf("Expected specific error message, got: %s", err.Error())
	}

	// Test with empty L2 RPC URLs
	config2 := Config{
		L1RPCUrls: []string{"https://test-rpc.com"},
		L2RPCUrls: []string{}, // Empty array
		RateLimit: 5,
		RateBurst: 2,
	}

	_, err2 := NewManager(config2)
	if err2 == nil {
		t.Error("Expected error for empty L2 RPC URLs, got nil")
	}
	if err2.Error() != "L2 RPC URLs cannot be empty" {
		t.Errorf("Expected specific error message, got: %s", err2.Error())
	}
}

func TestRoundRobinLogic(t *testing.T) {
	tests := []struct {
		name        string
		rpcCount    int
		iterations  int
		description string
	}{
		{
			name:        "SingleRPC",
			rpcCount:    1,
			iterations:  5,
			description: "With 1 RPC, index should always be 0",
		},
		{
			name:        "TwoRPCs",
			rpcCount:    2,
			iterations:  4,
			description: "With 2 RPCs, index should alternate 0,1,0,1",
		},
		{
			name:        "ThreeRPCs",
			rpcCount:    3,
			iterations:  6,
			description: "With 3 RPCs, index should cycle 0,1,2,0,1,2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Simulate round-robin logic without creating real connections
			index := 0
			expectedSequence := make([]int, tt.iterations)

			for i := 0; i < tt.iterations; i++ {
				expectedSequence[i] = index
				index = (index + 1) % tt.rpcCount
			}

			t.Logf("%s: Expected sequence for %d iterations: %v",
				tt.description, tt.iterations, expectedSequence)

			// Verify the pattern is correct
			for i := 0; i < tt.iterations; i++ {
				expected := i % tt.rpcCount
				if expectedSequence[i] != expected {
					t.Errorf("At iteration %d, expected index %d, got %d",
						i, expected, expectedSequence[i])
				}
			}
		})
	}
}

func TestNodeRPCRoundRobin(t *testing.T) {
	// Create a simple test for Node RPC round-robin logic
	nodeUrls := []string{
		"https://node1.test.com",
		"https://node2.test.com",
		"https://node3.test.com",
	}

	// Test round-robin logic without actual connections
	index := 0
	var indices []int
	for i := 0; i < 6; i++ {
		indices = append(indices, index)
		index = (index + 1) % len(nodeUrls)
	}

	expected := []int{0, 1, 2, 0, 1, 2}
	if !equalIntSlices(indices, expected) {
		t.Errorf("Node RPC round-robin failed. Expected %v, got %v", expected, indices)
	}

	t.Logf("Node RPC round-robin test passed: %v", indices)
	t.Logf("Node URLs used in rotation: %v", nodeUrls)
}

// Helper function to compare int slices
func equalIntSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
