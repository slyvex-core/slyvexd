package rpchandlers

import (
	"github.com/slyvex-core/slyvexd/app/appmessage"
	"github.com/slyvex-core/slyvexd/app/rpc/rpccontext"
	"github.com/slyvex-core/slyvexd/domain/consensus/utils/constants"
	"github.com/slyvex-core/slyvexd/infrastructure/network/netadapter/router"
)

// HandleGetCoinSupply handles the respectively named RPC command
func HandleGetCoinSupply(context *rpccontext.Context, _ *router.Router, _ appmessage.Message) (appmessage.Message, error) {
	if !context.Config.UTXOIndex {
		errorMessage := &appmessage.GetCoinSupplyResponseMessage{}
		errorMessage.Error = appmessage.RPCErrorf("Method unavailable when slyvexd is run without --utxoindex")
		return errorMessage, nil
	}

	circulatingSeepSupply, err := context.UTXOIndex.GetCirculatingSeepSupply()
	if err != nil {
		return nil, err
	}

	response := appmessage.NewGetCoinSupplyResponseMessage(
		constants.MaxSeep,
		circulatingSeepSupply,
	)

	return response, nil
}
