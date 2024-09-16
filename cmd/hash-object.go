package cmd

import (
	"fmt"
	"io"
	"lgit/base"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(hashObjectCmd)
}

var hashObjectCmd = &cobra.Command{
	Use:   "hash-object",
	Short: "hash file into object",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("file args is required")
		}
		path := args[0]

		file, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		data, err := io.ReadAll(file)
		if err != nil {
			log.Fatal(err)
		}

		oid, err := base.HashObject(data, "blob")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("object OID: ", oid)
	},
}
