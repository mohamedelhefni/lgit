package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(writeTreeCmd)
}

func isEscaped(path string) bool {
	if strings.Contains(path, GIT_DIR) {
		return true
	} else if strings.Contains(path, ".git") {
		return true
	}
	return false
}

func writeTree(path string) error {
	files, err := os.ReadDir(path)
	if err != nil {
		return err
	}
	for _, file := range files {
		filePath := path + "/" + file.Name()
		if isEscaped(filePath) {
			continue
		}
		if file.IsDir() {
			writeTree(filePath)
		} else {
			file, _ := os.ReadFile(filePath)
			fmt.Println(filePath, " - ", hashData(file))
		}
	}
	return nil
}

var writeTreeCmd = &cobra.Command{
	Use:   "write-tree",
	Short: "This command will take the current working directory and store it to the object database.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("file args is required")
		}
		err := writeTree(args[0])
		if err != nil {
			log.Fatal(err)
		}

	},
}
