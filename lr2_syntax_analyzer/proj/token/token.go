package token

import "fmt"

// Структура Position представляет положение лексемы в строке
type Position struct {
	Start int
	End   int
}

// Структура Token представляет лексему с ее типом и значением
type Token struct {
	Type     TokenType // Тип
	Value    string    // Значение
	Position Position  // Положение лексемы
}

// Функция получения имени токена, для соответствия интерфейсу символа
func (token Token) GetName() string {
	return token.Type.GetName()
}

// Фабричная функция для токенов, возвращающая замкнутую лямбда функцию для создания токена определённого типа
func tokenFactory(tokenType TokenType) func(string, Position) Token {
	return func(value string, position Position) Token {
		return Token{
			Value:    value,
			Type:     tokenType,
			Position: position,
		}
	}
}

// Функция определяющая как токен преобразуется в строку
func (token Token) String() string {
	return fmt.Sprintf("%s (%s)", token.Type, token.Value)
}

// Функции создания лексем определённых типов
var (
	Delimiter    = tokenFactory(DelimiterType)
	Identifier   = tokenFactory(IdentifierType)
	Assignment   = tokenFactory(AssignmentType)
	And          = tokenFactory(AndType)
	Or           = tokenFactory(OrType)
	Xor          = tokenFactory(XorType)
	Not          = tokenFactory(NotType)
	LeftParenth  = tokenFactory(LeftParenthType)
	RightParenth = tokenFactory(RightParenthType)
	Hex          = tokenFactory(HexType)
	Error        = tokenFactory(ErrorType)
	Comment      = tokenFactory(CommentType)
	Start        = Token{StartType, "", Position{0, 0}}
	EOF          = Token{EOFType, "EOF", Position{0, 0}}
)
