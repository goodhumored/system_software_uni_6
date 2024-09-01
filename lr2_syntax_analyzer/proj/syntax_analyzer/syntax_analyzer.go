package syntax_analyzer

import (
	"fmt"

	"goodhumored/lr2_syntax_analyzer/syntax_analyzer/nonterminal"
	"goodhumored/lr2_syntax_analyzer/syntax_analyzer/parse_tree"
	"goodhumored/lr2_syntax_analyzer/syntax_analyzer/precedence"
	"goodhumored/lr2_syntax_analyzer/syntax_analyzer/rule"
	"goodhumored/lr2_syntax_analyzer/token"
	"goodhumored/lr2_syntax_analyzer/token_table"
)

func AnalyzeSyntax(ruleTable rule.RuleTable, tokenTable token_table.TokenTable, matrix precedence.PrecedenceMatrix) (parse_tree.ParseTree, error) {
	rootNode := parse_tree.CreateNode(nonterminal.Null)
	tree := parse_tree.ParseTree{Root: &rootNode}
	tokens := tokenTable.GetTokens()
	tokenIndex := 1
	stack := symbolStack{tokens[0]}

	for {
		stackTerminal := stack.PeekNextTerminal()
		inputToken := tokens[tokenIndex]
		if isInputAccepted(inputToken, stack) {
			return tree, nil
		}
		if inputToken.Type == token.CommentType {
			tokenIndex += 1
			continue
		}

		fmt.Printf("Лексема: '%s' \n", tokens[tokenIndex].Value)
		fmt.Printf("Стек: %s \n", stack)

		prec := matrix.GetPrecedence(stackTerminal.Type, inputToken.Type)

		if prec == precedence.Lt || prec == precedence.Eq {
			tree.AddNode(&parse_tree.Node{Symbol: inputToken, Children: []*parse_tree.Node{}})
			stack = stack.Push(inputToken)
			tokenIndex += 1
			print("Сдвигаем\n")
		} else if prec == precedence.Gt {
			print("Сворачиваем стек\n")
			newStack, rule, err := reduce(stack, ruleTable)
			if err != nil {
				return tree, err
			}
			tree.Reduce(*rule)
			stack = newStack
		} else {
			return tree, fmt.Errorf("Ошибка в синтексе, неожиданное сочетание символов %s и %s (%d)", stackTerminal.GetName(), inputToken.GetName(), inputToken.Position.End)
		}
		println("==============")
	}
}

func isInputAccepted(currentToken token.Token, stack symbolStack) bool {
	nextTerminal := stack.PeekNextTerminal()
	nextSymbol := stack.Peek()
	return currentToken.Type == token.EOFType &&
		nextTerminal != nil &&
		nextTerminal.Type == token.Start.Type &&
		nextSymbol != nil &&
		nextSymbol == nonterminal.E
}

func reduce(stack symbolStack, ruleTable rule.RuleTable) (symbolStack, *rule.Rule, error) {
	for {
		if rule := ruleTable.GetRuleByRightSide(stack); rule != nil {
			fmt.Printf("Нашлось правило: %v, пушим %s в стек\n", rule, rule.Left)
			stack = append(stack[:len(stack)-len(rule.Right)], rule.Left)
			return stack, rule, nil
		} else {
			return stack, nil, fmt.Errorf("Не найдено правил для свёртки")
		}
	}
}
