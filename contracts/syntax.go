package contracts

type (
	SyntaxNode interface {
		GetChildren() []SyntaxNode
		Kind() SyntaxKind
	}

	ExpressionSyntax interface {
		SyntaxNode
	}

	AssignmentExpressionSyntax interface {
		ExpressionSyntax
		IdentifierToken() *SyntaxToken
		EqualsToken() *SyntaxToken
		Expression() ExpressionSyntax
	}

	BinaryExpressionSyntax interface {
		ExpressionSyntax
		OperatorToken() *SyntaxToken
		Right() ExpressionSyntax
		Left() ExpressionSyntax
	}

	LiteralExpressionSyntax interface {
		ExpressionSyntax
		LiteralToken() *SyntaxToken
		Value() interface{}
	}

	NameExpressionSyntax interface {
		ExpressionSyntax
		IdentifierToken() *SyntaxToken
	}

	ParanthesizedExpressionSyntax interface {
		ExpressionSyntax
		OpenParanthesisToken() *SyntaxToken
		Expression() ExpressionSyntax
		CloseParanthesisToken() *SyntaxToken
	}

	UnaryExpressionSyntax interface {
		ExpressionSyntax
		OperatorToken() *SyntaxToken
		Operand() ExpressionSyntax
	}
)
