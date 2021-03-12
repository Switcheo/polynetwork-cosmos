package main

import (
	"os"

	"github.com/Switcheo/polynetwork-cosmos/app"
	"github.com/Switcheo/polynetwork-cosmos/cmd/polynetwork-cosmosd/cmd"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
