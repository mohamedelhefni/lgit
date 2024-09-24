package cmd

import (
	"fmt"
	"lgit/base"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(branchCmd)
	branchCmd.PersistentFlags().StringP("name", "n", "", "name of the branch")
	branchCmd.PersistentFlags().StringP("oid", "o", "", "oid of the branch")
}

func createBranch(name, oid string) error {
	return base.SetRef("refs/heads/"+name, base.RefValue{Value: oid, Symbolic: false}, true)
}

var branchCmd = &cobra.Command{
	Use:   "branch",
	Short: "checkout new branch",
	Run: func(cmd *cobra.Command, args []string) {

		var branchName string
		var branchOid string

		if len(args) == 0 {
			current, _ := base.GetBranchName()
			branches, err := base.IterBranches()
			if err != nil {
				log.Fatal(err)
			}
			var prefix string
			for _, branch := range branches {
				prefix = "-"
				if strings.Contains(current, branch) {
					prefix = "*"
				}

				fmt.Println(prefix + " " + branch)
			}
			return
		}

		if name, err := cmd.Flags().GetString("name"); err != nil || name == "" {
			if len(args) >= 1 {
				branchName = args[0]
			}
		} else {
			branchName = name
		}

		if oid, err := cmd.Flags().GetString("oid"); err != nil || oid == "" {
			if len(args) >= 2 {
				branchOid = args[1]
			} else {
				branchOid = "@"
			}
		} else {
			branchOid = oid
		}

		if err := createBranch(branchName, base.GetOID(branchOid)); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Created barnch: ", branchName, "with oid: ", base.GetOID(branchOid))

	},
}
