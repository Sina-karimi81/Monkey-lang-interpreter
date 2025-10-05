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

// readIdentifier reads the input until it reaches a non letter chars and then slices the input
func (l *Lexer) readIdentifier() string {
	position := l.currentPosition
	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[position:l.currentPosition]
}

// readNumber reads the input until it reaches a non digit chars and then slices the input
func (l *Lexer) readNumber() string {
	position := l.currentPosition
	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.currentPosition]
}

// skipWhiteSpace skips white spaces characters so we can parse actual characters
func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) peekChar() byte {
	if l.currentPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.nextPosition]
	}
}

// NextToken look at New to find out why we have read the character as without reading the first character
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhiteSpace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{
				Type:    token.EQ,
				Literal: string(ch) + string(l.ch),
			}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{
				Type:    token.NotEq,
				Literal: string(ch) + string(l.ch),
			}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			identifier := l.readIdentifier()
			tok.Literal = identifier
			tok.Type = token.LookupIdent(identifier)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

// isLetter checks whether a character is allowed in an identifier or not
func isLetter(character byte) bool {
	return 'a' <= character && character <= 'z' ||
		'A' <= character && character <= 'Z' ||
		character == '_'
}

func isDigit(character byte) bool {
	return '0' <= character && character <= '9'
}

func newToken(tt token.TokenType, ch byte) token.Token {
	return token.Token{Type: tt, Literal: string(ch)}
}
