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
	var oid string
	oid = argOid
	if argOid == "" {
		head, err := base.GetRef("HEAD")
		if err != nil {
			return err
		}
		oid = head
	}
	for oid != "" {
		commit, err := base.GetCommit(oid)
		if err != nil {
			fmt.Println("err: ", err)
			continue
		}
		fmt.Println("commit: ", oid)
		fmt.Println(strings.TrimSpace(commit.Message))
		oid = commit.Parent
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
		err := showCommits(oid)
		if err != nil {
			log.Fatal(err)
		}

	},
}
