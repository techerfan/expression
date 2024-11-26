package contracts

import "reflect"

type (
	BoundNode interface {
		Kind() BoundNodeKind
	}

	BoundExpression interface {
		BoundNode
		Type() reflect.Kind
	}
)
