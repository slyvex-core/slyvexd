package rpchandlers

import (
	"github.com/slyvex-core/slyvexd/app/appmessage"
	"github.com/slyvex-core/slyvexd/app/rpc/rpccontext"
	"github.com/slyvex-core/slyvexd/domain/consensus/utils/consensushashing"
	"github.com/slyvex-core/slyvexd/domain/miningmanager/mempool"
	"github.com/slyvex-core/slyvexd/infrastructure/network/netadapter/router"
	"github.com/pkg/errors"
)

// HandleSubmitTransaction handles the respectively named RPC command
func HandleSubmitTransaction(context *rpccontext.Context, _ *router.Router, request appmessage.Message) (appmessage.Message, error) {
	submitTransactionRequest := request.(*appmessage.SubmitTransactionRequestMessage)

	domainTransaction, err := appmessage.RPCTransactionToDomainTransaction(submitTransactionRequest.Transaction)
	if err != nil {
		errorMessage := &appmessage.SubmitTransactionResponseMessage{}
		errorMessage.Error = appmessage.RPCErrorf("Could not parse transaction: %s", err)
		return errorMessage, nil
	}

	transactionID := consensushashing.TransactionID(domainTransaction)
	err = context.ProtocolManager.AddTransaction(domainTransaction, submitTransactionRequest.AllowOrphan)
	if err != nil {
		if !errors.As(err, &mempool.RuleError{}) {
			return nil, err
		}

		log.Debugf("Rejected transaction %s: %s", transactionID, err)
		errorMessage := &appmessage.SubmitTransactionResponseMessage{}
		errorMessage.Error = appmessage.RPCErrorf("Rejected transaction %s: %s", transactionID, err)
		return errorMessage, nil
	}

	response := appmessage.NewSubmitTransactionResponseMessage(transactionID.String())
	return response, nil
}
