package base

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"regexp"
)

func GetOID(name string) string {
	refs := []string{
		name,
		fmt.Sprintf("refs/%s", name),
		fmt.Sprintf("refs/tags/%s", name),
		fmt.Sprintf("refs/heads/%s", name),
	}

	for _, ref := range refs {
		if oid, err := GetRef(ref); err == nil {
			return oid
		}
	}

	// check if valid sha1
	if match, err := regexp.MatchString("^[a-fA-F0-9]{40}$", name); err == nil && match {
		return name
	}
	return ""
}

func SetRef(ref, oid string) error {
	refPath := GIT_DIR + "/" + ref
	if err := os.MkdirAll(path.Dir(refPath), os.ModePerm); err != nil {
		return err
	}
	return os.WriteFile(refPath, []byte(oid), 0644)
}

func GetRef(ref string) (string, error) {
	refPath := GIT_DIR + "/" + ref
	if stat, err := os.Stat(refPath); err == nil && !stat.IsDir() {
		content, err := os.ReadFile(refPath)
		return string(content), err
	}
	return "", errors.New("file not exist")
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
