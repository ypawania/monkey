package token

type TokenType string

type Token struct {
	Type TokenType
	literal string
}

const {
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	//IDENTIFIERS + literal
	IDENT = "IDENT"
	INT = "INT"

	//OPERATORS
	ASSIGN = "="
	PLUS = "+"
	MINUS = "-"
	BANG = "!"
	ASTERISK = "*"
	SLASH = "/"

	LT = "<"
	GT = ">"

	EQ = "=="
	NOT_EQ = "!="

	//DELIMETERS
	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
	
	//KEYWORDS
	FUNCTION = "FUNCTION"
	LET = "LET"
	TRUE = "TRUE"
	FALSE = "FALSE"
	IF = "IF"
	ELSE = "ELSE"
	RETURN = "RETURN"
}

