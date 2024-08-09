package contract

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ethereum-optimism/optimism/op-challenger/game/fault/types"
	"github.com/ethereum-optimism/optimism/op-service/client"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum-optimism/optimism/op-service/predeploys"
	"github.com/ethereum-optimism/optimism/op-service/sources"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
	"math/big"
	"strings"
	"testing"
)

func TestContract(t *testing.T) {
	l1rpc, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/lV2e-64nNnEMUA7UG0IT0uwjzlxEI512")
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

func TestBlockRange(t *testing.T) {
	l1rpc, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/PNunSRFo0FWRJMu5yrwBd6jF7G78YHrv")
	require.NoError(t, err)
	disputeGame := "0xfe90247bb9d191a6F3840815f34F61E9d8Df89ee"
	newDisputeGame, err := NewDisputeGame(common.HexToAddress(disputeGame), l1rpc)
	////startingBlockNumber  prestateBlock
	////l2BlockNumber        poststateBlock
	//
	prestateBlock, err := newDisputeGame.StartingBlockNumber(&bind.CallOpts{})
	require.NoError(t, err)
	poststateBlock, err := newDisputeGame.L2BlockNumber(&bind.CallOpts{})
	require.NoError(t, err)
	splitDepth, err := newDisputeGame.SplitDepth(&bind.CallOpts{})
	require.NoError(t, err)
	splitDepths := types.Depth(splitDepth.Uint64())

	pos := types.NewPositionFromGIndex(big.NewInt(1))
	traceIndex := pos.TraceIndex(splitDepths)
	fmt.Printf("traceIndex:%s\n", traceIndex)
	if !traceIndex.IsUint64() {
		fmt.Errorf("err:%s", traceIndex)
	}
	outputBlock := traceIndex.Uint64() + prestateBlock.Uint64() + 1
	if outputBlock > poststateBlock.Uint64() {
		outputBlock = poststateBlock.Uint64()
	}
	fmt.Printf("outputblock:%d\n", outputBlock)
	fmt.Printf("blockhash:%s", hexutil.Uint64(outputBlock))

	l2rpc, err := ethclient.Dial("https://opt-sepolia.g.alchemy.com/v2/FPgbOkDCgG8t0ppZ6TwZXLucr1wl_us4")
	require.NoError(t, err)
	defer l2rpc.Close()
	l2RPC := client.NewBaseRPCClient(l2rpc.Client())

	rollupClient := sources.NewRollupClient(l2RPC)
	output, err := rollupClient.OutputAtBlock(context.Background(), outputBlock)
	fmt.Printf("outputRoot:%s\n", common.Hash(output.OutputRoot))
}

//blockHash 0x0e494f1663e2e1b876f706668f1abebd762341aaefd9e7463cb3a109383a6f5b
//stateRoot 0xe67c67ddb0ac98bc8c395d782fc32bdcdc4590b93da6abef8da88ef9f62050c2
//storageHash 0x888e7e703509255745ed639a98c7dd1c8c84c98fae3a884c642b0343fbb69b3c

func TestHash(t *testing.T) {
	l2rpc, err := ethclient.Dial("https://opt-sepolia.g.alchemy.com/v2/FPgbOkDCgG8t0ppZ6TwZXLucr1wl_us4")
	require.NoError(t, err)
	defer l2rpc.Close()

	block, err := l2rpc.BlockByNumber(context.Background(), big.NewInt(12827274))
	require.NoError(t, err)
	l2RPC := client.NewBaseRPCClient(l2rpc.Client())

	var getProofResponse *eth.AccountResult
	err = l2RPC.CallContext(context.Background(), &getProofResponse, "eth_getProof", predeploys.L2ToL1MessagePasserAddr, []common.Hash{}, block.Hash().String())
	require.NoError(t, err)

	output := &eth.OutputV0{
		StateRoot:                eth.Bytes32(block.Root()),
		MessagePasserStorageRoot: eth.Bytes32(getProofResponse.StorageHash),
		BlockHash:                block.Hash(),
	}

	//output := &eth.OutputV0{
	//	StateRoot:                eth.Bytes32(common.HexToHash("0xe67c67ddb0ac98bc8c395d782fc32bdcdc4590b93da6abef8da88ef9f62050c2")),
	//	MessagePasserStorageRoot: eth.Bytes32(common.HexToHash("0x888e7e703509255745ed639a98c7dd1c8c84c98fae3a884c642b0343fbb69b3c")),
	//	BlockHash:                common.HexToHash("0x0e494f1663e2e1b876f706668f1abebd762341aaefd9e7463cb3a109383a6f5b"),
	//}
	outputRoot := eth.OutputRoot(output)
	fmt.Println(outputRoot)
}
