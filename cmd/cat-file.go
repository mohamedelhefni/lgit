package cmd

import (
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(catFileCmd)
}

func getObject(oid, type_ string) (string, error) {
	file, err := os.Open(GIT_DIR + "/objects/" + oid)
	if err != nil {
		return "", err
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(data[len([]byte(type_))+1:]), nil
}

func catFile(oid, type_ string) error {
	_, err := getObject(oid, type_)
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
