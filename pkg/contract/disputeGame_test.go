package contract

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
	"math/big"
	"strings"
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

func TestCallL2BlockNumber(t *testing.T) {
	l1rpc, err := ethclient.Dial("https://quaint-white-season.ethereum-sepolia.quiknode.pro/b5c30cbb548d8743f08dd175fe50e3e923259d30")
	require.NoError(t, err)
	parsed, err := abi.JSON(strings.NewReader(DisputeGameMetaData.ABI))
	callData, err := parsed.Pack("l2BlockNumber")
	adr := common.HexToAddress("0x2ab2Ed3ce15de144432B62c9F03e435B8bB513d0")
	result, _ := l1rpc.CallContract(context.Background(), ethereum.CallMsg{
		To:   &adr,
		Data: callData,
	}, nil)
	fmt.Println(hex.EncodeToString(result))
	integer := new(big.Int)
	a, ok := integer.SetString(hex.EncodeToString(result), 16)
	fmt.Println(ok)
	fmt.Println(a)
}

func TestCallStatus(t *testing.T) {
	l1rpc, err := ethclient.Dial("https://quaint-white-season.ethereum-sepolia.quiknode.pro/b5c30cbb548d8743f08dd175fe50e3e923259d30")
	require.NoError(t, err)
	parsed, err := abi.JSON(strings.NewReader(DisputeGameMetaData.ABI))
	callData, err := parsed.Pack("status")
	adr := common.HexToAddress("0x2ab2Ed3ce15de144432B62c9F03e435B8bB513d0")
	result, _ := l1rpc.CallContract(context.Background(), ethereum.CallMsg{
		To:   &adr,
		Data: callData,
	}, nil)
	fmt.Println(hex.EncodeToString(result))
}

func TestAllCredit(t *testing.T) {
	l1rpc, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/PNunSRFo0FWRJMu5yrwBd6jF7G78YHrv")
	require.NoError(t, err)
	disputeGame := "0xc9cb084c3ad4e36b719b60649f99ea9f13bb45b7"
	newDisputeGame, err := NewDisputeGame(common.HexToAddress(disputeGame), l1rpc)
	credit, err := newDisputeGame.Credit(&bind.CallOpts{}, common.HexToAddress("0x49277EE36A024120Ee218127354c4a3591dc90A9"))
	println(credit.Int64())
}
