package binder

import (
	"reflect"

	"github.com/techerfan/expression/contracts"
)

type BoundBinaryOperator struct {
	SyntaxKind contracts.SyntaxKind
	Kind       BoundBinaryOperatorKind
	LeftType   reflect.Kind
	RightType  reflect.Kind
	ResultType reflect.Kind
}

var binaryOperators = []*BoundBinaryOperator{
	NewBoundBinaryOperatorWithOneType(contracts.PlusToken, Addition, reflect.Float64),
	NewBoundBinaryOperatorWithOneType(contracts.MinusToken, Subtraction, reflect.Float64),
	NewBoundBinaryOperatorWithOneType(contracts.StarToken, Multiplication, reflect.Float64),
	NewBoundBinaryOperatorWithOneType(contracts.SlashToken, Division, reflect.Float64),
	NewBoundBinaryOperatorWithOneType(contracts.AmpersandToken, BitwiseAnd, reflect.Float64),
	NewBoundBinaryOperatorWithOneType(contracts.PipeToken, BitwiseOr, reflect.Float64),
	NewBoundBinaryOperatorWithOneType(contracts.GreaterGreaterToken, ShiftRight, reflect.Float64),
	NewBoundBinaryOperatorWithOneType(contracts.LesserLesserToken, ShiftLeft, reflect.Float64),
	NewBoundBinaryOperatorWithOneType(contracts.CaretToken, Xor, reflect.Float64),
	NewBoundBinaryOperatorWithOneType(contracts.PercentToken, Mod, reflect.Float64),

	NewBoundBinaryOperatorWithOperandType(contracts.EqualsEqualsToken, Equals, reflect.Float64, reflect.Bool),
	NewBoundBinaryOperatorWithOperandType(contracts.BangEqualsToken, NotEquals, reflect.Float64, reflect.Bool),
	NewBoundBinaryOperatorWithOperandType(contracts.GreaterToken, Greater, reflect.Float64, reflect.Bool),
	NewBoundBinaryOperatorWithOperandType(contracts.LesserToken, Lesser, reflect.Float64, reflect.Bool),
	NewBoundBinaryOperatorWithOperandType(contracts.GreaterOrEqualsToken, GreaterOrEquals, reflect.Float64, reflect.Bool),
	NewBoundBinaryOperatorWithOperandType(contracts.LesserOrEqualsToken, LesserOrEquals, reflect.Float64, reflect.Bool),

	NewBoundBinaryOperatorWithOneType(contracts.AmpersandAmpersandToken, LogicalAnd, reflect.Bool),
	NewBoundBinaryOperatorWithOperandType(contracts.AmpersandAmpersandToken, LogicalAnd, reflect.Float64, reflect.Bool),
	NewBoundBinaryOperator(contracts.AmpersandAmpersandToken, LogicalAnd, reflect.Float64, reflect.Bool, reflect.Bool),
	NewBoundBinaryOperator(contracts.AmpersandAmpersandToken, LogicalAnd, reflect.Bool, reflect.Float64, reflect.Bool),

	NewBoundBinaryOperatorWithOneType(contracts.PipePipeToken, LogicalOr, reflect.Bool),
	NewBoundBinaryOperatorWithOperandType(contracts.PipePipeToken, LogicalOr, reflect.Float64, reflect.Bool),
	NewBoundBinaryOperator(contracts.PipePipeToken, LogicalOr, reflect.Float64, reflect.Bool, reflect.Bool),
	NewBoundBinaryOperator(contracts.PipePipeToken, LogicalOr, reflect.Bool, reflect.Float64, reflect.Bool),

	NewBoundBinaryOperatorWithOneType(contracts.EqualsEqualsToken, Equals, reflect.Bool),
	NewBoundBinaryOperatorWithOneType(contracts.BangEqualsToken, NotEquals, reflect.Bool),
}

func NewBoundBinaryOperator(
	syntaxKind contracts.SyntaxKind,
	kind BoundBinaryOperatorKind,
	leftType reflect.Kind,
	rightType reflect.Kind,
	resulttype reflect.Kind,
) *BoundBinaryOperator {

	return &BoundBinaryOperator{
		SyntaxKind: syntaxKind,
		Kind:       kind,
		RightType:  rightType,
		LeftType:   leftType,
		ResultType: resulttype,
	}
}

func NewBoundBinaryOperatorWithOperandType(
	syntaxKind contracts.SyntaxKind,
	kind BoundBinaryOperatorKind,
	operandType reflect.Kind,
	resultType reflect.Kind,
) *BoundBinaryOperator {
	return NewBoundBinaryOperator(syntaxKind, kind, operandType, operandType, resultType)
}

func NewBoundBinaryOperatorWithOneType(
	syntaxKind contracts.SyntaxKind,
	kind BoundBinaryOperatorKind,
	t reflect.Kind,
) *BoundBinaryOperator {
	return NewBoundBinaryOperator(syntaxKind, kind, t, t, t)
}

func BindBoundBinaryOperator(
	kind contracts.SyntaxKind,
	leftType reflect.Kind,
	rightType reflect.Kind,
) *BoundBinaryOperator {
	for _, op := range binaryOperators {
		if op.SyntaxKind == kind && op.LeftType == leftType && op.RightType == rightType {
			return op
		}
	}
	return nil
}
