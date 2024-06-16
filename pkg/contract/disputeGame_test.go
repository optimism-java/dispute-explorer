package contract

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestContract(t *testing.T) {
	l1rpc, err := ethclient.Dial("https://quaint-white-season.ethereum-sepolia.quiknode.pro/b5c30cbb548d8743f08dd175fe50e3e923259d30")
	require.NoError(t, err)
	disputeGame, err := NewDisputeGame(common.HexToAddress("0x8304B519e45133A11E07b356443dC39bEf881D83"), l1rpc)
	require.NoError(t, err)
	count, err := disputeGame.ClaimDataLen(&bind.CallOpts{})
	require.NoError(t, err)
	fmt.Println(count)

}
