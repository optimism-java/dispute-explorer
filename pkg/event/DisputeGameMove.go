package event

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	DisputeGameMoveName = "Move"

	DisputeGameMovedHash = crypto.Keccak256([]byte("Move(uint256,bytes32,address)"))
)

type DisputeGameMove struct {
	ParentIndex *big.Int `json:"parentIndex"`
	Claim       string   `json:"claim"`
	Claimant    string   `json:"claimant"`
}

func (*DisputeGameMove) Name() string {
	return DisputeGameMoveName
}

func (*DisputeGameMove) EventHash() common.Hash {
	return common.BytesToHash(DisputeGameMovedHash)
}

func (t *DisputeGameMove) ToObj(data string) error {
	err := json.Unmarshal([]byte(data), &t)
	if err != nil {
		return err
	}
	return nil
}

func (*DisputeGameMove) Data(log types.Log) (string, error) {
	transfer := &DisputeGameMove{
		ParentIndex: big.NewInt(TopicToInt64(log, 1)),
		Claim:       TopicToHash(log, 2).Hex(),
		Claimant:    TopicToAddress(log, 3).Hex(),
	}
	data, err := ToJSON(transfer)
	if err != nil {
		return "", err
	}
	return data, nil
}
