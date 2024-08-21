package contract

import (
	"context"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum-optimism/optimism/op-service/client"
	"github.com/ethereum-optimism/optimism/op-service/sources"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
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

// Loading addresses from /workspace/optimism/packages/contracts-bedrock/deployments/3151908/kurtosis.json
// Saving AddressManager: 0x53Be5dc06De227757056873E6D3D66c1E3E5cE92
// Saving AnchorStateRegistry: 0x7681f947F40852C973F4360E8Ba93D78Fc18827D
// Saving AnchorStateRegistryProxy: 0xf130D1DBDB5AF7224781fDeE9728E2E3EE3f4af9
// Saving DelayedWETH: 0xeB09769574b8029EcA7333eD9025406CA10F20E1
// Saving DelayedWETHProxy: 0xdaA56Bfb3bB00Ab2aE3f37A77175fA876d752e37
// Saving DisputeGameFactory: 0x368C112Ab442e01F4654b4C8F9D676eCfd7Bee84
// Saving DisputeGameFactoryProxy: 0xf149eEe2E3F33d053e5F7C8F4eB3a2fF14B8b202
// Saving L1CrossDomainMessenger: 0xb4752935558c688cbFd066d1d534Acf81065f4Ad
// Saving L1CrossDomainMessengerProxy: 0x5c00fe7f934622f6ED2Ad7D7a972dEec59E122F7
// Saving L1ERC721Bridge: 0xe048EBA8aD6F52A58dd8F0A4CF2E638320358971
// Saving L1ERC721BridgeProxy: 0x5c60C1Ddae2fc7530581d84e6CB9894d7860BcF2
// Saving L1StandardBridge: 0xf2607e782CA09d8721840597e734adcEf549A20A
// Saving L1StandardBridgeProxy: 0x92D2037D5a6a94D0EcE1E82D5304b2727B5093e1
// Saving L2OutputOracle: 0xcb0D8F56920537E5b1387cA45E0A7EE64635f168
// Saving L2OutputOracleProxy: 0x579d4D43EC93fbB5f434319AB70174d300015414
// Saving Mips: 0xDC6F8c922f6B8Fc0e1c9Db3a50EA3a0be353A5fF
// Saving OptimismMintableERC20Factory: 0x986141DB5b1291EE646FE894B881702F21057e60
// Saving OptimismMintableERC20FactoryProxy: 0xa0BeD97c83A8032f05b70C8fd9242972471a9Cd3
// Saving OptimismPortal: 0xce01807a35abD57CC6Ac670e6383220EE832c290
// Saving OptimismPortal2: 0x35F9bAb4ec7c21a06A3A78C3Fe7fD0F4F1698B5C
// Saving OptimismPortalProxy: 0x49508B8f34a23cEFFC728EF8Ed00245820Dd2f2E
// Saving PreimageOracle: 0x9D6c1A3669Ab91FA6f578896950A827B183Cc84b
// Saving ProtocolVersions: 0xB9fc75b3448561785d96be4E96d1f56Ad58A8031
// Saving ProtocolVersionsProxy: 0x971f488763D9129510B01bE514BD64AC193897EB
// Saving ProxyAdmin: 0xDb50B44df5Ee921A869f11F829F8ED45746376Ef
// Saving SafeProxyFactory: 0x172B4d2CA55C12659664129d297FA6aA060Ff793
// Saving SafeSingleton: 0x595d35Eec105ba9674B7392c984D86b6F25760ec
// Saving SuperchainConfig: 0xD025a2E1818b20454d3f3628C3228f3c2ee95675
// Saving SuperchainConfigProxy: 0x3DF4Dc2cAFF6381eF943815a0397D3933232eA81
// Saving SystemConfig: 0xEfd0526348baE10523A8a6EC1F417756cA125342
// Saving SystemConfigProxy: 0x1693988fEfa147A321F4989d54f2B6CDe8C4DA33
// Saving SystemOwnerSafe: 0xde99354AA68b841d52aBea4818c3bC3B83A07b1E
func TestCreateGames(t *testing.T) {
	privateKey, err := crypto.HexToECDSA("bcdf20249abf0ed6d944c0288fad489e33f66b3960d9e6229c1cd214ed3bbe31")
	require.NoError(t, err)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(3151908))
	require.NoError(t, err)
	fmt.Println(auth.Signer)
	l1rpc, err := ethclient.Dial("http://localhost:32773/")
	require.NoError(t, err)
	defer l1rpc.Close()
	l2rpc, err := ethclient.Dial("http://localhost:32788/")
	require.NoError(t, err)
	defer l2rpc.Close()
	l2RPC := client.NewBaseRPCClient(l2rpc.Client())
	// DisputeGameFactoryProxy: 0xf149eEe2E3F33d053e5F7C8F4eB3a2fF14B8b202
	rollupClient := sources.NewRollupClient(l2RPC)

	anchorProxy, err := NewAnchorStateProxy(common.HexToAddress("0x0d65D421A0768C94F09e56F98eE2017DF2761786"), l1rpc)
	res, err := anchorProxy.Anchors(&bind.CallOpts{}, 0)
	fmt.Println(res.L2BlockNumber)
	fmt.Println(hex.EncodeToString(res.Root[:]))

	a, err := anchorProxy.DisputeGameFactory(&bind.CallOpts{})
	fmt.Printf("game:%s", a)

	disputeGameFactory, err := NewDisputeGameProxy(common.HexToAddress("0x8d8C34f63e9812d862CE46AC302ec0D7F72742B6"), l1rpc)
	require.NoError(t, err)
	count, err := disputeGameFactory.GameCount(&bind.CallOpts{})
	fmt.Printf("count: %s\n", count)
	initBonds, err := disputeGameFactory.InitBonds(&bind.CallOpts{}, 0)
	fmt.Printf("bonds:%s\n", initBonds.String()) // 0xc000622000
	owner, err := disputeGameFactory.Owner(&bind.CallOpts{})
	fmt.Printf("owner:%s\n", owner) // 0xc000622000
	imples, err := disputeGameFactory.GameImpls(&bind.CallOpts{}, 0)
	fmt.Printf("imple: %s\n", imples)
	l2BlockNumber := uint64(800)
	output, err := rollupClient.OutputAtBlock(context.Background(), l2BlockNumber)
	fmt.Printf("output:%s\n", common.Hash(output.OutputRoot))
	//
	extraData := make([]byte, 32)
	binary.BigEndian.PutUint64(extraData[24:], l2BlockNumber)
	tx, err := disputeGameFactory.Create(&bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
		GasPrice: big.NewInt(1000000007),
		GasLimit: uint64(299999),
		Value:    initBonds,
	}, uint32(0), common.Hash(output.OutputRoot), extraData)
	fmt.Println(tx)
	//
	//fmt.Printf("extraData:%s\n", hex.EncodeToString(extraData[:]))
	//
	//initBonds1, err := disputeGameFactory.InitBonds(&bind.CallOpts{}, 1)
	//fmt.Println(initBonds1.String()) //0xc000622000

	//tx, err := disputeGameFactory.SetInitBond(&bind.TransactOpts{
	//	From:   auth.From,
	//	Signer: auth.Signer,
	//}, 0, big.NewInt(10000))
	//require.NoError(t, err)
	//fmt.Println(tx)

	//tx, err := transactions.PadGasEstimate(&bind.TransactOpts{
	//	From:   auth.From,
	//	Signer: auth.Signer,
	//}, 2, func(opts *bind.TransactOpts) (*types.Transaction, error) {
	//	return disputeGameFactory.Create(opts, 0, common.Hash(output.OutputRoot), extraData)
	//})

	//if err != nil {
	//	fmt.Println(errors.WithStack(err))
	//}
	//fmt.Println(tx)
}

