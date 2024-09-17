package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(tagCmd)
	tagCmd.PersistentFlags().StringP("name", "n", "", "name of the tag")
	tagCmd.PersistentFlags().StringP("oid", "o", "", "the OID of the commit")
}

func createTag(message string) error {

	return nil
}

var tagCmd = &cobra.Command{
	Use:   "tag",
	Short: "tag current commit into readable message",
	Run: func(cmd *cobra.Command, args []string) {

		name, err := cmd.Flags().GetString("name")
		if err != nil {
			log.Fatal(err)
		}

		oid, err := cmd.Flags().GetString("oid")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("name is", name, "oid is", oid)

	},
}
