package lexer

import "Monkey-Lang/token"

// Lexer this represents an actual lexer implementation, with a lookahead of 1 with the means of currentPosition and nextPosition
type Lexer struct {
	input           string
	currentPosition int  // points to current char
	nextPosition    int  // the position after the current char
	ch              byte // current char
}

// New creates a lexer and reads one character into memory
func New(input string) *Lexer {
	l := &Lexer{
		input: input,
	}
	l.readChar() // to make sure we have read one character

	return l
}

// readChar function reads one char at a time by comparing the current position of the lexer and the size of input
func (l *Lexer) readChar() {
	if l.nextPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.nextPosition]
	}
	l.currentPosition = l.nextPosition
	l.nextPosition++ // incrementing the lookahead
}

// NextToken look at New to find out why we have read the character as without reading the first character
func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	l.readChar()
	return tok
}

func newToken(tt token.TokenType, ch byte) token.Token {
	return token.Token{Type: tt, Literal: string(ch)}
}
