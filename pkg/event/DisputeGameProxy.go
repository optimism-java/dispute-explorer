package event

import (
	"encoding/json"
	"github.com/spf13/cast"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	DisputeGameCreatedName = "DisputeGameCreated"

	DisputeGameCreatedHash = crypto.Keccak256([]byte("DisputeGameCreated(address,uint32,bytes32)"))
)

type DisputeGameCreated struct {
	DisputeProxy string `json:"disputeProxy"`
	GameType     uint32 `json:"gameType"`
	RootClaim    string `json:"rootClaim"`
}

func (*DisputeGameCreated) Name() string {
	return DisputeGameCreatedName
}

func (*DisputeGameCreated) EventHash() common.Hash {
	return common.BytesToHash(DisputeGameCreatedHash)
}

func (t *DisputeGameCreated) ToObj(data string) error {
	err := json.Unmarshal([]byte(data), &t)
	if err != nil {
		return err
	}
	return nil
}

func (*DisputeGameCreated) Data(log types.Log) (string, error) {
	transfer := &DisputeGameCreated{
		DisputeProxy: TopicToAddress(log, 1).Hex(),
		GameType:     cast.ToUint32(TopicToInt64(log, 2)),
		RootClaim:    TopicToHash(log, 3).Hex(),
	}
	data, err := ToJSON(transfer)
	if err != nil {
		return "", err
	}
	return data, nil
}
