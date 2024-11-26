package contracts

func GetKeyworkKind(text string) SyntaxKind {
	switch text {
	case "true":
		return TrueKeyword
	case "false":
		return FalseKeyword
	default:
		return IdentifierToken
	}
}

func (s SyntaxKind) GetUnaryOperatorPrecedence() uint16 {
	switch s {
	case PlusToken:
		fallthrough
	case MinusToken:
		fallthrough
	case BangToken:
		return 11

	default:
		return 0
	}
}

func (s SyntaxKind) GetBinaryOperatorPrecedence() uint16 {
	switch s {
	case StarToken:
		fallthrough
	case PercentToken:
		fallthrough
	case SlashToken:
		return 10

	case PlusToken:
		fallthrough
	case MinusToken:
		return 9

	case LesserLesserToken:
		fallthrough
	case GreaterGreaterToken:
		return 8

	case GreaterToken:
		fallthrough
	case GreaterOrEqualsToken:
		fallthrough
	case LesserToken:
		fallthrough
	case LesserOrEqualsToken:
		return 7

	case EqualsEqualsToken:
		fallthrough
	case BangEqualsToken:
		return 6

	case AmpersandToken:
		return 5
	case CaretToken:
		return 4
	case PipeToken:
		return 3
	case AmpersandAmpersandToken:
		return 2
	case PipePipeToken:
		return 1

	default:
		return 0
	}
}
