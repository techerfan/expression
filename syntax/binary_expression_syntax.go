package syntax

import (
	"github.com/techerfan/expression/contracts"
)

type binaryExpressionSyntax struct {
	operatorToken *contracts.SyntaxToken
	right         contracts.ExpressionSyntax
	left          contracts.ExpressionSyntax
}

func NewBinaryExpressionSyntax(
	left contracts.ExpressionSyntax,
	operatorToken *contracts.SyntaxToken,
	right contracts.ExpressionSyntax,
) binaryExpressionSyntax {
	return binaryExpressionSyntax{
		left:          left,
		operatorToken: operatorToken,
		right:         right,
	}
}

func (b binaryExpressionSyntax) GetChildren() []contracts.SyntaxNode {
	return []contracts.SyntaxNode{
		b.left,
		b.operatorToken,
		b.right,
	}
}

func (b binaryExpressionSyntax) Kind() contracts.SyntaxKind {
	return contracts.BinaryExpression
}

func (b binaryExpressionSyntax) OperatorToken() *contracts.SyntaxToken {
	return b.operatorToken
}

func (b binaryExpressionSyntax) Right() contracts.ExpressionSyntax {
	return b.right
}

func (b binaryExpressionSyntax) Left() contracts.ExpressionSyntax {
	return b.left
}
