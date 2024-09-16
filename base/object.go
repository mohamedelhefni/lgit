package base

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"os"
)

func SetHead(oid string) error {
	return os.WriteFile(GIT_DIR+"/HEAD", []byte(oid), 0644)
}

func GetHead() (string, error) {
	content, err := os.ReadFile(GIT_DIR + "/HEAD")
	return string(content), err
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
