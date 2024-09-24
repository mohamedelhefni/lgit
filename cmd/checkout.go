package cmd

import (
	"lgit/base"
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(checkoutCmd)
}

func checkout(name string) error {
	oid := base.GetOID(name)
	commit, err := base.GetCommit(oid)
	if err != nil {
		return err
	}

	var HEAD base.RefValue
	if base.IsBranch(name) {
		HEAD = base.RefValue{Symbolic: true, Value: "refs/heads/" + name}
	} else {
		HEAD = base.RefValue{Symbolic: false, Value: oid}
	}

	err = base.ReadTree(commit.Tree)
	if err != nil {
		return err
	}
	return base.SetRef("HEAD", HEAD, false)
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
