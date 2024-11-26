package syntax

import "github.com/techerfan/expression/contracts"

type assignmentExpressionSyntax struct {
	identifierToken *contracts.SyntaxToken
	equalsToken     *contracts.SyntaxToken
	expression      contracts.ExpressionSyntax
}

func NewAssignmentExpressionSyntax(
	identifierToken *contracts.SyntaxToken,
	equalsToken *contracts.SyntaxToken,
	expression contracts.ExpressionSyntax,
) assignmentExpressionSyntax {
	return assignmentExpressionSyntax{
		identifierToken: identifierToken,
		equalsToken:     equalsToken,
		expression:      expression,
	}
}

func (a assignmentExpressionSyntax) GetChildren() []contracts.SyntaxNode {
	return []contracts.SyntaxNode{
		a.identifierToken,
		a.equalsToken,
		a.expression,
	}
}

func (a assignmentExpressionSyntax) Kind() contracts.SyntaxKind {
	return contracts.AssignmentExpression
}

func (a assignmentExpressionSyntax) IdentifierToken() *contracts.SyntaxToken {
	return a.identifierToken
}

func (a assignmentExpressionSyntax) EqualsToken() *contracts.SyntaxToken {
	return a.equalsToken
}

func (a assignmentExpressionSyntax) Expression() contracts.ExpressionSyntax {
	return a.expression
}
