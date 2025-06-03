package lexer

import (
	"github.com/codecrafters-io/interpreter-starter-go/token"
	"fmt"
	"strconv"
)

type Lexer struct {
	source   string
	filename string
	start    int
	current  int
	line     uint32
	column   uint32
}

func New(source string, filename string) *Lexer {
	return &Lexer {
		source: source,
		filename: filename,
		start: 0,
		current: 0,
		line: 1,
		column: 1,
	}
}

func (l *Lexer) IsAtEnd() bool {
	return l.current >= len(l.source)
}

func (l *Lexer) advance() rune {
	ch := rune(l.source[l.current])
	l.current++
	l.column++
	return ch
}

func (l *Lexer) peek() rune {
	if l.IsAtEnd() {
		return 0 
	}
	return rune(l.source[l.current])
}

func (l *Lexer) peekNext() rune {
	if l.current + 1 >= len(l.source) {
		return '0'
	}
	return rune(l.source[l.current + 1])
}

func (l *Lexer) resetCursor() {
	l.line++
	l.column = 1
}

// TOKENIZER
func (l *Lexer) skipWhitespace() {
	for l.peek() == ' ' || l.peek() == '\r' || l.peek() == '\t' || l.peek() == '\n' {
		if l.peek() == '\n' {
			l.resetCursor()
		}
		l.advance()
	}

}

func (l *Lexer) isAlpha(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_'
}

func (l* Lexer) tokenizeIdent() (token.Token, error) {
	// alpha numeric
	for l.isAlpha(l.peek()) || l.isDigit(l.peek()) {
		l.advance()
	}

	ident := string(l.source[l.start: l.current])
	keyword := token.LookUpIdent(ident)

	return token.New(keyword, ident , "", l.line, l.column), nil
}

func (l* Lexer) tokenizeString() (token.Token, error) {
	for l.peek() != '"' && !l.IsAtEnd() {
		l.advance()
	}

	if l.IsAtEnd() {	
		return token.New(token.EOF, "", "", l.line, l.column), fmt.Errorf("[line %v] Error: Unterminated string.\n", l.line)
	}

	l.advance()
	// trimming the surrounding quotes
	literal := string(l.source[l.start + 1: l.current - 1])
	return token.New(token.STR_LITERAL, fmt.Sprintf("\"%s\"", literal), literal, l.line, l.column), nil
}

func (l *Lexer) isDigit(c rune) bool {
	return c >= '0' && c <='9'
}

func (l* Lexer) tokenizeNumber() (token.Token, error) {
	for l.isDigit(l.peek()) {
		l.advance()
	}

	// look for any fractional part
	if l.peek() == '.' && l.isDigit(l.peekNext()) {
		// consume '.'
		l.advance()
		for l.isDigit(l.peek()) {
			l.advance()
		}
	}

	number := string(l.source[l.start: l.current])
	// convert into float before passing it as literal
	f, err := strconv.ParseFloat(number, 64)
	if err != nil {
		return token.New(token.ERROR, "", "", l.line, l.column), err
	}
	return token.New(token.NUM_LITERAL, number , f, l.line, l.column), nil
}


func (l *Lexer) NextToken() (token.Token, error) {
	l.skipWhitespace()
	l.start = l.current
	if l.IsAtEnd() {
		return token.New(token.EOF, "", "", l.line, l.column), nil
	}

	ch := l.advance()

	switch ch {
		// LITERALS
		case '"':
			return l.tokenizeString()	
		case '(': 
			return token.New(token.LPAREN, "(", "", l.line, l.column), nil
		case ')':
			return token.New(token.RPAREN, ")", "", l.line, l.column), nil
		case '{': 
			return token.New(token.LBRACE, "{", "", l.line, l.column), nil
		case '}':
			return token.New(token.RBRACE, "}", "", l.line, l.column), nil
		case ',':
			return token.New(token.COMMA, ",", "", l.line, l.column), nil
		case '.': 
			return token.New(token.DOT, ".", "", l.line, l.column), nil
		case '-': 
			return token.New(token.MINUS, "-", "", l.line, l.column), nil
		case '+': 
			return token.New(token.PLUS, "+", "", l.line, l.column), nil
		case ';': 
			return token.New(token.SEMICOLON, ";", "", l.line, l.column), nil
		case '/': 
			// single-line comment
			if l.peek() == '/' {
				l.advance()
				for !l.IsAtEnd() && l.peek() != '\n' {
					l.advance()
				}
				return l.NextToken()
			} 
			return token.New(token.DIV, "/", "", l.line, l.column), nil
		case '*':
			return token.New(token.STAR, "*", "", l.line, l.column), nil
		case '=':
			if l.peek() == '=' {
				l.advance()
				return token.New(token.EQ, "==", "", l.line, l.column), nil
			}
			return token.New(token.ASSIGN, "=", "", l.line, l.column), nil
		case '!':
			if l.peek() == '=' {
				l.advance()
				return token.New(token.NE, "!=", "", l.line, l.column), nil
			}
			return token.New(token.NOT, "!", "", l.line, l.column), nil
			
		case '>':
			if l.peek() == '=' {
				l.advance()
				return token.New(token.GE, ">=", "", l.line, l.column), nil
			}
			return token.New(token.GT, ">", "", l.line, l.column), nil
		case '<':
			if l.peek() == '=' {
				l.advance()
				return token.New(token.LE, "<=", "", l.line, l.column), nil
			}
			return token.New(token.LT, "<", "", l.line, l.column), nil
		default:
			if l.isDigit(ch) {
				return l.tokenizeNumber()
			}

			if l.isAlpha(ch) {
				return l.tokenizeIdent()
			}
			return token.New(token.ERROR, string(ch), "", l.line, l.column), fmt.Errorf("[line %v] Error: Unexpected character: %s\n", l.line, string(ch))
	}
}
