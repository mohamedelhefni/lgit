package base

import (
	"container/list"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func IterCommits(initOids []string) (chan string, error) {
	ch := make(chan string)

	go func() {
		defer close(ch)
		oids := list.New()
		for _, oid := range initOids {
			oids.PushBack(oid)
		}
		visited := map[string]bool{}
		for oids.Len() > 0 {
			ele := oids.Front()
			oids.Remove(ele)
			oid := ele.Value.(string)
			if oid == "" || visited[oid] {
				continue
			}
			ch <- oid
			visited[oid] = true
			commit, err := GetCommit(oid)
			if err != nil {
				fmt.Println("err: ", err)
				continue
			}
			if commit.Parent != "" {
				oids.PushBack(commit.Parent)
			}

		}

	}()

	return ch, nil
}

type RefResult struct {
	Refname string
	Ref     string
}

func (r RefResult) String() string {
	return r.Refname + " -> " + r.Ref
}

func IterBranches() ([]string, error) {
	var branches []string
	res, err := IterRefs("refs/heads/", true)
	if err != nil {
		return branches, err
	}
	for branch := range res {
		branches = append(branches, GetBranchRelativeName(branch.Refname))
	}
	return branches, nil
}

func IterRefs(prefix string, deref bool) (chan RefResult, error) {
	ch := make(chan RefResult)

	go func() {
		defer close(ch)
		refs := []string{"HEAD"}

		err := filepath.Walk(filepath.Join(GIT_DIR, "refs"),
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if !info.IsDir() {
					relPath, err := filepath.Rel(GIT_DIR, path)
					if err != nil {
						return err
					}
					refs = append(refs, relPath)
				}
				return nil
			})

		if err != nil {
			fmt.Println("Error walking through refs:", err)
			return
		}

		for _, refname := range refs {
			if !strings.HasPrefix(refname, prefix) {
				continue
			}
			ref, err := GetRef(refname, deref)
			if err != nil {
				fmt.Printf("Error getting ref for %s: %v\n", refname, err)
				continue
			}
			ch <- RefResult{
				Refname: refname,
				Ref:     ref.Value,
			}
		}
	}()
	return ch, nil
}

func IterTree(oid string) ([]Entry, error) {
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
		if len(objectArr) < 2 {
			continue
		}
		type_, oid, name := objectArr[0], objectArr[1], objectArr[2]
		entries = append(entries, Entry{
			Name: name,
			Type: type_,
			OID:  oid,
		})
	}
	return entries, nil
}
