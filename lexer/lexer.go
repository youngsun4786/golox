package lexer

import (
	"github.com/codecrafters-io/interpreter-starter-go/token"
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

func (l *Lexer) isAtEnd() bool {
	return l.current >= len(l.source)
}

func (l *Lexer) advance() rune {
	ch := rune(l.source[l.current])
	l.current++
	l.column++
	if ch == '\n' {
		l.resetCursor()
	}
	return ch
}

func (l *Lexer) peek() rune {
	if l.isAtEnd() {
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
		l.advance()
	}

}

func (l *Lexer) NextToken() token.Token {
	l.skipWhitespace()
	l.start = l.current
	if l.isAtEnd() {
		return token.New(token.EOF, "", "", l.line, l.column)
	}

	ch := l.advance()

	// comments
	if ch == '/' && l.peek() == '/' {
		ch = l.advance()
		for !l.isAtEnd() && ch != '\n' {
			l.advance()
			ch = l.peek()
		}
		return l.NextToken()
	}

	switch ch {
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



