package contracts

import (
	"fmt"
	"reflect"
)

type DiagnosticBag struct {
	Diagnostics []*Diagnostic
}

var diagnosticBagInstance *DiagnosticBag

func GetDiangnosticBag() *DiagnosticBag {
	if diagnosticBagInstance == nil {
		diagnosticBagInstance = &DiagnosticBag{}
	}

	return diagnosticBagInstance
}

func (d *DiagnosticBag) AddRange(diagBag *DiagnosticBag) {
	if d.Diagnostics == nil {
		d.Diagnostics = []*Diagnostic{}
	}
	d.Diagnostics = append(d.Diagnostics, diagBag.Diagnostics...)
}

func (d *DiagnosticBag) ReportBadCharacter(position int, character rune) {
	span := NewTextSpan(position, 1)
	d.report(span, fmt.Sprintf("Bad character input: '%c'", character))
}

func (d *DiagnosticBag) ReportInvalidNumber(
	span *TextSpan,
	text string,
	t reflect.Kind,
) {
	d.report(span, fmt.Sprintf("The number %s is not valid %s.", text, t.String()))
}

func (d *DiagnosticBag) ReportUndefinedUnaryOperator(
	span *TextSpan,
	operatorText string,
	operandType reflect.Kind,
) {
	d.report(span, fmt.Sprintf("Unary operator '%s' is not defined for type '%s'.", operatorText, operandType.String()))
}

func (d *DiagnosticBag) ReportUnexpectedToken(
	span *TextSpan,
	actualKind SyntaxKind,
	expectedKind SyntaxKind,
) {
	d.report(span, fmt.Sprintf("Error: Unexpected token <%s>, expected <%s>.", GetSyntaxKindName(actualKind), GetSyntaxKindName(expectedKind)))
}

func (d *DiagnosticBag) ReportUndefinedBinaryOperator(
	span *TextSpan,
	operatorText string,
	leftType reflect.Kind,
	rightType reflect.Kind,
) {
	d.report(span, fmt.Sprintf("Binary operator %s is not defined for type %s and %s", operatorText, leftType.String(), rightType.String()))
}

func (d *DiagnosticBag) ReportUndefinedName(span *TextSpan, name string) {
	d.report(span, fmt.Sprintf("variable %s does not exist", name))
}

func (d *DiagnosticBag) report(span *TextSpan, message string) {
	d.Diagnostics = append(d.Diagnostics, &Diagnostic{
		Message: message,
		Span:    span,
	})
}
