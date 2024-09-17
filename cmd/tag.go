package cmd

import (
	"lgit/base"
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(tagCmd)
	tagCmd.PersistentFlags().StringP("name", "n", "", "name of the tag")
	tagCmd.PersistentFlags().StringP("oid", "o", "", "the OID of the commit")
}

func createTag(name, oid string) error {
	return base.SetRef("refs/tags/"+name, oid)
}

var tagCmd = &cobra.Command{
	Use:   "tag",
	Short: "tag current commit into readable message",
	Run: func(cmd *cobra.Command, args []string) {
		var tagName string
		var tagOid string

		if name, err := cmd.Flags().GetString("name"); err != nil || name == "" {
			if len(args) > 1 {
				tagName = args[0]
			}
		} else {
			tagName = name
		}

		if oid, err := cmd.Flags().GetString("oid"); err != nil || oid == "" {
			if len(args) >= 2 {
				tagOid = args[1]
			}
		} else {
			tagOid = oid
		}

		if err := createTag(tagName, base.GetOID(tagOid)); err != nil {
			log.Fatal(err)
		}

	},
}
