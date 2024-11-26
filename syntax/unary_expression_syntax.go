package syntax

import "github.com/techerfan/expression/contracts"

type unaryExpressionSyntax struct {
	operatorToken *contracts.SyntaxToken
	operand       contracts.ExpressionSyntax
}

func NewUnaryExpressionSyntax(
	operatorToken *contracts.SyntaxToken,
	operand contracts.ExpressionSyntax,
) unaryExpressionSyntax {
	return unaryExpressionSyntax{
		operatorToken: operatorToken,
		operand:       operand,
	}
}

func (u unaryExpressionSyntax) GetChildren() []contracts.SyntaxNode {
	return []contracts.SyntaxNode{
		u.operatorToken,
		u.operand,
	}
}

func (u unaryExpressionSyntax) Kind() contracts.SyntaxKind {
	return contracts.UnaryExpression
}

func (u unaryExpressionSyntax) OperatorToken() *contracts.SyntaxToken {
	return u.operatorToken
}

func (u unaryExpressionSyntax) Operand() contracts.ExpressionSyntax {
	return u.operand
}
