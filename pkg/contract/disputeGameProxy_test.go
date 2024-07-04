package contract

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetGameCount(t *testing.T) {
	l1rpc, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/PNunSRFo0FWRJMu5yrwBd6jF7G78YHrv")
	require.NoError(t, err)
	disputeGame, err := NewDisputeGameProxy(common.HexToAddress("0x05F9613aDB30026FFd634f38e5C4dFd30a197Fa1"), l1rpc)
	require.NoError(t, err)
	count, err := disputeGame.GameCount(&bind.CallOpts{
		BlockHash: common.HexToHash("0xeab46f31379b051a350dbb6f569c6ccae8043a1637cc37958e8e15e3210e220b"),
	})
	require.NoError(t, err)
	fmt.Println(count)
}
