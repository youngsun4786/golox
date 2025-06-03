package lexer

import (
	"github.com/codecrafters-io/interpreter-starter-go/token"
	"fmt"
	"os"
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
func (l* Lexer) tokenizeString() token.Token {
	for l.peek() != '"' && !l.IsAtEnd() {
		l.advance()
	}

	if l.IsAtEnd() {
		fmt.Fprintf(os.Stderr, "[line %v] Error: Unterminated string.\n", l.line)
		return token.New(token.EOF, "", "", l.line, l.column)
	}

	l.advance()
	// trimming the surrounding quotes
	literal := string(l.source[l.start + 1: l.current - 1])
	return token.New(token.STR_LITERAL, fmt.Sprintf("\"%s\"", literal), literal, l.line, l.column)
}

func (l *Lexer) NextToken() token.Token {
	l.skipWhitespace()
	l.start = l.current
	if l.IsAtEnd() {
		return token.New(token.EOF, "", "", l.line, l.column)
	}

	ch := l.advance()

	switch ch {
		// LITERALS
		case '"':
			return l.tokenizeString()
		case '(': 
			return token.New(token.LPAREN, "(", "", l.line, l.column)
		case ')':
			return token.New(token.RPAREN, ")", "", l.line, l.column)
		case '{': 
			return token.New(token.LBRACE, "{", "", l.line, l.column)
		case '}':
			return token.New(token.RBRACE, "}", "", l.line, l.column)
		case ',':
			return token.New(token.COMMA, ",", "", l.line, l.column)
		case '.': 
			return token.New(token.DOT, ".", "", l.line, l.column)
		case '-': 
			return token.New(token.MINUS, "-", "", l.line, l.column)
		case '+': 
			return token.New(token.PLUS, "+", "", l.line, l.column)	
		case ';': 
			return token.New(token.SEMICOLON, ";", "", l.line, l.column)
		case '/': 
			// single-line comment
			if l.peek() == '/' {
				l.advance()
				for !l.IsAtEnd() && l.peek() != '\n' {
					l.advance()
				}
				return l.NextToken()
			} 
			return token.New(token.DIV, "/", "", l.line, l.column)
		case '*':
			return token.New(token.STAR, "*", "", l.line, l.column)
		case '=':
			if l.peek() == '=' {
				l.advance()
				return token.New(token.EQ, "==", "", l.line, l.column)
			}
			return token.New(token.ASSIGN, "=", "", l.line, l.column)
		case '!':
			if l.peek() == '=' {
				l.advance()
				return token.New(token.NE, "!=", "", l.line, l.column)
			}
			return token.New(token.NOT, "!", "", l.line, l.column)
			
		case '>':
			if l.peek() == '=' {
				l.advance()
				return token.New(token.GE, ">=", "", l.line, l.column)
			}
			return token.New(token.GT, ">", "", l.line, l.column)
		case '<':
			if l.peek() == '=' {
				l.advance()
				return token.New(token.LE, "<=", "", l.line, l.column)
			}
			return token.New(token.LT, "<", "", l.line, l.column)
		default:
			return token.New(token.ERROR, string(ch), "", l.line, l.column)	
	}
}
