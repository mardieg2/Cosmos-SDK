package main

import (
	"os"

	"github.com/cosmos/cosmos-sdk/client"
)

func main() {
    rootCmd, _ := client.NewRootCmd()
    if err := rootCmd.Execute(); err != nil {
        os.Exit(1)
    }
}
