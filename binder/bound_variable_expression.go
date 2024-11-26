package binder

import (
	"reflect"

	"github.com/techerfan/expression/contracts"
)

type BoundVariableExpression struct {
	Variable *contracts.VariableSymbol
}

func NewBoundVariableExpression(variable *contracts.VariableSymbol) BoundVariableExpression {
	return BoundVariableExpression{
		Variable: variable,
	}
}

func (b BoundVariableExpression) Type() reflect.Kind {
	return b.Variable.Type
}

func (b BoundVariableExpression) Kind() contracts.BoundNodeKind {
	return contracts.VariableExpressionNode
}
