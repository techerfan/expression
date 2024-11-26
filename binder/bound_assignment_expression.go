package binder

import (
	"reflect"

	"github.com/techerfan/expression/contracts"
)

type BoundAssignmentExpression struct {
	Variable   *contracts.VariableSymbol
	Expression contracts.BoundExpression
}

func NewBoundAssignmentExpression(
	variable *contracts.VariableSymbol,
	boundExpression contracts.BoundExpression,
) BoundAssignmentExpression {
	return BoundAssignmentExpression{
		Variable:   variable,
		Expression: boundExpression,
	}
}

func (b BoundAssignmentExpression) Kind() contracts.BoundNodeKind {
	return contracts.AssignmentExpressionNode
}

func (b BoundAssignmentExpression) Type() reflect.Kind {
	return b.Variable.Type
}
