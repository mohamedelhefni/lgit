package base

import (
	"strings"
)

func GetCommit(oid string) (Commit, error) {
	var commit Commit
	data, err := GetObject(oid, "commit")
	if err != nil {
		return Commit{}, err
	}

	lines := strings.Split(data, "\n")
	for _, line := range lines {
		parts := strings.Split(line, " ")
		if len(parts) < 2 {
			continue
		}
		key, value := parts[0], parts[1]
		if key == "tree" {
			commit.Tree = value
		} else if key == "parent" {
			commit.Parent = value
		} else {
			continue
			// fmt.Println("Unkown field ", key)
		}
	}
	commit.Message = strings.Join(lines[3:], "\n")
	return commit, nil
}
