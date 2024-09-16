package base

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func isEscaped(path string) bool {
	if strings.Contains(path, GIT_DIR) {
		return true
	} else if strings.Contains(path, ".git") {
		return true
	} else if strings.Contains(path, "lgit") {
		return true
	}
	return false
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
	tree, err := func() (string, error) {
		file, err := os.Open(GIT_DIR + "/objects/" + oid)
		if err != nil {
			return "", err
		}
		defer file.Close()
		data, err := io.ReadAll(file)
		if err != nil {
			return "", err
		}
		return string(data[len([]byte("tree"))+1:]), nil
	}()
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

func ReadTree(oid string) error {
	emptyCurrentDir(".")
	tree, err := getTree(oid, "./")
	if err != nil {
		return err
	}

	for path, oid := range tree {
		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			return err
		}
		data, err := GetObject(oid, "blob")
		if err != nil {
			return err
		}
		if err := os.WriteFile(path, []byte(data), 0644); err != nil {
			return err
		}
	}

	return nil
}

func WriteTree(path string) (string, error) {
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
			oid, _ = WriteTree(filePath)
		} else {
			file, _ := os.ReadFile(filePath)
			type_ = "blob"
			oid, _ = HashObject(file, "blob")
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
	return HashObject([]byte(tree), "tree")
}
