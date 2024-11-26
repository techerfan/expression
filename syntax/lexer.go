package syntax

import (
	"reflect"
	"strconv"
	"unicode"

	"github.com/techerfan/expression/contracts"
)

type lexer struct {
	text        string
	position    int
	diagnostics *contracts.DiagnosticBag
}

func newLexer(text string) *lexer {
	return &lexer{
		text: text,
		diagnostics: &contracts.DiagnosticBag{
			Diagnostics: []*contracts.Diagnostic{},
		},
	}
}

func (l *lexer) nextToken() *contracts.SyntaxToken {
	if l.position >= len(l.text) {
		return contracts.NewSyntaxToken(contracts.EndOfFileToken, l.position, "0", nil)
	}

	start := l.position

	if unicode.IsDigit(l.current()) {
		hasFloatingPoint := false
		for unicode.IsDigit(l.current()) || (l.current() == '.' && !hasFloatingPoint) {
			if l.current() == '.' {
				hasFloatingPoint = true
			}
			l.next()
		}

		length := l.position - start
		text := l.text[start : start+length]

		if !hasFloatingPoint {
			if val, err := strconv.ParseFloat(text, 64); err != nil {
				l.diagnostics.ReportInvalidNumber(contracts.NewTextSpan(start, length), text, reflect.TypeOf(val).Kind())
			} else {
				return contracts.NewSyntaxToken(contracts.NumberToken, start, text, val)
			}
		} else {
			if val, err := strconv.ParseFloat(text, 64); err != nil {
				l.diagnostics.ReportInvalidNumber(contracts.NewTextSpan(start, length), text, reflect.TypeOf(val).Kind())
			} else {
				return contracts.NewSyntaxToken(contracts.NumberToken, start, text, val)
			}
		}
	}

	if unicode.IsSpace(l.current()) {
		for unicode.IsSpace(l.current()) {
			l.next()
		}

		length := l.position - start
		text := l.text[start : start+length]
		return contracts.NewSyntaxToken(contracts.WhitespaceToken, start, text, 0)
	}

	if unicode.IsLetter(l.current()) {
		for unicode.IsLetter(l.current()) || unicode.IsDigit(l.current()) || l.current() == '_' {
			l.next()
		}

		length := l.position - start
		text := l.text[start : start+length]
		kind := contracts.GetKeyworkKind(text)
		return contracts.NewSyntaxToken(kind, start, text, text)
	}

	pos := l.position
	switch l.current() {
	case '+':
		l.position += 1
		return contracts.NewSyntaxToken(contracts.PlusToken, pos, "+", nil)
	case '-':
		l.position += 1
		return contracts.NewSyntaxToken(contracts.MinusToken, pos, "-", nil)
	case '*':
		l.position += 1
		return contracts.NewSyntaxToken(contracts.StarToken, pos, "*", nil)
	case '/':
		l.position += 1
		return contracts.NewSyntaxToken(contracts.SlashToken, pos, "/", nil)
	case '(':
		l.position += 1
		return contracts.NewSyntaxToken(contracts.OpenParanthesisToken, pos, "(", nil)
	case ')':
		l.position += 1
		return contracts.NewSyntaxToken(contracts.CloseParanthesisToken, pos, ")", nil)
	case '%':
		l.position += 1
		return contracts.NewSyntaxToken(contracts.PercentToken, pos, "%", nil)
	case '^':
		l.position += 1
		return contracts.NewSyntaxToken(contracts.CaretToken, pos, "^", nil)
	case '>':
		if l.lookahead() == '>' {
			l.position += 2
			return contracts.NewSyntaxToken(contracts.GreaterGreaterToken, start, ">>", nil)
		} else if l.lookahead() == '=' {
			l.position += 2
			return contracts.NewSyntaxToken(contracts.GreaterOrEqualsToken, start, ">=", nil)
		}
		l.position += 1
		return contracts.NewSyntaxToken(contracts.GreaterToken, start, ">", nil)
	case '<':
		if l.lookahead() == '<' {
			l.position += 2
			return contracts.NewSyntaxToken(contracts.LesserLesserToken, start, "<<", nil)
		} else if l.lookahead() == '=' {
			l.position += 2
			return contracts.NewSyntaxToken(contracts.LesserOrEqualsToken, start, "<=", nil)
		}
		l.position += 1
		return contracts.NewSyntaxToken(contracts.LesserToken, start, "<", nil)
	case '&':
		if l.lookahead() == '&' {
			l.position += 2
			return contracts.NewSyntaxToken(contracts.AmpersandAmpersandToken, start, "&&", nil)
		}
		l.position += 1
		return contracts.NewSyntaxToken(contracts.AmpersandToken, start, "&", nil)
	case '|':
		if l.lookahead() == '|' {
			l.position += 2
			return contracts.NewSyntaxToken(contracts.PipePipeToken, start, "||", nil)
		}
		l.position += 1
		return contracts.NewSyntaxToken(contracts.PipeToken, start, "|", nil)
	case '=':
		if l.lookahead() == '=' {
			l.position += 2
			return contracts.NewSyntaxToken(contracts.EqualsEqualsToken, start, "==", nil)
		}
		l.position += 1
		return contracts.NewSyntaxToken(contracts.EqualsToken, start, "=", nil)
	case '!':
		if l.lookahead() == '=' {
			l.position += 2
			return contracts.NewSyntaxToken(contracts.BangEqualsToken, start, "!=", nil)
		}
		l.position += 1
		return contracts.NewSyntaxToken(contracts.BangToken, start, "!", nil)
	}

	l.diagnostics.ReportBadCharacter(l.position, l.current())
	l.position += 1
	return contracts.NewSyntaxToken(contracts.BadToken, pos, l.text[l.position-1:l.position], nil)
}

func (l *lexer) next() {
	l.position += 1
}

func (l *lexer) peek(offset int) rune {
	index := l.position + offset

	if index >= len(l.text) {
		return -1
	}
	return rune(l.text[index])
}

func (l *lexer) current() rune {
	return l.peek(0)
}

func (l *lexer) lookahead() rune {
	return l.peek(1)
}
