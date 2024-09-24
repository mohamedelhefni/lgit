package cmd

import (
	"lgit/base"
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(resetCmd)
}

func reset(oid string) error {
	return base.SetRef("HEAD", base.RefValue{
		Value:    oid,
		Symbolic: false,
	}, true)
}

var resetCmd = &cobra.Command{
	Use: "reset",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("file args is required")
		}
		err := reset(base.GetOID(args[0]))
		if err != nil {
			log.Fatal(err)
		}

	},
}
