package binder

type BoundBinaryOperatorKind int

const (
	Addition BoundBinaryOperatorKind = iota + 1
	Subtraction
	Multiplication
	Division
	LogicalAnd
	LogicalOr
	Equals
	NotEquals
	BitwiseAnd
	BitwiseOr
	ShiftLeft
	ShiftRight
	Xor
	Greater
	Lesser
	GreaterOrEquals
	LesserOrEquals
	Mod
)
