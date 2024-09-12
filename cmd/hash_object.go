package cmd

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(hashObjectCmd)
}

func hash_data(data []byte) string {
	hash := sha1.Sum(data)
	return hex.EncodeToString(hash[:])
}

var hashObjectCmd = &cobra.Command{
	Use:   "hash-object",
	Short: "hash file into object",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("file args is required")
		}
		file, err := os.Open(args[0])
		if err != nil {
			log.Fatal("err: ", err)
		}
		defer file.Close()
		data, err := io.ReadAll(file)
		if err != nil {
			log.Fatal(err)
		}
		oid := hash_data(data)
		fmt.Println("object OID: ", oid)
		outFile, err := os.Create(GIT_DIR + "/objects/" + oid)
		if err != nil {
			log.Fatal(err)
		}
		defer outFile.Close()
		_, err = outFile.Write(data)
		if err != nil {
			log.Fatal(err)
		}
	},
}
