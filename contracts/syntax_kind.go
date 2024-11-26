package contracts

type SyntaxKind int

const (
	BadToken SyntaxKind = iota + 1
	EndOfFileToken
	NumberToken
	WhitespaceToken
	PlusToken
	MinusToken
	SlashToken
	StarToken
	BangToken
	EqualsToken
	AmpersandAmpersandToken
	PipePipeToken
	EqualsEqualsToken
	BangEqualsToken
	OpenParanthesisToken
	CloseParanthesisToken
	IdentifierToken
	GreaterToken
	LesserToken
	GreaterOrEqualsToken
	LesserOrEqualsToken
	AmpersandToken
	PipeToken
	GreaterGreaterToken
	LesserLesserToken
	CaretToken
	PercentToken

	// Keywords
	FalseKeyword
	TrueKeyword

	// Expressions
	LiteralExpression
	NameExpression
	UnaryExpression
	BinaryExpression
	ParenthesizedExpression
	AssignmentExpression
)

func GetSyntaxKindName(kind SyntaxKind) string {
	switch kind {
	case BadToken:
		return "BadToken"
	case EndOfFileToken:
		return "EndOfFileToken"
	case NumberToken:
		return "NumberToken"
	case WhitespaceToken:
		return "WhitespaceToken"
	case PlusToken:
		return "PlusToken"
	case MinusToken:
		return "MinusToken"
	case SlashToken:
		return "SlashToken"
	case StarToken:
		return "StarToken"
	case BangToken:
		return "BangToken"
	case EqualsToken:
		return "EqualsToken"
	case AmpersandAmpersandToken:
		return "AmpersandAmpersandToken"
	case PipePipeToken:
		return "PipePipeToken"
	case EqualsEqualsToken:
		return "EqualsEqualsToken"
	case BangEqualsToken:
		return "BangEqualsToken"
	case OpenParanthesisToken:
		return "OpenParanthesisToken"
	case CloseParanthesisToken:
		return "CloseParanthesisToken"
	case IdentifierToken:
		return "IdentifierToken"
	case GreaterToken:
		return "GreaterToken"
	case LesserToken:
		return "LesserToken"
	case GreaterOrEqualsToken:
		return "GreaterOrEqualsToken"
	case LesserOrEqualsToken:
		return "LesserOrEqualsToken"
	case AmpersandToken:
		return "AmpersandToken"
	case PipeToken:
		return "PipeToken"
	case GreaterGreaterToken:
		return "GreaterGreaterToken"
	case LesserLesserToken:
		return "LesserLesserToken"
	case CaretToken:
		return "CaretToken"
	case PercentToken:
		return "PercentToken"
	case FalseKeyword:
		return "FalseKeyword"
	case TrueKeyword:
		return "TrueKeyword"
	case LiteralExpression:
		return "LiteralExpression"
	case NameExpression:
		return "NameExpression"
	case UnaryExpression:
		return "UnaryExpression"
	case BinaryExpression:
		return "BinaryExpression"
	case ParenthesizedExpression:
		return "ParenthesizedExpression"
	case AssignmentExpression:
		return "AssignmentExpression"
	}

	return ""
}
