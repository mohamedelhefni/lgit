package base

type Commit struct {
	Tree    string
	Parent  string
	Message string
}

func (c Commit) String() string {
	return c.Message + "(" + c.Tree[:10] + ")"
}

type Entry struct {
	Name string
	Type string
	OID  string
}

type RefValue struct {
	Value    string
	Symbolic bool
}
