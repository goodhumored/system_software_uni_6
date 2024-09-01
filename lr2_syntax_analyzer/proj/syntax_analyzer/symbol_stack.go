package syntax_analyzer

import (
	"fmt"

	"goodhumored/lr2_syntax_analyzer/syntax_analyzer/rule"
	"goodhumored/lr2_syntax_analyzer/token"
)

type symbolStack []rule.Symbol

func (s symbolStack) Push(e rule.Symbol) symbolStack {
	return append(s, e)
}

func (s symbolStack) Pop() (symbolStack, rule.Symbol) {
	length := len(s)
	if length == 0 {
		return s, nil
	}
	return s[:length-1], s[length-1]
}

func (s symbolStack) Peek() rule.Symbol {
	length := len(s)
	if length == 0 {
		return nil
	}
	return s[length-1]
}

func (s symbolStack) PeekN(n int) rule.Symbol {
	length := len(s)
	if length == 0 {
		return nil
	}
	return s[length-n-1]
}

func (s symbolStack) PeekNextTerminal() *token.Token {
	for i := range s {
		symbol := s.PeekN(i)
		if token, ok := symbol.(token.Token); ok {
			return &token
		}
	}
	return nil
}

func (s symbolStack) String() string {
	str := ""
	for _, i := range s {
		str += fmt.Sprintf("%s ", i.GetName())
	}
	return str
}

func (s symbolStack) Print() {
	fmt.Print(s.String())
}
