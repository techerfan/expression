package expression

import (
	"reflect"

	"github.com/techerfan/expression/contracts"
)

type EvalutaionResult struct {
	Diagnostics []*contracts.Diagnostic
	Value       interface{}
	// some logics may not accept boolean results,
	// hence, we need to cast them to float (false: 0 and true: 1).
	FloatCastedValue float64
}

func NewEvaluationResult(diagnostics []*contracts.Diagnostic, value interface{}) *EvalutaionResult {
	var floatValue float64 = 0
	if reflect.TypeOf(value).Kind() == reflect.Bool {
		if value.(bool) {
			floatValue = 1
		}
	} else if reflect.TypeOf(value).Kind() == reflect.Float64 {
		floatValue = value.(float64)
	}

	return &EvalutaionResult{
		Diagnostics:      diagnostics,
		Value:            value,
		FloatCastedValue: floatValue,
	}
}
