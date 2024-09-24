package cmd

import (
	"fmt"
	"lgit/base"
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(logCmd)
}

func showCommits(argOid string) error {
	refs := map[string][]string{}
	refCh, err := base.IterRefs("", false)
	if err != nil {
		return err
	}
	for refResult := range refCh {
		if _, ok := refs[refResult.Ref]; !ok {
			refs[refResult.Ref] = []string{}
		}
		refs[refResult.Ref] = append(refs[refResult.Ref], refResult.Refname)
	}

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
		printCommit(CommitArgs{
			Oid:    oid,
			Commit: commit,
			Refs:   refs[oid],
		})

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
