package contracts

import "reflect"

type VariableSymbol struct {
	Name string
	Type reflect.Kind
}

func NewVariableSymbol(name string, t reflect.Kind) *VariableSymbol {
	return &VariableSymbol{
		Name: name,
		Type: t,
	}
}
