package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(readTreeCmd)
}

func emptyCurrentDir(path string) error {
	files, err := os.ReadDir(".")
	if err != nil {
		return err
	}

	for _, file := range files {
		filePath := path + "/" + file.Name()
		if isEscaped(filePath) {
			continue
		}
		if !file.IsDir() {
			os.Remove(filePath)
		} else {
			os.RemoveAll(filePath)
		}
	}
	return nil
}

func treeIter(oid string) ([]Entry, error) {
	var entries []Entry
	tree, err := getObject(oid, "tree")
	if err != nil {
		return entries, err
	}

	for _, object := range strings.Split(tree, "\n") {
		objectArr := strings.Split(object, " ")
		type_, oid, name := objectArr[0], objectArr[1], objectArr[2]
		entries = append(entries, Entry{
			Name: name,
			Type: type_,
			OID:  oid,
		})
	}
	return entries, nil
}

func getTree(oid, basePath string) (map[string]string, error) {
	result := make(map[string]string)
	entries, err := treeIter(oid)
	if err != nil {
		return nil, err
	}
	for _, entry := range entries {
		if strings.Contains(entry.Name, "/") {
			return nil, fmt.Errorf("invalid name %s in tree", entry.Name)
		}
		if entry.Name == ".." || entry.Name == "." {
			return nil, fmt.Errorf("invalid name %s in tree", entry.Name)
		}
		if entry.Name == "lgit" {
			continue
		}

		path := filepath.Join(basePath, entry.Name)
		if entry.Type == "blob" {
			result[path] = entry.OID
		} else if entry.Type == "tree" {
			subTree, err := getTree(entry.OID, path)
			if err != nil {
				return nil, err
			}
			for subPath, subOID := range subTree {
				result[subPath] = subOID
			}
		} else {
			return nil, fmt.Errorf("unknown tree entry %s", entry.Type)
		}
	}
	return result, nil
}

func readTree(oid string) error {
	emptyCurrentDir(".")
	tree, err := getTree(oid, "./")
	if err != nil {
		return err
	}

	for path, oid := range tree {
		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			return err
		}
		data, err := getObject(oid, "blob")
		if err != nil {
			return err
		}
		if err := os.WriteFile(path, []byte(data), 0644); err != nil {
			return err
		}
	}

	return nil
}

var readTreeCmd = &cobra.Command{
	Use:   "read-tree",
	Short: "print tree content byt it's OID ",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("file args is required")
		}
		err := readTree(args[0])
		if err != nil {
			log.Fatal(err)
		}

	},
}
