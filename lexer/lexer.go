package lexer

import (
	"github.com/youngsun4786/golox/token"
)

type Lexer struct {
	source   string
	filename string
	start    int
	current  int
	line     uint32
	column   uint32
}


func NewLexer(source string, filename string) *Lexer {
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
	return nil
}



