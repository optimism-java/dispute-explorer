package handler

import (
	"context"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/optimism-java/dispute-explorer/internal/blockchain"
	"github.com/optimism-java/dispute-explorer/internal/schema"
	"github.com/optimism-java/dispute-explorer/internal/svc"
	"github.com/optimism-java/dispute-explorer/pkg/contract"
	"github.com/optimism-java/dispute-explorer/pkg/event"
	"github.com/optimism-java/dispute-explorer/pkg/log"
	"github.com/pkg/errors"
)

func FilterAddAndRemove(ctx *svc.ServiceContext, evt *schema.SyncEvent) error {
	dispute := event.DisputeGameCreated{}
	if evt.EventName == dispute.Name() && evt.EventHash == dispute.EventHash().String() {
		err := dispute.ToObj(evt.Data)
		if err != nil {
			log.Errorf("[FilterDisputeContractAndAdd] event data to DisputeGameCreated err: %s\n", err)
			return errors.WithStack(err)
		}
		blockchain.AddContract(dispute.DisputeProxy)
	}
	disputeResolved := event.DisputeGameResolved{}
	if evt.EventName == disputeResolved.Name() && evt.EventHash == disputeResolved.EventHash().String() {
		err := disputeResolved.ToObj(evt.Data)
		if err != nil {
			log.Errorf("[FilterDisputeContractAndAdd] event data to DisputeGameResolved err: %s\n", err)
			return errors.WithStack(err)
		}
		blockchain.RemoveContract(evt.ContractAddress)
	}
	return nil
}

func BatchFilterAddAndRemove(ctx *svc.ServiceContext, events []*schema.SyncEvent) error {
	for _, evt := range events {
		err := FilterAddAndRemove(ctx, evt)
		if err != nil {
			log.Errorf("[BatchFilterAddAndRemove] FilterAddAndRemove: %s\n", err)
			return err
		}
	}
	return nil
}

func getGameInput(ctx *svc.ServiceContext, txHash string) (map[uint64]map[string]interface{}, error) {
	abiObject, err := abi.JSON(strings.NewReader(contract.DisputeGameProxyMetaData.ABI))
	if err != nil {
		return nil, fmt.Errorf("[getGameInput] parse abi error: %s", errors.WithStack(err))
	}
	tx, _, err := ctx.L1RPC.TransactionByHash(context.Background(), common.HexToHash(txHash))
	if err != nil {
		return nil, fmt.Errorf("[getGameInput] get tx error: %s", errors.WithStack(err))
	}
	inputsMap, methodName := decodeTransactionInputData(abiObject, tx.Data())

	if methodName != "create" {
		return nil, fmt.Errorf("[getGameInput] methodName is :  %s parse method error: %s", methodName, errors.WithStack(err))
	}

	e := inputsMap["_gameType"].(uint32)
	f := inputsMap["_rootClaim"].([32]byte)
	g := inputsMap["_extraData"].([]byte)

	fmt.Println("_gameType", e)
	fmt.Println("_rootClaim:", hex.EncodeToString(f[:]))
	fmt.Println("_extraData: ", hex.EncodeToString(g))
	return nil, nil
}

func decodeTransactionInputData(contractABI abi.ABI, data []byte) (map[string]interface{}, string) {
	methodSigData := data[:4]
	inputsSigData := data[4:]

	method, err := contractABI.MethodById(methodSigData)
	if err != nil {
		log.Errorf("[DecodeTransactionInputData] parse abi error: %s\n", errors.WithStack(err))
	}

	inputsMap := make(map[string]interface{})

	if err := method.Inputs.UnpackIntoMap(inputsMap, inputsSigData); err != nil {
		log.Errorf("[DecodeTransactionInputData] parse abi error: %s\n", errors.WithStack(err))
	}
	return inputsMap, method.Name
}
