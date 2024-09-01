package main

import (
	"fmt"
	"os"

	"goodhumored/lr2_syntax_analyzer/syntax_analyzer"
	"goodhumored/lr2_syntax_analyzer/syntax_analyzer/nonterminal"
	"goodhumored/lr2_syntax_analyzer/syntax_analyzer/precedence"
	"goodhumored/lr2_syntax_analyzer/syntax_analyzer/rule"
	"goodhumored/lr2_syntax_analyzer/token"
	"goodhumored/lr2_syntax_analyzer/token_analyzer"
)

// Правила грамматики
var rulesTable = rule.RuleTable{Rules: []rule.Rule{
	{Left: nonterminal.E, Right: []rule.Symbol{token.IdentifierType, token.AssignmentType, nonterminal.E, token.DelimiterType}},
	{Left: nonterminal.E, Right: []rule.Symbol{nonterminal.E, token.OrType, nonterminal.E}},
	{Left: nonterminal.E, Right: []rule.Symbol{nonterminal.E, token.XorType, nonterminal.E}},
	{Left: nonterminal.E, Right: []rule.Symbol{nonterminal.E, token.AndType, nonterminal.E}},
	{Left: nonterminal.E, Right: []rule.Symbol{token.NotType, token.LeftParenthType, nonterminal.E, token.RightParenthType}},
	{Left: nonterminal.E, Right: []rule.Symbol{token.LeftParenthType, nonterminal.E, token.RightParenthType}},
	{Left: nonterminal.E, Right: []rule.Symbol{token.IdentifierType}},
}}

var precedenceMatrix = precedence.PrecedenceMatrix{
	token.IdentifierType:   map[token.TokenType]precedence.PrecedenceType{token.AssignmentType: precedence.Eq, token.RightParenthType: precedence.Gt, token.OrType: precedence.Gt, token.XorType: precedence.Gt, token.AndType: precedence.Gt, token.DelimiterType: precedence.Gt},
	token.AssignmentType:   map[token.TokenType]precedence.PrecedenceType{token.IdentifierType: precedence.Lt, token.LeftParenthType: precedence.Lt, token.RightParenthType: precedence.Lt, token.NotType: precedence.Lt, token.OrType: precedence.Lt, token.XorType: precedence.Lt, token.AndType: precedence.Lt, token.DelimiterType: precedence.Eq},
	token.LeftParenthType:  map[token.TokenType]precedence.PrecedenceType{token.IdentifierType: precedence.Lt, token.LeftParenthType: precedence.Lt, token.RightParenthType: precedence.Eq, token.NotType: precedence.Lt, token.OrType: precedence.Lt, token.XorType: precedence.Lt, token.AndType: precedence.Lt},
	token.RightParenthType: map[token.TokenType]precedence.PrecedenceType{token.RightParenthType: precedence.Gt, token.OrType: precedence.Gt, token.XorType: precedence.Gt, token.AndType: precedence.Gt, token.DelimiterType: precedence.Gt},
	token.NotType:          map[token.TokenType]precedence.PrecedenceType{token.LeftParenthType: precedence.Lt},
	token.OrType:           map[token.TokenType]precedence.PrecedenceType{token.IdentifierType: precedence.Lt, token.LeftParenthType: precedence.Lt, token.RightParenthType: precedence.Gt, token.NotType: precedence.Lt, token.OrType: precedence.Gt, token.XorType: precedence.Gt, token.AndType: precedence.Lt, token.DelimiterType: precedence.Gt},
	token.XorType:          map[token.TokenType]precedence.PrecedenceType{token.IdentifierType: precedence.Lt, token.LeftParenthType: precedence.Lt, token.RightParenthType: precedence.Gt, token.NotType: precedence.Lt, token.OrType: precedence.Gt, token.XorType: precedence.Gt, token.AndType: precedence.Lt, token.DelimiterType: precedence.Gt},
	token.AndType:          map[token.TokenType]precedence.PrecedenceType{token.IdentifierType: precedence.Lt, token.LeftParenthType: precedence.Lt, token.RightParenthType: precedence.Gt, token.NotType: precedence.Lt, token.OrType: precedence.Gt, token.XorType: precedence.Gt, token.AndType: precedence.Gt, token.DelimiterType: precedence.Gt},
}

func main() {
	source := getInput("./input.txt") // читаем файл

	// выводим содержимое
	println("Содержимое входного файла:\n")
	fmt.Println(source)

	// запускаем распознание лексем
	tokenTable := token_analyzer.RecogniseTokens(source)

	// выводим лексемы
	fmt.Println("Таблица лексем:")
	fmt.Println(tokenTable)

	if errors := tokenTable.GetErrors(); len(errors) > 0 {
		fmt.Printf("Во время лексического анализа было обнаружено: %d ошибок:\n", len(errors))
		for _, error := range errors {
			fmt.Printf("Неожиданный символ '%s'\n", error.Value)
		}
		return
	}

	// запускаем синтаксический анализатор
	tree, error := syntax_analyzer.AnalyzeSyntax(rulesTable, *tokenTable, precedenceMatrix)
	if error != nil {
		fmt.Printf("Ошибка при синтаксическом анализе строки: %s", error)
	} else {
		fmt.Println("Строка принята!!!")
		tree.Print()
	}
}

// Читает файл с входными данными, вызывает панику в случае неудачи
func getInput(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(data)
}
