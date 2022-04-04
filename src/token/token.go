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

	// Double perators
	PLUS_AS = "+="
	MIN_AS  = "-="
	MULT_AS = "*="
	DIV_AS  = "/="

	// Comparators
	LT = "<"
	GT = ">"

	// Double comparators
	EQ     = "=="
	NOT_EQ = "!="
	LT_EQ  = "<="
	GT_EQ  = ">="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	DOT       = "."

	LPAREN  = "("
	RPAREN  = ")"
	LBRACE  = "{"
	RBRACE  = "}"
	LSQUARE = "["
	RSQUARE = "]"

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
	PRINT    = "PRINT"
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
	"print":    PRINT,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

/* TO BE ADDED LATER: */

// static 						(<- a classical, non-reactive variable/object, might need a better keyword)
// class						(<- this one will be tricky to implement)
// & ^ |						(<- binary operators)
// && ^^ ||						(<- logical operators)
// str char array float object 	(<- to define / change types)
// "" '' ``						(<- books says how to do these towards the very end)
