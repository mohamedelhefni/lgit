package cmd

import (
	"fmt"
	"lgit/base"
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(writeTreeCmd)
}

var writeTreeCmd = &cobra.Command{
	Use:   "write-tree",
	Short: "This command will take the current working directory and store it to the object database.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("file args is required")
		}
		oid, err := base.WriteTree(args[0])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("oid is: ", oid)

	},
}
