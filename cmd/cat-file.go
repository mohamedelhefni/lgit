package cmd

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(catFileCmd)
}

func catFile(path, type_ string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	fmt.Println(string(data[len([]byte(type_))+1:]))

	return nil
}

var catFileCmd = &cobra.Command{
	Use:   "cat-file",
	Short: "print objecct content byt it's OID ",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("file args is required")
		}
		path := GIT_DIR + "/objects/" + args[0]
		err := catFile(path, "blob")
		if err != nil {
			log.Fatal(err)
		}

	},
}
