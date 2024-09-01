package precedence

import (
	"goodhumored/lr2_syntax_analyzer/token"
)

type PrecedenceMatrix map[token.TokenType]map[token.TokenType]PrecedenceType

func (matrix PrecedenceMatrix) GetPrecedence(left, right token.TokenType) PrecedenceType {
	if left == token.StartType {
		return Lt
	}
	if right == token.EOFType {
		return Gt
	}
	if val, ok := matrix[left]; ok {
		if precedence, ok := val[right]; ok {
			return precedence
		}
	}
	return Undefined
}
