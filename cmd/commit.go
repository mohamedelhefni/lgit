package cmd

import (
	"fmt"
	"lgit/base"
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(commitCmd)
	commitCmd.PersistentFlags().StringP("message", "m", "", "message for changes in commit")
}

func commit(message string) error {
	tree, err := base.WriteTree("./")
	if err != nil {
		return err
	}
	commitHeaders := fmt.Sprintf("tree %s\n", tree)
	head, err := base.GetHead()
	if err == nil {
		commitHeaders += fmt.Sprintf("parent %s\n", head)
	}
	commitHeaders += "\n"
	commitHeaders += fmt.Sprintf("%s\n", message)
	oid, err := base.HashObject([]byte(commitHeaders), "commit")
	fmt.Println("commit id: ", oid)
	return base.SetHead(oid)
}

var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "commit current changes with message",
	Run: func(cmd *cobra.Command, args []string) {

		message, err := cmd.Flags().GetString("message")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("message is", message)

		err = commit(message)
		if err != nil {
			log.Fatal(err)
		}

	},
}
