package contracts

type SyntaxToken struct {
	Position int
	Text     string
	SynKind  SyntaxKind
	Value    interface{}
	Span     *TextSpan
}

func NewSyntaxToken(kind SyntaxKind, position int, text string, value interface{}) *SyntaxToken {
	return &SyntaxToken{
		SynKind:  kind,
		Position: position,
		Text:     text,
		Value:    value,
	}
}

func (s *SyntaxToken) Kind() SyntaxKind {
	return s.SynKind
}

func (s *SyntaxToken) GetChildren() []SyntaxNode {
	return nil
}
