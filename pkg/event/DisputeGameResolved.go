package event

import (
	"encoding/json"

	"github.com/spf13/cast"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	DisputeGameResolvedName = "Resolved"

	DisputeGameResolvedHash = crypto.Keccak256([]byte("Resolved(uint8)"))
)

type DisputeGameResolved struct {
	Status uint8 `json:"status"`
}

func (*DisputeGameResolved) Name() string {
	return DisputeGameResolvedName
}

func (*DisputeGameResolved) EventHash() common.Hash {
	return common.BytesToHash(DisputeGameResolvedHash)
}

func (t *DisputeGameResolved) ToObj(data string) error {
	err := json.Unmarshal([]byte(data), &t)
	if err != nil {
		return err
	}
	return nil
}

func (*DisputeGameResolved) Data(log types.Log) (string, error) {
	transfer := &DisputeGameResolved{
		Status: cast.ToUint8(TopicToInt64(log, 1)),
	}
	data, err := ToJSON(transfer)
	if err != nil {
		return "", err
	}
	return data, nil
}
