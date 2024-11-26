package expression_test

import (
	"reflect"
	"testing"

	"github.com/techerfan/expression"
	"github.com/techerfan/expression/contracts"
	"github.com/techerfan/expression/syntax"

	"github.com/stretchr/testify/assert"
)

type tc struct {
	expression string
	result     interface{}
}

var testCases = []tc{
	{expression: "1 + 2", result: 3.0},
	{expression: "15 - 9", result: 6.0},
	{expression: "2 * 3", result: 6.0},
	{expression: "6 / 2", result: 3.0},
	{expression: "1 + 2 * 3", result: 7.0},
	{expression: "(1 + 2) * 3", result: 9.0},
	{expression: "true && false", result: false},
	{expression: "true || false", result: true},
	{expression: "!true", result: false},
	{expression: "!false", result: true},
	{expression: "1 == 2", result: false},
	{expression: "1 != 2", result: true},
	{expression: "true != false", result: true},
	{expression: "true == false", result: false},
	{expression: "test1 - test2", result: -5.0},
	{expression: "test1 + test2", result: 25.0},
	{expression: "test3 = 3", result: 3.0},
	{expression: "test3 + 2", result: 5.0},
	{expression: "test3 - 2", result: 1.0},
	{expression: "test3 * test1", result: 30.0},
	{expression: "test3 && test1", result: true},
	{expression: "test3 && 0", result: false},
	{expression: "0 || 0", result: false},
	{expression: "test3 || 0", result: true},
	{expression: "0 || test3", result: true},
	{expression: "test3 / 10", result: 0.3},
	{expression: "test3", result: 3.0},
	{expression: "4 << 1", result: 8},
	{expression: "5 >> 1", result: 2},
	{expression: "5 % 3", result: 2},
	{expression: "5 > 5", result: false},
	{expression: "5 < 5", result: false},
	{expression: "5 < 4", result: false},
	{expression: "5 >= 5", result: true},
	{expression: "5 >= 6", result: false},
	{expression: "5 <= 5", result: true},
	{expression: "5 <= 6", result: true},
	{expression: "5 | 8", result: 13},
	{expression: "5 & 4", result: 4},
	{expression: "5 ^ 12", result: 9},
}

var variables = map[*contracts.VariableSymbol]interface{}{
	contracts.NewVariableSymbol("test1", reflect.Float64): 10.0,
	contracts.NewVariableSymbol("test2", reflect.Float64): 15.0,
}

func TestEvaluationResult(t *testing.T) {
	for _, tc := range testCases {
		s := syntax.Parse(tc.expression)
		comp := expression.NewCompilation(s)
		result := comp.Evaluate(variables)
		assert.EqualValues(t, tc.result, result.Value)
	}
}

// Tests if variables are extracted correctly
func TestVariables(t *testing.T) {
	var testExpression = "(test1 + test2) * test3 + dnpm1_rtu_1_tag-1"
	var vars = []*struct {
		name string
		res  bool
	}{
		{name: "test1", res: false},
		{name: "test2", res: false},
		{name: "test3", res: false},
		{name: "dnpm1_rtu_1_tag", res: false},
	}

	s := syntax.Parse(testExpression)
	var sigs = s.Variables()
	for _, s := range sigs {
		for _, v := range vars {
			if s == v.name {
				v.res = true
			}
		}
	}

	for _, v := range vars {
		if !v.res {
			t.Fail()
		}
	}
}
