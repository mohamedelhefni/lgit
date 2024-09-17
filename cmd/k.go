package cmd

import (
	"fmt"
	"lgit/base"
	"log"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(kCmd)
}

func k() error {
	dot := strings.Builder{}
	dot.WriteString("digraph commits {\n")

	oids := make(map[string]bool)

	refs, err := base.IterRefs()
	if err != nil {
		return err
	}

	for ref := range refs {
		dot.WriteString(fmt.Sprintf("\"%s\" [shape=note]\n", ref.Refname))
		dot.WriteString(fmt.Sprintf("\"%s\" -> \"%s\"\n", ref.Refname, ref.Ref))
		oids[ref.Ref] = true
	}

	oidsSlice := make([]string, 0, len(oids))
	for oid := range oids {
		oidsSlice = append(oidsSlice, oid)
	}

	commitsOids, err := base.IterCommits(oidsSlice)
	if err != nil {
		return err
	}

	for oid := range commitsOids {
		commit, err := base.GetCommit(oid)
		if err != nil {
			return err
		}
		dot.WriteString(fmt.Sprintf("\"%s\" [shape=box style=filled label=\"%s\"]\n", oid, commit))
		if commit.Parent != "" {
			dot.WriteString(fmt.Sprintf("\"%s\" -> \"%s\"\n", oid, commit.Parent))
		}
	}

	dot.WriteString("}")
	dotString := dot.String()
	fmt.Println(dotString)

	cmd := exec.Command("dot", "-Tpng", "-o", "commits.png")
	cmd.Stdin = strings.NewReader(dotString)
	return cmd.Run()
}

var kCmd = &cobra.Command{
	Use:   "k",
	Short: "visualize the mess! ",
	Run: func(cmd *cobra.Command, args []string) {
		if err := k(); err != nil {
			log.Fatal(err)
		}
	},
}
