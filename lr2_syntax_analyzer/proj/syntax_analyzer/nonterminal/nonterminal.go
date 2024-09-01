package nonterminal

type NonTerminal struct {
	Name string
}

func (nt NonTerminal) GetName() string {
	return nt.Name
}

var (
	E    = NonTerminal{"E"}
	Null = NonTerminal{"/"}
)
