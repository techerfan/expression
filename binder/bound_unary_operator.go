package binder

import (
	"reflect"

	"github.com/techerfan/expression/contracts"
)

type BoundUnaryOperator struct {
	SyntaxKind  contracts.SyntaxKind
	Kind        BoundUnaryOperatorKind
	OperandType reflect.Kind
	ResultType  reflect.Kind
}

var unaryOperators = []*BoundUnaryOperator{
	NewBoundUnaryOperatorWithOperandType(contracts.BangToken, LogicalNegation, reflect.Bool),
	NewBoundUnaryOperatorWithOperandType(contracts.PlusToken, Identity, reflect.Bool),
	NewBoundUnaryOperatorWithOperandType(contracts.MinusToken, Negation, reflect.Bool),
}

func NewBoundUnaryOperatorWithOperandType(
	syntaxKind contracts.SyntaxKind,
	kind BoundUnaryOperatorKind,
	operandType reflect.Kind,
) *BoundUnaryOperator {
	return NewBoundUnaryOperator(syntaxKind, kind, operandType, operandType)
}

func NewBoundUnaryOperator(
	syntaxKind contracts.SyntaxKind,
	kind BoundUnaryOperatorKind,
	operandType reflect.Kind,
	resultType reflect.Kind,
) *BoundUnaryOperator {
	return &BoundUnaryOperator{
		SyntaxKind:  syntaxKind,
		Kind:        kind,
		OperandType: operandType,
		ResultType:  resultType,
	}
}

func BindBoundUnaryOperator(kind contracts.SyntaxKind, operandType reflect.Kind) *BoundUnaryOperator {
	for _, op := range unaryOperators {
		if op.SyntaxKind == kind && op.OperandType == operandType {
			return op
		}
	}
	return nil
}
