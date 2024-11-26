package binder

import (
	"reflect"

	"github.com/techerfan/expression/contracts"
)

type BoundLiteralExpression struct {
	Value interface{}
}

func NewBoundLiteralExpression(value interface{}) BoundLiteralExpression {
	return BoundLiteralExpression{
		Value: value,
	}
}

func (b BoundLiteralExpression) Type() reflect.Kind {
	return reflect.TypeOf(b.Value).Kind()
}

func (b BoundLiteralExpression) Kind() contracts.BoundNodeKind {
	return contracts.LiteralExpressionNode
}
