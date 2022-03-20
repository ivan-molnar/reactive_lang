package token

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF" // End Of File
	NL      = "NL"  // New Line

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1343456

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	EQ     = "=="
	NOT_EQ = "!="
	LT_EQ  = "<="
	GT_EQ  = ">="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	CONSTANT = "CONSTANT"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	ELIF     = "ELIF"
	SWITCH   = "SWITCH"
	FOR      = "FOR"
	WHILE    = "WHILE"
	BREAK    = "BREAK"
	CONTINUE = "CONTINUE"
	RETURN   = "RETURN"
)

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"func":     FUNCTION,
	"let":      LET,
	"const":    CONSTANT,
	"true":     TRUE,
	"false":    FALSE,
	"if":       IF,
	"else":     ELSE,
	"elif":     ELIF,
	"switch":   SWITCH,
	"for":      FOR,
	"while":    WHILE,
	"break":    BREAK,
	"continue": CONTINUE,
	"return":   RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

/* TO BE ADDED LATER: */

// += -= *= /=
// static (<- a classical, non-reactive variable/object)
// class object
// & ^ |
// && ^^ ||
// [] (<- when arrays will be a thing)
// str char array float bool
// Math
