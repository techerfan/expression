package binder

type BoundUnaryOperatorKind int

const (
	Identity BoundUnaryOperatorKind = iota + 1
	Negation
	LogicalNegation
)
