package token

import "fmt"

type TokenType string

type Token struct {
	Type TokenType
	Lexeme string
	Literal interface{}
	Position Position
}

type Position struct {
	Line uint32
	Col uint32
}

func (p *Position) String() string {
	return fmt.Sprintf("%d:%d", p.Line, p.Col)
}

func New(typ TokenType, lexeme string, literal interface{}, line, col uint32) Token {
	return Token {
		Type: typ,
		Lexeme: lexeme,
		Literal: literal,
		Position: Position{Line: line, Col: col},
	}
}

const (
	STR_LITERAL TokenType = "STR_LITERAL"
	NUM_LITERAL = "NUM_LITERAL"
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
	COMMA = ","
	DOT = "."
	MINUS = "-"
	PLUS = "+"
	SEMICOLON = ";"
	DIV = "/"
	STAR = "*"
	// OPERATORS
	ASSIGN = "="
	NOT = "!"
	EQ = "=="
	NE = "!="
	GT = ">"
	GE = ">="
	LT = "<"
	LE = "<="
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
		case STR_LITERAL:
			return "STRING"
		case NUM_LITERAL:
			return "NUMBER"
		case LPAREN: 
			return "LEFT_PAREN"
		case RPAREN: 
			return "RIGHT_PAREN"
		case LBRACE: 
			return "LEFT_BRACE"
		case RBRACE: 
			return "RIGHT_BRACE"
		case COMMA: 
			return "COMMA"
		case DOT: 
			return "DOT"
		case MINUS:	
			return "MINUS"
		case PLUS: 
			return "PLUS"
		case SEMICOLON: 
			return "SEMICOLON"
		case DIV: 
			return "SLASH"
		case STAR: 
			return "STAR"
		case ASSIGN:
			return "EQUAL"
		case NOT:
			return "BANG"	
		case EQ:
			return "EQUAL_EQUAL"	
		case NE:
			return "BANG_EQUAL"	
		case GT:
			return "GREATER"	
		case GE:
			return "GREATER_EQUAL"	
		case LT:
			return "LESS"	
		case LE:
			return "LESS_EQUAL"	
		case EOF:
			return "EOF"	
		default:
			return ""
	}
}

func formatFloat(f float64) string {
    // Check if it's a whole number
	if f == float64(int64(f)) {
		return fmt.Sprintf("%.1f", f) 
	}
	// Normal formatting for non-whole numbers
	return fmt.Sprintf("%g", f) 
}


func (t* Token) String() string {
	if t.Literal == "" {
		t.Literal = "null"
	}

	    // Handle float64 specially to preserve precision
	if f, ok := t.Literal.(float64); ok {
		return fmt.Sprintf("%s %s %s", t.Type, t.Lexeme, formatFloat(f))		
	}

	return fmt.Sprintf("%s %s %v", t.TokenToStr(), t.Lexeme, t.Literal)
//	return fmt.Sprintf("%s %s %s at %d:%d", t.Type, t.Lexeme, t.Literal, t.Position.Line, t.Position.Col)
}
