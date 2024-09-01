package main

import (
	"fmt"
	"os"
	"strings"

	"goodhumored/lr1_analyzer/token"
)

func main() {
	tt := &TokenTable{}               // создаём таблицу лексем
	source := getInput("./input.txt") // читаем файл

	// выводим содержимое
	println("Содержимое входного файла:\n")
	fmt.Println(source)

	// запускаем распознание лексем
	recogniseTokens(source, tt)

	// выводим лексемы
	fmt.Println("Таблица лексем:")
	fmt.Print(tt)
}

// Читает файл с входными данными, вызывает панику в случае неудачи
func getInput(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(data)
}

// Распознаёт токены в данной строке построчно и записывает в таблицу
func recogniseTokens(source string, tokenTable *TokenTable) {
	for _, line := range strings.Split(source, "\n") {
		recogniseTokensLine(line, tokenTable)
	}
}

// Распознаёт лексемы в данной строке и записывает в таблицу
func recogniseTokensLine(line string, tokenTable *TokenTable) {
	for {
		line = strings.Trim(line, " ") // обрезаем пробельные символы в строке
		if len(line) == 0 {            // если строка пустая - завершаем обработку строки
			return
		}
		nextToken := getNextToken(line)      // ищем очередную лексему
		tokenTable.Add(nextToken)            // добавляем лексему в таблицу
		line = line[nextToken.Position.End:] // вырезаем обработанную часть
	}
}

// Ищет очередную лексему в строке
func getNextToken(str string) token.Token {
	// проходим по всем шаблонам лексем
	for _, tokenPattern := range tokenPatterns {
		res := tokenPattern.Pattern.FindStringIndex(str)
		if res != nil {
			return tokenPattern.Type(str[res[0]:res[1]], token.Position{res[0], res[1]})
		}
	}
	return token.Error(str[0:1], token.Position{0, 1})
}