func TestDecode(t *testing.T) {
	Method := crypto.Keccak256([]byte("AlreadyInitialized()"))[:4]
	fmt.Printf("Method: %x\n", Method)
	Method2 := crypto.Keccak256([]byte("AnchorRootNotFound()"))[:4]
	fmt.Printf("Method: %x\n", Method2)
	Method3 := crypto.Keccak256([]byte("setImplementation(uint32,address)"))[:4]
	fmt.Printf("Method: %x\n", Method3)
	// 0x14f6b1a30000000000000000000000000000000000000000000000000000000000000000 000000000000000000000000b926ffb5a8bba2fe9cea2c888228176ca683301b
	// 0x14f6b1a30000000000000000000000000000000000000000000000000000000000000001 000000000000000000000000d41afd0c05272d6b3586b542b53856bd183effcb
	// resolve()
	Method4 := crypto.Keccak256([]byte("resolve()"))[:4]
	fmt.Printf("Method: %x\n", Method4)
	Method5 := crypto.Keccak256([]byte("setInitBond(uint32,address)"))[:4]
	fmt.Printf("Method: %x\n", Method5)
	Method6 := crypto.Keccak256([]byte("create(uint32,byte32,bytes)"))[:4]
	fmt.Printf("Method: %x\n", Method6)
	Method7 := crypto.Keccak256([]byte("transferOwnership(address)"))[:4]
	fmt.Printf("Method: %x\n", Method7)
}

func TestNewDisputeGame(t *testing.T) {
	privateKey, err := crypto.HexToECDSA("bcdf20249abf0ed6d944c0288fad489e33f66b3960d9e6229c1cd214ed3bbe31")
	require.NoError(t, err)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(3151908))
	require.NoError(t, err)
	fmt.Println(auth.From)
	fmt.Println(auth.Signer)

	l1rpc, err := ethclient.Dial("http://localhost:32773")
	require.NoError(t, err)
	disputeGame, err := NewDisputeGame(common.HexToAddress("0xb926ffb5a8bba2fe9cea2c888228176ca683301b"), l1rpc)
	require.NoError(t, err)
	count, err := disputeGame.ClaimDataLen(&bind.CallOpts{})
	require.NoError(t, err)
	fmt.Println(count)

	status, err := disputeGame.Status(&bind.CallOpts{})
	fmt.Println(status)

	gameData, err := disputeGame.GameData(&bind.CallOpts{})
	fmt.Println(gameData.GameType)
	fmt.Println(gameData.ExtraData)
	fmt.Println(gameData.RootClaim)
}
