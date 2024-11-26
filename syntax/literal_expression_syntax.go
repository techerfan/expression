package syntax

import "github.com/techerfan/expression/contracts"

type literalExpressionSyntax struct {
	literalToken *contracts.SyntaxToken
	value        interface{}
}

func NewLiteralExpressionSyntax(literalToken *contracts.SyntaxToken) literalExpressionSyntax {
	return literalExpressionSyntax{
		literalToken: literalToken,
		value:        literalToken.Value,
	}
}

func NewLiteralExpressionSyntaxWithValue(literalToken *contracts.SyntaxToken, value interface{}) literalExpressionSyntax {
	return literalExpressionSyntax{
		literalToken: literalToken,
		value:        value,
	}
}

func (l literalExpressionSyntax) GetChildren() []contracts.SyntaxNode {
	return []contracts.SyntaxNode{
		l.literalToken,
	}
}

func (l literalExpressionSyntax) Kind() contracts.SyntaxKind {
	return contracts.LiteralExpression
}

func (l literalExpressionSyntax) LiteralToken() *contracts.SyntaxToken {
	return l.literalToken
}

func (l literalExpressionSyntax) Value() interface{} {
	return l.value
}
