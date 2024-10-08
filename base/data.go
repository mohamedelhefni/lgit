package base

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path"
	"regexp"
	"strings"
)

func GetOID(name string) string {
	if name == "@" {
		name = "HEAD"
	}
	refs := []string{
		name,
		fmt.Sprintf("refs/%s", name),
		fmt.Sprintf("refs/tags/%s", name),
		fmt.Sprintf("refs/heads/%s", name),
	}

	for _, ref := range refs {
		if val, err := GetRef(ref, false); err == nil {
			if val.Value != "" {
				value, _ := GetRef(ref, true)
				return value.Value
			}
		}
	}

	// check if valid sha1
	if match, err := regexp.MatchString("^[a-fA-F0-9]{40}$", name); err == nil && match {
		return name
	}
	return ""
}

func SetRef(ref string, value RefValue, deref bool) error {
	ref, _, _ = getRef(ref, deref)
	refPath := GIT_DIR + "/" + ref
	if err := os.MkdirAll(path.Dir(refPath), os.ModePerm); err != nil {
		return err
	}
	var val string
	if value.Symbolic {
		val = "ref: " + value.Value
	} else {
		val = value.Value
	}
	return os.WriteFile(refPath, []byte(val), 0644)
}

func GetRef(ref string, deref bool) (RefValue, error) {
	_, refValue, err := getRef(ref, deref)
	return refValue, err
}

func getRef(ref string, deref bool) (string, RefValue, error) {
	refPath := GIT_DIR + "/" + ref
	var value string
	if stat, err := os.Stat(refPath); err == nil && !stat.IsDir() {
		content, err := os.ReadFile(refPath)
		if err != nil {
			return "", RefValue{}, err
		}
		value = string(content)
	}
	symobilc := value != "" && strings.HasPrefix(value, "ref:")
	if symobilc {
		value = strings.TrimSpace(strings.Split(value, ":")[1])
		if deref {
			return getRef(value, true)
		}
	}
	return ref, RefValue{Value: value, Symbolic: symobilc}, nil
}

func HashData(data []byte) string {
	hash := sha1.Sum(data)
	return hex.EncodeToString(hash[:])
}

func HashObject(data []byte, type_ string) (string, error) {
	var buff []byte
	buff = append(buff, []byte(type_)...)
	buff = append(buff, '\x00')
	buff = append(buff, data...)
	oid := HashData(buff)
	outFile, err := os.Create(GIT_DIR + "/objects/" + oid)
	if err != nil {
		return "", err
	}
	defer outFile.Close()
	_, err = outFile.Write(buff)
	if err != nil {
		return "", err
	}
	return oid, nil
}

func GetObject(oid, type_ string) (string, error) {
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
