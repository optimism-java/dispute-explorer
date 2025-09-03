package blockchain

import (
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	config "github.com/optimism-java/dispute-explorer/internal/types"
	"github.com/optimism-java/dispute-explorer/pkg/event"
)

type Event interface {
	Name() string
	EventHash() common.Hash
	Data(log types.Log) (string, error)
	ToObj(data string) error
}

var (
	events      = make([]common.Hash, 0)
	contracts   = make([]common.Address, 0)
	contractsMu sync.RWMutex // Protects concurrent access to contracts slice
	initOnce    sync.Once    // Ensures initialization only runs once
	EventMap    = make(map[common.Hash][]Event, 0)
	EIP1155     = make([]common.Address, 0)
)

func init() {
	Register(&event.DisputeGameCreated{})
	Register(&event.DisputeGameMove{})
	Register(&event.DisputeGameResolved{})
}

// InitContracts initializes contract addresses (runs only once)
func InitContracts() {
	initOnce.Do(func() {
		cfg := config.GetConfig()
		disputeGameProxys := strings.Split(cfg.DisputeGameProxyContract, ",")
		for _, one := range disputeGameProxys {
			AddContract(one)
		}
	})
}

func Register(event Event) {
	if len(EventMap[event.EventHash()]) == 0 {
		events = append(events, event.EventHash())
	}
	EventMap[event.EventHash()] = append(EventMap[event.EventHash()], event)
}

func AddContract(contract string) {
	contractsMu.Lock()
	defer contractsMu.Unlock()

	addr := common.HexToAddress(contract)
	// Check if already exists to prevent duplicate additions
	for _, existing := range contracts {
		if existing == addr {
			return // Already exists, return directly
		}
	}
	contracts = append(contracts, addr)
}

func RemoveContract(contract string) {
	contractsMu.Lock()
	defer contractsMu.Unlock()

	addr := common.HexToAddress(contract)
	// Use reverse iteration to safely remove all matching contracts
	for i := len(contracts) - 1; i >= 0; i-- {
		if contracts[i] == addr {
			contracts = append(contracts[:i], contracts[i+1:]...)
		}
	}
}

func GetContracts() []common.Address {
	contractsMu.RLock()
	defer contractsMu.RUnlock()

	// Return a copy to prevent external modifications
	result := make([]common.Address, len(contracts))
	copy(result, contracts)
	return result
}

func GetEvents() []common.Hash {
	return events
}

func GetEvent(eventHash common.Hash) Event {
	EventList := EventMap[eventHash]
	Event := EventList[0]
	return Event
}
