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
	head, err := base.GetRef("HEAD", true)
	if err == nil {
		commitHeaders += fmt.Sprintf("parent %s\n", head.Value)
	}
	commitHeaders += "\n\n"
	commitHeaders += fmt.Sprintf("%s\n", message)
	oid, err := base.HashObject([]byte(commitHeaders), "commit")
	fmt.Println("commit id: ", oid)
	return base.SetRef("HEAD", base.RefValue{Value: oid, Symbolic: false}, true)
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
