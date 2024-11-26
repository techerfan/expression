package syntax

import "github.com/techerfan/expression/contracts"

type parser struct {
	tokens      []*contracts.SyntaxToken
	position    int
	diagnostics *contracts.DiagnosticBag
}

func newParser(text string) *parser {
	tokens := []*contracts.SyntaxToken{}

	lexer := newLexer(text)

	var token *contracts.SyntaxToken

	for {
		token = lexer.nextToken()

		if token.Kind() != contracts.WhitespaceToken && token.Kind() != contracts.BadToken {
			tokens = append(tokens, token)
		}

		if token.Kind() == contracts.EndOfFileToken {
			break
		}
	}

	diagnostics := &contracts.DiagnosticBag{}
	diagnostics.AddRange(lexer.diagnostics)

	return &parser{
		tokens:      tokens,
		diagnostics: diagnostics,
	}
}

// Returns all the tokens used in the expression
func (p *parser) Tokens() []*contracts.SyntaxToken {
	return p.tokens
}

// Parses the expression
func (p *parser) Parse() *SyntaxTree {
	expression := p.parseExpression()
	endOfFileToken := p.matchToken(contracts.EndOfFileToken)
	return NewSyntaxTree(p.diagnostics.Diagnostics, expression, endOfFileToken, p.tokens)
}

// Starting point of the parser
func (p *parser) parseExpression() contracts.ExpressionSyntax {
	return p.parseAssignmentExpression()
}

func (p *parser) parseAssignmentExpression() contracts.ExpressionSyntax {
	if p.peek(0).Kind() == contracts.IdentifierToken && p.peek(1).Kind() == contracts.EqualsToken {
		identifierToken := p.nextToken()
		operatorToken := p.nextToken()
		right := p.parseAssignmentExpression()

		return NewAssignmentExpressionSyntax(identifierToken, operatorToken, right)
	}

	return p.parseBinaryExpression(0)
}

func (p *parser) parseBinaryExpression(parentPrecedence uint16) contracts.ExpressionSyntax {
	var left contracts.ExpressionSyntax
	unaryOperatorPrecedence := p.current().Kind().GetUnaryOperatorPrecedence()
	if unaryOperatorPrecedence != 0 && unaryOperatorPrecedence >= parentPrecedence {
		operatorToken := p.nextToken()
		operand := p.parseBinaryExpression(unaryOperatorPrecedence)
		left = NewUnaryExpressionSyntax(operatorToken, operand)
	} else {
		left = p.parsePrimaryExpression()
	}

	for {
		precedence := p.current().Kind().GetBinaryOperatorPrecedence()
		if precedence == 0 || precedence <= uint16(parentPrecedence) {
			break
		}

		operatorToken := p.nextToken()
		right := p.parseBinaryExpression(precedence)
		left = NewBinaryExpressionSyntax(left, operatorToken, right)
	}

	return left
}

func (p *parser) parsePrimaryExpression() contracts.ExpressionSyntax {
	switch p.current().Kind() {
	case contracts.OpenParanthesisToken:
		left := p.nextToken()
		expression := p.parseExpression()
		right := p.matchToken(contracts.CloseParanthesisToken)
		return NewParanthesizedExpressionSyntax(left, expression, right)

	case contracts.TrueKeyword:
		fallthrough
	case contracts.FalseKeyword:
		keywordToken := p.nextToken()
		var value bool = false
		if keywordToken.Kind() == contracts.TrueKeyword {
			value = true
		}
		return NewLiteralExpressionSyntaxWithValue(keywordToken, value)
	case contracts.IdentifierToken:
		identifierToken := p.nextToken()
		return NewNameExpressionSyntax(identifierToken)
	default:
		numberToken := p.matchToken(contracts.NumberToken)
		return NewLiteralExpressionSyntax(numberToken)
	}
}

func (p *parser) nextToken() *contracts.SyntaxToken {
	current := p.current()
	p.position += 1
	return current
}

func (p *parser) matchToken(kind contracts.SyntaxKind) *contracts.SyntaxToken {
	if p.current().Kind() == kind {
		return p.nextToken()
	}

	p.diagnostics.ReportUnexpectedToken(p.current().Span, p.current().Kind(), kind)
	return contracts.NewSyntaxToken(kind, p.current().Position, "", nil)
}

func (p *parser) peek(offset int) *contracts.SyntaxToken {
	index := p.position + offset

	if index >= len(p.tokens) {
		return p.tokens[len(p.tokens)-1]
	}
	return p.tokens[index]
}

func (p *parser) current() *contracts.SyntaxToken {
	return p.peek(0)
}

func (p *parser) lookahead() *contracts.SyntaxToken {
	return p.peek(1)
}
