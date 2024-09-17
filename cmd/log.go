package cmd

import (
	"fmt"
	"lgit/base"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(logCmd)
}

func showCommits(argOid string) error {
	iter, err := base.IterCommits([]string{argOid})
	if err != nil {
		return err
	}

	for oid := range iter {
		commit, err := base.GetCommit(oid)
		if err != nil {
			fmt.Println("err: ", err)
			continue
		}
		fmt.Println("commit: ", oid)
		fmt.Println(strings.TrimSpace(commit.Message))
	}

	return nil
}

var logCmd = &cobra.Command{
	Use:   "log",
	Short: "shot commit messages ",
	Run: func(cmd *cobra.Command, args []string) {
		var oid string
		if len(args) != 0 {
			oid = args[0]
		}
		if oid == "" {
			oid = "@"
		}

		err := showCommits(base.GetOID(oid))
		if err != nil {
			log.Fatal(err)
		}

	},
}
