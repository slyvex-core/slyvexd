package main

import (
	"context"
	"fmt"

	"github.com/slyvex-core/slyvexd/cmd/slyvexwallet/daemon/client"
	"github.com/slyvex-core/slyvexd/cmd/slyvexwallet/daemon/pb"
	"github.com/slyvex-core/slyvexd/cmd/slyvexwallet/utils"
)

func balance(conf *balanceConfig) error {
	daemonClient, tearDown, err := client.Connect(conf.DaemonAddress)
	if err != nil {
		return err
	}
	defer tearDown()

	ctx, cancel := context.WithTimeout(context.Background(), daemonTimeout)
	defer cancel()
	response, err := daemonClient.GetBalance(ctx, &pb.GetBalanceRequest{})
	if err != nil {
		return err
	}

	pendingSuffix := ""
	if response.Pending > 0 {
		pendingSuffix = " (pending)"
	}
	if conf.Verbose {
		pendingSuffix = ""
		println("Address                                                                       Available             Pending")
		println("-----------------------------------------------------------------------------------------------------------")
		for _, addressBalance := range response.AddressBalances {
			fmt.Printf("%s %s %s\n", addressBalance.Address, utils.FormatSvx(addressBalance.Available), utils.FormatSvx(addressBalance.Pending))
		}
		println("-----------------------------------------------------------------------------------------------------------")
		print("                                                 ")
	}
	fmt.Printf("Total balance, SVX %s %s%s\n", utils.FormatSvx(response.Available), utils.FormatSvx(response.Pending), pendingSuffix)

	return nil
}
