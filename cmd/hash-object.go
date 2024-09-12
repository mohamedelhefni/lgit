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

func hashData(data []byte) string {
	hash := sha1.Sum(data)
	return hex.EncodeToString(hash[:])
}

func hashObject(path, type_ string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	var buff []byte
	buff = append(buff, []byte(type_)...)
	buff = append(buff, '\x00')
	buff = append(buff, data...)
	oid := hashData(buff)
	fmt.Println("object OID: ", oid)
	outFile, err := os.Create(GIT_DIR + "/objects/" + oid)
	if err != nil {
		return err
	}
	defer outFile.Close()
	_, err = outFile.Write(buff)
	if err != nil {
		return err
	}
	return nil
}

var hashObjectCmd = &cobra.Command{
	Use:   "hash-object",
	Short: "hash file into object",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("file args is required")
		}
		path := args[0]
		err := hashObject(path, "blob")
		if err != nil {
			log.Fatal(err)
		}
	},
}
