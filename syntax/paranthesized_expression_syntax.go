package syntax

import "github.com/techerfan/expression/contracts"

type paranthesizedExpressionSyntax struct {
	openParanthesisToken  *contracts.SyntaxToken
	expression            contracts.ExpressionSyntax
	closeParanthesisToken *contracts.SyntaxToken
}

func NewParanthesizedExpressionSyntax(
	OpenParanthesisToken *contracts.SyntaxToken,
	expression contracts.ExpressionSyntax,
	closeParanthesisToken *contracts.SyntaxToken,
) paranthesizedExpressionSyntax {
	return paranthesizedExpressionSyntax{
		openParanthesisToken:  OpenParanthesisToken,
		expression:            expression,
		closeParanthesisToken: closeParanthesisToken,
	}
}

func (p paranthesizedExpressionSyntax) GetChildren() []contracts.SyntaxNode {
	return []contracts.SyntaxNode{
		p.openParanthesisToken,
		p.expression,
		p.closeParanthesisToken,
	}
}

func (p paranthesizedExpressionSyntax) Kind() contracts.SyntaxKind {
	return contracts.ParenthesizedExpression
}

func (p paranthesizedExpressionSyntax) OpenParanthesisToken() *contracts.SyntaxToken {
	return p.openParanthesisToken
}

func (p paranthesizedExpressionSyntax) Expression() contracts.ExpressionSyntax {
	return p.expression
}

func (p paranthesizedExpressionSyntax) CloseParanthesisToken() *contracts.SyntaxToken {
	return p.closeParanthesisToken
}
