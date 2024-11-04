package main

import (
	"os"

	"github.com/cosmos/cosmos-sdk/server"
)

func main() {
    rootCmd, _ := server.NewRootCmd()
    if err := rootCmd.Execute(); err != nil {
        os.Exit(1)
    }
}
