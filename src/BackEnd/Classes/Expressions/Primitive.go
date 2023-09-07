package expressions

import (
	env "TSwift/Classes/Env"
	utils "TSwift/Classes/Utils"
	"fmt"
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

func (ar *Primitive) LineN() int {
	return ar.Line
}

func (ar *Primitive) ColumnN() int {
	return ar.Column
}

func (pr *Primitive) Exec(env *env.Env) *utils.ReturnType {
	switch pr.Type {
	case utils.INT:
		intValue, _ := strconv.Atoi(pr.Value.(string))
		return &utils.ReturnType{Value: intValue, Type: pr.Type}
	case utils.FLOAT:
		floatValue, _ := strconv.ParseFloat(pr.Value.(string), 64)
		return &utils.ReturnType{Value: floatValue, Type: pr.Type}
	case utils.BOOLEAN:
		return &utils.ReturnType{Value: pr.Value.(string) == "true", Type: pr.Type}
	case utils.NIL:
		return &utils.ReturnType{Value: "nil", Type: pr.Type}
	default:
		pr.Value = strings.ReplaceAll(fmt.Sprintf("%v", pr.Value), "\\n", "\n")
		pr.Value = strings.ReplaceAll(fmt.Sprintf("%v", pr.Value), "\\t", "\t")
		pr.Value = strings.ReplaceAll(fmt.Sprintf("%v", pr.Value), "\\r", "\r")
		pr.Value = strings.ReplaceAll(fmt.Sprintf("%v", pr.Value), "\\\"", "\"")
		pr.Value = strings.ReplaceAll(fmt.Sprintf("%v", pr.Value), "\\'", "'")
		pr.Value = strings.ReplaceAll(fmt.Sprintf("%v", pr.Value), "\\\\", "\\")
		return &utils.ReturnType{Value: fmt.Sprintf("%v", pr.Value), Type: pr.Type}
	}
}
