package base

type Commit struct {
	Tree    string
	Parent  string
	Message string
}

type Entry struct {
	Name string
	Type string
	OID  string
}
