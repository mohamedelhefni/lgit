package cmd

import (
	"lgit/base"
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(showCmd)
}

func show(oid string) error {
	commit, err := base.GetCommit(oid)
	if err != nil {
		return err
	}
	printCommit(CommitArgs{
		Oid:    oid,
		Commit: commit,
	})
	return nil
}

var showCmd = &cobra.Command{
	Use: "show",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("file args is required")
		}
		err := show(base.GetOID(args[0]))
		if err != nil {
			log.Fatal(err)
		}

	},
}
