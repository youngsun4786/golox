package token

import "fmt"

type TokenType string

type Token struct {
	Type TokenType
	Lexeme string
	Literal string
	Position Position
}

type Position struct {
	Line uint32
	Col uint32
}

func (p *Position) String() string {
	return fmt.Sprintf("%d:%d", p.Line, p.Col)
}

func New(typ TokenType, lexeme, literal string, line, col uint32) Token {
	return Token {
		Type: typ,
		Lexeme: lexeme,
		Literal: literal,
		Position: Position{Line: line, Col: col},
	}
}

const (
	LPAREN TokenType = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
	ERROR = "ERROR"
	EOF = "EOF"
)


var keywords = map[string]TokenType{}



func LookUpIdent(ident string) (TokenType, error) {
	if tok, ok := keywords[ident]; ok {
		return tok, nil
	}
	return ERROR, fmt.Errorf(("LookupIdent: Could not identify ident. Unexpected\n"))
}

func (t*Token) TokenToStr() string {
	switch t.Type {
		case LPAREN: 
			return "LEFT_PAREN"
		case RPAREN: 
			return "RIGHT_PAREN"
		case LBRACE: 
			return "LEFT_BRACE"
		case RBRACE: 
			return "RIGHT_BRACE"
		case EOF:
			return "EOF"
		default:
			return ""
	}
}


func (t* Token) String() string {
	var literal string
	if t.Literal == "" {
		literal = "null"
	}

	return fmt.Sprintf("%s %s %s", t.TokenToStr(), t.Lexeme, literal)
//	return fmt.Sprintf("%s %s %s at %d:%d", t.Type, t.Lexeme, t.Literal, t.Position.Line, t.Position.Col)
}
