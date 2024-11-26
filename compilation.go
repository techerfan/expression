package expression

import (
	"github.com/techerfan/expression/binder"
	"github.com/techerfan/expression/contracts"
	"github.com/techerfan/expression/syntax"
)

type Compilation struct {
	Syntax *syntax.SyntaxTree
}

func NewCompilation(syntax *syntax.SyntaxTree) *Compilation {
	return &Compilation{
		Syntax: syntax,
	}
}

func (c *Compilation) Evaluate(variables map[*contracts.VariableSymbol]interface{}) *EvalutaionResult {
	bndr := binder.NewBinder(variables)
	boundExpression := bndr.BindExpression(c.Syntax.Root)

	diagnostics := append(c.Syntax.Diagnostics, bndr.Diagnostic.Diagnostics...)
	if len(diagnostics) > 0 {
		return NewEvaluationResult(diagnostics, nil)
	}

	eval := newEvaluator(boundExpression, variables)
	value := eval.Evaluate()
	return NewEvaluationResult([]*contracts.Diagnostic{}, value)
}
