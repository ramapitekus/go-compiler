package scanner

type Token struct {
	Type   TokenType
	Length int
	Line   int
	Start  int32
}

type TokenType uint8

const (
	TOKEN_LEFT_PAREN TokenType = iota
	TOKEN_RIGHT_PAREN
	TOKEN_LEFT_BRACE
	TOKEN_RIGHT_BRACE
	TOKEN_COMMA
	TOKEN_DOT
	TOKEN_MINUS
	TOKEN_PLUS
	TOKEN_SEMICOLON
	TOKEN_SLASH
	TOKEN_STAR
	TOKEN_BANG
	TOKEN_BANG_EQUAL
	TOKEN_EQUAL
	TOKEN_EQUAL_EQUAL
	TOKEN_GREATER
	TOKEN_GREATER_EQUAL
	TOKEN_LESS
	TOKEN_LESS_EQUAL
	TOKEN_IDENTIFIER
	TOKEN_STRING
	TOKEN_NUMBER
	TOKEN_AND
	TOKEN_CLASS
	TOKEN_ELSE
	TOKEN_FALSE
	TOKEN_FOR
	TOKEN_FUN
	TOKEN_IF
	TOKEN_NIL
	TOKEN_OR
	TOKEN_PRINT
	TOKEN_RETURN
	TOKEN_SUPER
	TOKEN_THIS
	TOKEN_TRUE
	TOKEN_VAR
	TOKEN_WHILE
	TOKEN_ERROR
	TOKEN_EOF
)

type Scanner struct {
	source  string
	start   int32
	current int32
	line    int
}

func InitScanner(source string) *Scanner {
	return &Scanner{start: 0, current: 0, line: 1, source: source}
}

func (scanner *Scanner) ScanToken() Token {
	scanner.start = scanner.current
	// TODO: implement parser - read next character and check if it matches.
	return makeToken(TOKEN_AND)
}

func (scanner *Scanner) makeToken(tokenType TokenType) Token {
	return Token{Type: tokenType, Start: scanner.start, Length: int(scanner.current - scanner.start), Line: scanner.line}
}

func (scanner *Scanner) errorToken(message string) Token {
	return Token{Type: TOKEN_ERROR, Line: scanner.line} // TODO: Add start and length
}