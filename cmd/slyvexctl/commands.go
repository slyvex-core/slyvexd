package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/slyvex-core/slyvexd/infrastructure/network/netadapter/server/grpcserver/protowire"
)

var commandTypes = []reflect.Type{
	reflect.TypeOf(protowire.SlyvexdMessage_AddPeerRequest{}),
	reflect.TypeOf(protowire.SlyvexdMessage_GetConnectedPeerInfoRequest{}),
	reflect.TypeOf(protowire.SlyvexdMessage_GetPeerAddressesRequest{}),
	reflect.TypeOf(protowire.SlyvexdMessage_GetCurrentNetworkRequest{}),
	reflect.TypeOf(protowire.SlyvexdMessage_GetInfoRequest{}),

	reflect.TypeOf(protowire.SlyvexdMessage_GetBlockRequest{}),
	reflect.TypeOf(protowire.SlyvexdMessage_GetBlocksRequest{}),
	reflect.TypeOf(protowire.SlyvexdMessage_GetHeadersRequest{}),
	reflect.TypeOf(protowire.SlyvexdMessage_GetBlockCountRequest{}),
	reflect.TypeOf(protowire.SlyvexdMessage_GetBlockDagInfoRequest{}),
	reflect.TypeOf(protowire.SlyvexdMessage_GetSelectedTipHashRequest{}),
	reflect.TypeOf(protowire.SlyvexdMessage_GetVirtualSelectedParentBlueScoreRequest{}),
	reflect.TypeOf(protowire.SlyvexdMessage_GetVirtualSelectedParentChainFromBlockRequest{}),
	reflect.TypeOf(protowire.SlyvexdMessage_ResolveFinalityConflictRequest{}),
	reflect.TypeOf(protowire.SlyvexdMessage_EstimateNetworkHashesPerSecondRequest{}),

	reflect.TypeOf(protowire.SlyvexdMessage_GetBlockTemplateRequest{}),
	reflect.TypeOf(protowire.SlyvexdMessage_SubmitBlockRequest{}),

	reflect.TypeOf(protowire.SlyvexdMessage_GetMempoolEntryRequest{}),
	reflect.TypeOf(protowire.SlyvexdMessage_GetMempoolEntriesRequest{}),
	reflect.TypeOf(protowire.SlyvexdMessage_GetMempoolEntriesByAddressesRequest{}),

	reflect.TypeOf(protowire.SlyvexdMessage_SubmitTransactionRequest{}),

	reflect.TypeOf(protowire.SlyvexdMessage_GetUtxosByAddressesRequest{}),
	reflect.TypeOf(protowire.SlyvexdMessage_GetBalanceByAddressRequest{}),
	reflect.TypeOf(protowire.SlyvexdMessage_GetCoinSupplyRequest{}),

	reflect.TypeOf(protowire.SlyvexdMessage_BanRequest{}),
	reflect.TypeOf(protowire.SlyvexdMessage_UnbanRequest{}),
}

type commandDescription struct {
	name       string
	parameters []*parameterDescription
	typeof     reflect.Type
}

type parameterDescription struct {
	name   string
	typeof reflect.Type
}

func commandDescriptions() []*commandDescription {
	commandDescriptions := make([]*commandDescription, len(commandTypes))

	for i, commandTypeWrapped := range commandTypes {
		commandType := unwrapCommandType(commandTypeWrapped)

		name := strings.TrimSuffix(commandType.Name(), "RequestMessage")
		numFields := commandType.NumField()

		var parameters []*parameterDescription
		for i := 0; i < numFields; i++ {
			field := commandType.Field(i)

			if !isFieldExported(field) {
				continue
			}

			parameters = append(parameters, &parameterDescription{
				name:   field.Name,
				typeof: field.Type,
			})
		}
		commandDescriptions[i] = &commandDescription{
			name:       name,
			parameters: parameters,
			typeof:     commandTypeWrapped,
		}
	}

	return commandDescriptions
}

func (cd *commandDescription) help() string {
	sb := &strings.Builder{}
	sb.WriteString(cd.name)
	for _, parameter := range cd.parameters {
		_, _ = fmt.Fprintf(sb, " [%s]", parameter.name)
	}
	return sb.String()
}
