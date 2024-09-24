package cmd

import (
	"fmt"
	"lgit/base"
	"strings"
)

type CommitArgs struct {
	Oid    string
	Commit base.Commit
	Refs   []string
}

func printCommit(args CommitArgs) {
	var refsString string
	if len(args.Refs) > 0 {
		refsString = fmt.Sprintf("(%s)", strings.Join(args.Refs, ","))
	}
	fmt.Println("commit: ", args.Oid, refsString)
	fmt.Println(strings.TrimSpace(args.Commit.Message))
}
