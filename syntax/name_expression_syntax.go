package syntax

import "github.com/techerfan/expression/contracts"

type nameExpressionSyntax struct {
	identifierToken *contracts.SyntaxToken
}

func NewNameExpressionSyntax(identifierToken *contracts.SyntaxToken) nameExpressionSyntax {
	return nameExpressionSyntax{
		identifierToken: identifierToken,
	}
}

func (n nameExpressionSyntax) GetChildren() []contracts.SyntaxNode {
	return []contracts.SyntaxNode{
		n.identifierToken,
	}
}

func (n nameExpressionSyntax) Kind() contracts.SyntaxKind {
	return contracts.NameExpression
}

func (n nameExpressionSyntax) IdentifierToken() *contracts.SyntaxToken {
	return n.identifierToken
}
