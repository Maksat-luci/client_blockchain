package main

import (
	"fmt"
	// "github.com/cosmos/cosmos-sdk/client"
	// "github.com/cometbft/cometbft/test/e2e/app"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/server"
	// "github.com/cosmos/cosmos-sdk/simapp"

	"cosmossdk.io/simapp"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	simcli "github.com/cosmos/cosmos-sdk/x/simulation/client/cli"
)

func main() {
	fmt.Println(sendTx())
}

var (
	DefaultNodeHome = "/home/max/golang/client"
)

func sendTx() string {
	config := simcli.NewConfigFromFlags()
	config.ChainID = "mdc"
	db, _, logger, _, err := simtestutil.SetupSimulation(config, "leveldb-app-sim", "Simulation", simcli.FlagVerboseValue, simcli.FlagEnabledValue)
	if err != nil {
		return err.Error()
	}
	appOptions := make(simtestutil.AppOptionsMap, 0)
	appOptions[flags.FlagHome] = DefaultNodeHome
	appOptions[server.FlagInvCheckPeriod] = simcli.FlagPeriodValue
	app := simapp.NewSimApp(logger, db, nil, true, appOptions, fauxMerkleModeOpt)
	// txBuilder := app.
	_ = app
	return ""
}

func fauxMerkleModeOpt(bapp *baseapp.BaseApp) {
	bapp.SetFauxMerkleMode()
}

// func sendTx() string {
// 	ctx := context.Background()
// 	addressPrefix := "mdc"
// 	client, err := cosmosclient.New(ctx, cosmosclient.WithAddressPrefix(addressPrefix))
// 	if err != nil {
// 		return err.Error()
// 	}
// 	acc, err := client.Account("mdc12la80nam767spymufyurufl702wk0kl7yplef0")
// 	if err != nil {
// 		return err.Error()
// 	}
// 	acc2, err := client.Account("mdc12gm7lw872wvz00h7cws26k3ndd4dsa9yzvx54y")
// 	if err != nil {
// 		return err.Error()
// 	}

// 	addr1, err :=  acc.Address(addressPrefix)
// 	if err != nil {
// 		return err.Error()
// 	}
// 	addr2, err := acc2.Address(addressPrefix)
// 	if err != nil {
// 		return err.Error()
// 	}
// 	msg := &types.MsgSend{
// 		FromAddress:addr1,
// 		ToAddress: addr2,
// 		Amount: sdk.NewCoins(sdk.NewInt64Coin("stake", 99)),
// 	}

// 	txResp, err := client.BroadcastTx(ctx, acc, msg)
// 	if err != nil {
// 		return err.Error()
// 	}
// 	resp := "MsgCreatePost:\n\n" + txResp.Data + string(txResp.Code) + " " + txResp.Data

// 	return resp
// }
