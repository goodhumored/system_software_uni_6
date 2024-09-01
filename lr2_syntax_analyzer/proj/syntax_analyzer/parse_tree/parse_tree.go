package parse_tree

import (
	"fmt"

	"goodhumored/lr2_syntax_analyzer/syntax_analyzer/rule"
)

type ParseTree struct {
	Root *Node
}

func (tree ParseTree) Print() {
	tree.Root.Print("", true)
}

func (tree ParseTree) FindNode(symbolName string) *Node {
	return tree.Root.GetChildWithName(symbolName)
}

func (tree *ParseTree) AddNode(node *Node) {
	tree.Root.AddChild(node)
}

func (tree *ParseTree) Reduce(rule rule.Rule) {
	fmt.Printf("Применяем правило %s к дереву\n", rule)
	if tree.Root.Reduce(rule) {
		fmt.Printf("Успешно применено\n")
	} else {
		fmt.Printf("Правило %s применить не удалось\n", rule)
	}
}
