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

type Entry struct {
	Name string
	Type string
	OID  string
}

func writeTree(path string) (string, error) {
	var entries []Entry
	files, err := os.ReadDir(path)
	if err != nil {
		return "", err
	}

	for _, file := range files {
		filePath := path + "/" + file.Name()
		if isEscaped(filePath) {
			continue
		}
		var type_ string
		var oid string

		if file.IsDir() {
			type_ = "tree"
			oid, _ = writeTree(filePath)
		} else {
			file, _ := os.ReadFile(filePath)
			type_ = "blob"
			oid, _ = hashObject(file, "blob")
		}
		entries = append(entries, Entry{
			Name: file.Name(),
			Type: type_,
			OID:  oid,
		})

	}
	var treeNodes []string
	for _, entry := range entries {
		treeNodes = append(treeNodes, fmt.Sprintf("%s %s %s", entry.Type, entry.OID, entry.Name))
	}
	tree := strings.Join(treeNodes, "\n")
	return hashObject([]byte(tree), "tree")
}

var writeTreeCmd = &cobra.Command{
	Use:   "write-tree",
	Short: "This command will take the current working directory and store it to the object database.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("file args is required")
		}
		oid, err := writeTree(args[0])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("oid is: ", oid)

	},
}
