package expressions

import (
	env "TSwift/Classes/Env"
	utils "TSwift/Classes/Utils"
	"strconv"
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
		intValue, err := strconv.Atoi(pr.Value.(string))
		if err != nil {
		}
		return &utils.ReturnType{Value: intValue, Type: pr.Type}
	case utils.FLOAT:
		floatValue, err := strconv.ParseFloat(pr.Value.(string), 64)
		if err != nil {
		}
		return &utils.ReturnType{Value: floatValue, Type: pr.Type}
	case utils.BOOLEAN:
		return &utils.ReturnType{Value: pr.Value.(string) == "true", Type: pr.Type}
	case utils.STRING:
		pr.Value = strings.ReplaceAll(pr.Value.(string), "\\n", "\n")
		pr.Value = strings.ReplaceAll(pr.Value.(string), "\\t", "\t")
		pr.Value = strings.ReplaceAll(pr.Value.(string), "\\\"", "\"")
		pr.Value = strings.ReplaceAll(pr.Value.(string), "\\'", "'")
		pr.Value = strings.ReplaceAll(pr.Value.(string), "\\\\", "\\")
		return &utils.ReturnType{Value: pr.Value.(string), Type: pr.Type}
	default:
		return &utils.ReturnType{Value: "nil", Type: pr.Type}
	}
}
