package rule

import (
	"goodhumored/lr2_syntax_analyzer/syntax_analyzer/nonterminal"
)

type RuleTable struct {
	Rules []Rule
}

func (ruleTable RuleTable) GetRuleByLeftSide(nonterminal nonterminal.NonTerminal) *Rule {
	for _, rule := range ruleTable.Rules {
		if rule.Left == nonterminal {
			return &rule
		}
	}
	return nil
}

func (ruleTable RuleTable) GetRuleByRightSide(tokenTypes []Symbol) *Rule {
	for _, rule := range ruleTable.Rules {
		if isApplyable(rule.Right, tokenTypes) {
			return &rule
		}
	}
	return nil
}

func isApplyable(ruleSymbols, targetSymbols []Symbol) bool {
	ruleLen := len(ruleSymbols)
	targetLen := len(targetSymbols)
	lenDiff := targetLen - ruleLen
	if lenDiff < 0 {
		return false
	}
	for i, ruleSymbol := range ruleSymbols {
		if ruleSymbol.GetName() != targetSymbols[i+lenDiff].GetName() {
			return false
		}
	}
	return true
}
