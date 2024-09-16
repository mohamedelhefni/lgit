package cmd

import (
	"lgit/base"
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(catFileCmd)
}

func catFile(oid, type_ string) error {
	_, err := base.GetObject(oid, type_)
	return err
}

var catFileCmd = &cobra.Command{
	Use:   "cat-file",
	Short: "print objecct content byt it's OID ",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("file args is required")
		}
		err := catFile(args[0], "blob")
		if err != nil {
			log.Fatal(err)
		}

	},
}
