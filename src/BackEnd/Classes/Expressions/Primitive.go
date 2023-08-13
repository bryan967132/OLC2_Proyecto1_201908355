package expressions

import (
	env "TSwift/Classes/Env"
	utils "TSwift/Classes/Utils"
	"strings"
)

type Primitive struct {
	Line    int
	Column  int
	TypeExp utils.TypeExp
	Value   interface{}
	Type    utils.Type
}

func NewPrimitive(line, column int, value interface{}, typeD utils.Type) *Primitive {
	return &Primitive{line, column, utils.PRIMITIVE, value, typeD}
}

func (pr *Primitive) Exec(env *env.Env) *utils.ReturnType {
	switch pr.Type {
	case utils.INT:
		return &utils.ReturnType{Value: pr.Value.(int), Type: pr.Type}
	case utils.FLOAT:
		return &utils.ReturnType{Value: pr.Value.(float64), Type: pr.Type}
	case utils.BOOLEAN:
		return &utils.ReturnType{Value: pr.Value.(bool), Type: pr.Type}
	default:
		pr.Value = strings.ReplaceAll(pr.Value.(string), "\\n", "\n")
		pr.Value = strings.ReplaceAll(pr.Value.(string), "\\t", "\t")
		pr.Value = strings.ReplaceAll(pr.Value.(string), "\\\"", "\"")
		pr.Value = strings.ReplaceAll(pr.Value.(string), "\\'", "'")
		pr.Value = strings.ReplaceAll(pr.Value.(string), "\\\\", "\\")
		return &utils.ReturnType{Value: pr.Value.(string), Type: pr.Type}
	}
}
