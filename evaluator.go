package expression

import (
	"reflect"

	"github.com/techerfan/expression/binder"
	"github.com/techerfan/expression/contracts"
)

type evaluator struct {
	root      contracts.BoundExpression
	variables map[*contracts.VariableSymbol]interface{}
}

func newEvaluator(root contracts.BoundExpression, variables map[*contracts.VariableSymbol]interface{}) *evaluator {
	return &evaluator{
		root:      root,
		variables: variables,
	}
}

func (e *evaluator) Evaluate() interface{} {
	return e.evaluateExpression(e.root)
}

func (e *evaluator) evaluateExpression(node contracts.BoundExpression) interface{} {
	if n, ok := node.(binder.BoundLiteralExpression); ok {
		return n.Value
	}

	if v, ok := node.(binder.BoundVariableExpression); ok {
		return e.variables[v.Variable]
	}

	if a, ok := node.(binder.BoundAssignmentExpression); ok {
		value := e.evaluateExpression(a.Expression)
		e.variables[a.Variable] = value
		return value
	}

	if u, ok := node.(binder.BoundUnaryExpression); ok {
		operand := e.evaluateExpression(u.Operand)

		switch u.Op.Kind {
		case binder.Identity:
			return operand.(float64)
		case binder.Negation:
			return -(operand.(float64))
		case binder.LogicalNegation:
			return !(operand.(bool))
		default:
			panic("Unexpected unary operator")
		}
	}

	if b, ok := node.(binder.BoundBinaryExpression); ok {
		left := e.evaluateExpression(b.Left)
		right := e.evaluateExpression(b.Right)

		switch b.Op.Kind {
		case binder.Addition:
			return left.(float64) + right.(float64)
		case binder.Subtraction:
			return left.(float64) - right.(float64)
		case binder.Multiplication:
			return left.(float64) * right.(float64)
		case binder.Division:
			return left.(float64) / right.(float64)
		case binder.LogicalAnd:
			if reflect.TypeOf(left).Kind() == reflect.Bool && reflect.TypeOf(right).Kind() == reflect.Bool {
				return left.(bool) && right.(bool)
			} else if reflect.TypeOf(left).Kind() == reflect.Float64 && reflect.TypeOf(right).Kind() == reflect.Float64 {
				return (left.(float64) != 0) && (right.(float64) != 0)
			}
			return false
		case binder.LogicalOr:
			if reflect.TypeOf(left).Kind() == reflect.Bool && reflect.TypeOf(right).Kind() == reflect.Bool {
				return left.(bool) || right.(bool)
			} else if reflect.TypeOf(left).Kind() == reflect.Float64 && reflect.TypeOf(right).Kind() == reflect.Float64 {
				return (left.(float64) != 0) || (right.(float64) != 0)
			}
			return false
		case binder.Equals:
			if reflect.TypeOf(left).Kind() == reflect.Bool && reflect.TypeOf(right).Kind() == reflect.Bool {
				return left.(bool) == right.(bool)
			} else if reflect.TypeOf(left).Kind() == reflect.Float64 && reflect.TypeOf(right).Kind() == reflect.Float64 {
				return left.(float64) == right.(float64)
			}
			return false
		case binder.NotEquals:
			if reflect.TypeOf(left).Kind() == reflect.Bool && reflect.TypeOf(right).Kind() == reflect.Bool {
				return left.(bool) != right.(bool)
			} else if reflect.TypeOf(left).Kind() == reflect.Float64 && reflect.TypeOf(right).Kind() == reflect.Float64 {
				return left.(float64) != right.(float64)
			}
			return false
		case binder.BitwiseAnd:
			return int64(left.(float64)) & int64(right.(float64))
		case binder.BitwiseOr:
			return int64(left.(float64)) | int64(right.(float64))
		case binder.ShiftLeft:
			return int64(left.(float64)) << int64(right.(float64))
		case binder.ShiftRight:
			return int64(left.(float64)) >> int64(right.(float64))
		case binder.Xor:
			return int64(left.(float64)) ^ int64(right.(float64))
		case binder.Greater:
			return left.(float64) > right.(float64)
		case binder.Lesser:
			return left.(float64) < right.(float64)
		case binder.GreaterOrEquals:
			return left.(float64) >= right.(float64)
		case binder.LesserOrEquals:
			return left.(float64) <= right.(float64)
		case binder.Mod:
			return int64(left.(float64)) % int64(right.(float64))
		default:
			panic("Unexpected binary operator")
		}
	}

	panic("Unexpected node")
}
