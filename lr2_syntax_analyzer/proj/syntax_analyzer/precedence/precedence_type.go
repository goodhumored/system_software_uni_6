package precedence

type PrecedenceType struct {
	Name string
}

var (
	Lt        = PrecedenceType{"<"}
	Eq        = PrecedenceType{"="}
	Gt        = PrecedenceType{">"}
	Undefined = PrecedenceType{"-"}
)
