package env

import utils "TSwift/Classes/Utils"

type Symbol struct {
	Value   interface{}
	id      string
	Type    utils.Type
	arrType utils.Type
}

func NewSymbol(value interface{}, id string, Type, arrType utils.Type) *Symbol {
	return &Symbol{value, id, Type, arrType}
}
