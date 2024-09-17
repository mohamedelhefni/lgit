package cmd

import (
	"lgit/base"
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(readTreeCmd)
}

var readTreeCmd = &cobra.Command{
	Use:   "read-tree",
	Short: "print tree content byt it's OID ",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("file args is required")
		}
		err := base.ReadTree(base.GetOID(args[0]))
		if err != nil {
			log.Fatal(err)
		}

	},
}
