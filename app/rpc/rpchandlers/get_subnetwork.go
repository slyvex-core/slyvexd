package rpchandlers

import (
	"github.com/slyvex-core/slyvexd/app/appmessage"
	"github.com/slyvex-core/slyvexd/app/rpc/rpccontext"
	"github.com/slyvex-core/slyvexd/infrastructure/network/netadapter/router"
)

// HandleGetSubnetwork handles the respectively named RPC command
func HandleGetSubnetwork(context *rpccontext.Context, _ *router.Router, request appmessage.Message) (appmessage.Message, error) {
	response := &appmessage.GetSubnetworkResponseMessage{}
	response.Error = appmessage.RPCErrorf("not implemented")
	return response, nil
}
