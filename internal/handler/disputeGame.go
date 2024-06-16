package handler

import (
	"github.com/optimism-java/dispute-explorer/internal/blockchain"
	"github.com/optimism-java/dispute-explorer/internal/schema"
	"github.com/optimism-java/dispute-explorer/pkg/event"
	"github.com/optimism-java/dispute-explorer/pkg/log"
	"github.com/pkg/errors"
)

func FilterAddAndRemove(evt *schema.SyncEvent) error {
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

func BatchFilterAddAndRemove(events []*schema.SyncEvent) error {
	for _, evt := range events {
		err := FilterAddAndRemove(evt)
		if err != nil {
			log.Errorf("[BatchFilterAddAndRemove] FilterAddAndRemove: %s\n", err)
			return err
		}
	}
	return nil
}
