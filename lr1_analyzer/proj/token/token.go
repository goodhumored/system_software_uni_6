package token

// Типы лексем
const (
	DelimiterType   = "delimiter"   // Разделитель
	IdentifierType  = "identifier"  // Идентификатор
	HexType         = "hex_number"  // Шестнадцатиричное число
	AssignmentType  = "assignment"  // Присваивание
	OperatorType    = "operator"    // Логический оператор
	ParenthesesType = "parentheses" // Скобки
	ErrorType       = "error"       // Ошибка
	CommentType     = "comment"     // Комментарий
)

// Структура Position представляет положение лексемы в строке
type Position struct {
	Start int
	End   int
}

// Структура Token представляет лексему с ее типом и значением
type Token struct {
	Type     string   // Тип
	Value    string   // Значение
	Position Position // Положение лексемы
}

// Фабричная функция для токенов, возвращающая замкнутую лямбда функцию для создания токена определённого типа
func tokenFactory(tokenType string) func(string, Position) Token {
	return func(value string, position Position) Token {
		return Token{
			Value:    value,
			Type:     tokenType,
			Position: position,
		}
	}
}

// Функции создания лексем определённых типов
var (
	Delimiter   = tokenFactory(DelimiterType)
	Identifier  = tokenFactory(IdentifierType)
	Assignment  = tokenFactory(AssignmentType)
	Operator    = tokenFactory(OperatorType)
	Parentheses = tokenFactory(ParenthesesType)
	Hex         = tokenFactory(HexType)
	Error       = tokenFactory(ErrorType)
	Comment     = tokenFactory(CommentType)
)
