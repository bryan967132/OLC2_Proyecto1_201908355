package env

import utils "TSwift/Classes/Utils"

type Symbol struct {
	Value   interface{}
	Id      string
	Type    utils.Type
	ArrType utils.Type
}

func NewSymbol(value interface{}, id string, Type, arrType utils.Type) *Symbol {
	return &Symbol{value, id, Type, arrType}
}
