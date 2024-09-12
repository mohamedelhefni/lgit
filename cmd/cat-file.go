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

var catFileCmd = &cobra.Command{
	Use:   "cat-file",
	Short: "print objecct content byt it's OID ",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("file args is required")
		}
		file, err := os.Open(GIT_DIR + "/objects/" + args[0])
		if err != nil {
			log.Fatal("err: ", err)
		}
		defer file.Close()
		data, err := io.ReadAll(file)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(data))

	},
}
