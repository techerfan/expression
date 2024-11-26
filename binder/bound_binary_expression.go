package binder

import (
	"reflect"

	"github.com/techerfan/expression/contracts"
)

type BoundBinaryExpression struct {
	Left  contracts.BoundExpression
	Op    BoundBinaryOperator
	Right contracts.BoundExpression
}

func NewBoundBinaryExpression(
	left contracts.BoundExpression,
	op BoundBinaryOperator,
	right contracts.BoundExpression,
) BoundBinaryExpression {
	return BoundBinaryExpression{
		Left:  left,
		Op:    op,
		Right: right,
	}
}

func (b BoundBinaryExpression) Type() reflect.Kind {
	return b.Op.ResultType
}

func (b BoundBinaryExpression) Kind() contracts.BoundNodeKind {
	return contracts.BinaryExpressionNode
}
