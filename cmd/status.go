package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"lgit/base"
	"log"
)

func init() {
	rootCmd.AddCommand(statusCmd)
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "get current status",
	Run: func(cmd *cobra.Command, args []string) {
		HEAD := base.GetOID("@")
		branch, err := base.GetBranchName()
		if err != nil {
			log.Fatal(err)
		}

		if branch != "" {
			fmt.Println("on branch ", branch)
		} else {
			fmt.Println("HEAD attached to ", HEAD[:10])
		}

	},
}
