package binder

import (
	"github.com/techerfan/expression/contracts"
)

type Binder struct {
	Diagnostic *contracts.DiagnosticBag
	variables  map[*contracts.VariableSymbol]interface{}
}

func NewBinder(variables map[*contracts.VariableSymbol]interface{}) *Binder {
	return &Binder{
		Diagnostic: &contracts.DiagnosticBag{},
		variables:  variables,
	}
}

func (b *Binder) BindExpression(syntax contracts.ExpressionSyntax) contracts.BoundExpression {
	switch syntax.Kind() {
	case contracts.LiteralExpression:
		return b.BindLiteralExpression(syntax.(contracts.LiteralExpressionSyntax))
	case contracts.BinaryExpression:
		return b.BindBinaryExpression(syntax.(contracts.BinaryExpressionSyntax))
	case contracts.UnaryExpression:
		return b.BindUnaryExpression(syntax.(contracts.UnaryExpressionSyntax))
	case contracts.ParenthesizedExpression:
		return b.BindParenthesizedExpression(syntax.(contracts.ParanthesizedExpressionSyntax))
	case contracts.NameExpression:
		return b.BindNameExpression(syntax.(contracts.NameExpressionSyntax))
	case contracts.AssignmentExpression:
		return b.BindAssignmentExpression(syntax.(contracts.AssignmentExpressionSyntax))
	default:
		panic("Unexpected syntax")
	}
}

func (b *Binder) BindLiteralExpression(syntax contracts.LiteralExpressionSyntax) contracts.BoundExpression {
	var value interface{} = syntax.Value()
	if value == nil {
		value = 0
	}

	return NewBoundLiteralExpression(value)
}

func (b *Binder) BindBinaryExpression(syntax contracts.BinaryExpressionSyntax) contracts.BoundExpression {
	var boundLeft = b.BindExpression(syntax.Left())
	var boundRight = b.BindExpression(syntax.Right())
	var boundOperator = BindBoundBinaryOperator(syntax.OperatorToken().Kind(), boundLeft.Type(), boundRight.Type())

	if boundOperator == nil {
		b.Diagnostic.ReportUndefinedBinaryOperator(
			syntax.OperatorToken().Span,
			syntax.OperatorToken().Text,
			boundLeft.Type(),
			boundRight.Type(),
		)
		return boundLeft
	}

	return NewBoundBinaryExpression(boundLeft, *boundOperator, boundRight)
}

func (b *Binder) BindUnaryExpression(syntax contracts.UnaryExpressionSyntax) contracts.BoundExpression {
	var boundOperand = b.BindExpression(syntax.Operand())
	var boundOperator = BindBoundUnaryOperator(syntax.OperatorToken().Kind(), boundOperand.Type())

	if boundOperator == nil {
		b.Diagnostic.ReportUndefinedUnaryOperator(syntax.OperatorToken().Span, syntax.OperatorToken().Text, boundOperand.Type())
		return boundOperand
	}

	return NewBoundUnaryExpression(*boundOperator, boundOperand)
}

func (b *Binder) BindParenthesizedExpression(syntax contracts.ParanthesizedExpressionSyntax) contracts.BoundExpression {
	return b.BindExpression(syntax.Expression())
}

func (b *Binder) BindNameExpression(syntax contracts.NameExpressionSyntax) contracts.BoundExpression {
	var name = syntax.IdentifierToken().Text

	var variable *contracts.VariableSymbol = nil
	for key := range b.variables {
		if key.Name == name {
			variable = key
		}
	}

	if variable == nil {
		b.Diagnostic.ReportUndefinedName(syntax.IdentifierToken().Span, name)
		return NewBoundLiteralExpression(0)
	}

	return NewBoundVariableExpression(variable)
}

func (b *Binder) BindAssignmentExpression(syntax contracts.AssignmentExpressionSyntax) contracts.BoundExpression {
	var name = syntax.IdentifierToken().Text
	var boundExpression = b.BindExpression(syntax.Expression())

	var existingVariable *contracts.VariableSymbol = nil
	for key := range b.variables {
		if key.Name == name {
			existingVariable = key
		}
	}

	if existingVariable != nil {
		delete(b.variables, existingVariable)
	}

	var variable = contracts.NewVariableSymbol(name, boundExpression.Type())
	b.variables[variable] = nil

	return NewBoundAssignmentExpression(variable, boundExpression)
}
