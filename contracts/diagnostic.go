package contracts

type Diagnostic struct {
	Message string
	Span    *TextSpan
}

func NewDiagnostic(message string, span *TextSpan) *Diagnostic {
	return &Diagnostic{
		Message: message,
		Span:    span,
	}
}
