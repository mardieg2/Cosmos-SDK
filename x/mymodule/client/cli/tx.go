package cli

import (
	"github.com/spf13/cobra"
)

func NewTxCmd() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "mymodule",
        Short: "My module transactions subcommands",
    }

    cmd.AddCommand(
        NewCreateSomethingCmd(),
    )

    return cmd
}

func NewCreateSomethingCmd() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "create-something [args]",
        Short: "Create something",
        RunE: func(cmd *cobra.Command, args []string) error {
            // Implement your command logic here
            return nil
        },
    }

    return cmd
}
