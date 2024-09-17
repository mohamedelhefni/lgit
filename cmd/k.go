package cmd

import (
	"fmt"
	"lgit/base"
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(kCmd)
}

func k() error {
	refs, err := base.IterRefs()
	if err != nil {
		return err
	}
	for ref := range refs {
		fmt.Println(ref)
	}

	return nil
}

var kCmd = &cobra.Command{
	Use:   "k",
	Short: "visualize the mess! ",
	Run: func(cmd *cobra.Command, args []string) {
		if err := k(); err != nil {
			log.Fatal(err)
		}
	},
}
