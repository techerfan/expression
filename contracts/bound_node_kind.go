package contracts

type BoundNodeKind int

const (
	LiteralExpressionNode BoundNodeKind = iota + 1
	UnaryExpressionNode
	BinaryExpressionNode
	VariableExpressionNode
	AssignmentExpressionNode
)
