package parse_tree

import (
	"fmt"

	"goodhumored/lr2_syntax_analyzer/syntax_analyzer/rule"
)

type Node struct {
	Symbol   rule.Symbol
	Children []*Node
}

func CreateNode(s rule.Symbol) Node {
	return Node{s, []*Node{}}
}

func (n *Node) AddChild(child *Node) {
	n.Children = append(n.Children, child)
}

func (n *Node) GetChildWithName(name string) *Node {
	if n.Symbol.GetName() == name {
		return n
	}
	for _, child := range n.Children {
		if childWithName := child.GetChildWithName(name); childWithName != nil {
			return childWithName
		}
	}
	return nil
}

func (node *Node) Reduce(rule rule.Rule) bool {
	if !node.CanApplyRule(rule) {
		return false
	}
	lenDiff := len(node.Children) - len(rule.Right)

	nodes := make([]*Node, len(rule.Right))
	copy(nodes, node.Children[lenDiff:])

	node.Children = append(node.Children[:lenDiff], &Node{rule.Left, nodes})
	return true
}

func (node Node) CanApplyRule(rule rule.Rule) bool {
	lenDiff := len(node.Children) - len(rule.Right)
	if lenDiff < 0 {
		return false
	}
	for i, rule := range rule.Right {
		if rule.GetName() != node.Children[i+lenDiff].Symbol.GetName() {
			return false
		}
	}
	return true
}

func (node *Node) Print(prefix string, isTail bool) {
	// Выводим символ узла с отступом
	var branch, prefixSuffix string
	if isTail {
		prefixSuffix = "    "
		branch = "└── "
	} else {
		branch = "├── "
		prefixSuffix = "│   "
	}
	fmt.Println(prefix + branch + node.Symbol.GetName())

	// Рекурсивно выводим дочерние узлы
	for i := 0; i < len(node.Children)-1; i++ {
		node.Children[i].Print(prefix+prefixSuffix, false)
	}
	if len(node.Children) > 0 {
		node.Children[len(node.Children)-1].Print(prefix+prefixSuffix, true)
	}
}
