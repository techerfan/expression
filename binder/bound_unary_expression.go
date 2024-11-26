package binder

import (
	"reflect"

	"github.com/techerfan/expression/contracts"
)

type BoundUnaryExpression struct {
	Op      BoundUnaryOperator
	Operand contracts.BoundExpression
}

func NewBoundUnaryExpression(op BoundUnaryOperator, operand contracts.BoundExpression) BoundUnaryExpression {
	return BoundUnaryExpression{
		Op:      op,
		Operand: operand,
	}
}

func (b BoundUnaryExpression) Type() reflect.Kind {
	return b.Op.ResultType
}

func (b BoundUnaryExpression) Kind() contracts.BoundNodeKind {
	return contracts.UnaryExpressionNode
}
