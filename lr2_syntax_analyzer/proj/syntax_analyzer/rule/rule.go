package rule

import (
	"fmt"

	"goodhumored/lr2_syntax_analyzer/syntax_analyzer/nonterminal"
)

type Symbol interface {
	GetName() string
}

type Rule struct {
	Left  nonterminal.NonTerminal
	Right []Symbol
}

func (r Rule) String() string {
	return fmt.Sprintf("%s -> %s", r.Left.GetName(), r.Right)
}
