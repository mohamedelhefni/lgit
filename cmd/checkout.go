package cmd

import (
	"lgit/base"
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(checkoutCmd)
}

func checkout(oid string) error {
	commit, err := base.GetCommit(oid)
	if err != nil {
		return err
	}
	err = base.ReadTree(commit.Tree)
	if err != nil {
		return err
	}
	return base.SetHead(oid)
}

var checkoutCmd = &cobra.Command{
	Use: "checkout",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("file args is required")
		}
		err := checkout(args[0])
		if err != nil {
			log.Fatal(err)
		}

	},
}
