package main

import (
	"fmt"
	"goodhumored/lr1_analyzer/token"

	"strings"
)

// Таблица лексем
type TokenTable struct {
	tokens []token.Token
}

// Метод добавления лексемы в таблицу
func (ts *TokenTable) Add(token token.Token) {
	ts.tokens = append(ts.tokens, token)
}

// Вспомогательная функция для генерации строки с таблицей лексем
func (ts *TokenTable) String() string {
	if len(ts.tokens) == 0 {
		return "Ни одного токена не найдено"
	}

	// Определяем максимальную ширину столбца
	maxTypeLen := len("Тип")
	maxValueLen := len("Значение")
	for _, token := range ts.tokens {
		if len(token.Type) > maxTypeLen {
			maxTypeLen = len(token.Type)
		}
		if len(token.Value) > maxValueLen {
			maxValueLen = len(token.Value)
		}
	}

	// создаем шапку и рамки
	header := fmt.Sprintf("| %-*s | %-*s |", maxTypeLen, "Тип", maxValueLen, "Значение")
	border := fmt.Sprintf("+-%s-+-%s-+", strings.Repeat("-", maxTypeLen), strings.Repeat("-", maxValueLen))

	// Собираем таблицу
	res := border + "\n" + header + "\n" + border + "\n"
	for _, token := range ts.tokens {
		res += fmt.Sprintf("| %-*s | %-*s |\n", maxTypeLen, token.Type, maxValueLen, token.Value)
	}
	res += border

	return res
}

// Вспомогательная функция для печати таблицы
func (ts *TokenTable) Print() {
	fmt.Println(ts.String())
}
