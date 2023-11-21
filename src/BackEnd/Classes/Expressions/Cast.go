package expressions

import (
	env "TSwift/Classes/Env"
	interfaces "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Cast struct {
	Line    int
	Column  int
	Destiny utils.Type
	TypeExp utils.TypeExp
	Exp     interfaces.Expression
}

func NewCast(line, column int, destiny utils.Type, exp interfaces.Expression) *Cast {
	return &Cast{line, column, destiny, utils.CAST, exp}
}

func (ct *Cast) LineN() int {
	return ct.Line
}

func (ct *Cast) ColumnN() int {
	return ct.Column
}

func (ct *Cast) Exec(env *env.Env) *utils.ReturnType {
	value := ct.Exp.Exec(env)
	if ct.Destiny == utils.INT {
		if value.Type == utils.FLOAT {
			floatValue, _ := strconv.ParseFloat(fmt.Sprintf("%v", value.Value), 64)
			return &utils.ReturnType{Value: int(floatValue), Type: ct.Destiny}
		}
		if value.Type == utils.STRING || value.Type == utils.CHAR {
			if ct.isValidNumericalString(fmt.Sprintf("%v", value.Value)) {
				floatValue, _ := strconv.ParseFloat(fmt.Sprintf("%v", value.Value), 64)
				return &utils.ReturnType{Value: int(floatValue), Type: ct.Destiny}
			}
			env.SetError(fmt.Sprintf("La cadena \"%s\" no tiene formato numérico para castear a \"Int\"", value.Value), ct.Exp.LineN(), ct.Exp.ColumnN())
			return &utils.ReturnType{Value: "nil", Type: utils.NIL}
		}
		env.SetError(fmt.Sprintf("No hay casteo de \"%s\" a \"Int\"", ct.getType(value.Type)), ct.Exp.LineN(), ct.Exp.ColumnN())
		return &utils.ReturnType{Value: "nil", Type: utils.NIL}
	}
	if ct.Destiny == utils.FLOAT {
		if value.Type == utils.STRING {
			if ct.isValidNumericalString(fmt.Sprintf("%v", value.Value)) {
				floatValue, _ := strconv.ParseFloat(fmt.Sprintf("%v", value.Value), 64)
				return &utils.ReturnType{Value: floatValue, Type: ct.Destiny}
			}
			env.SetError(fmt.Sprintf("La cadena \"%s\" no tiene formato numérico para castear a \"Int\"", value.Value), ct.Exp.LineN(), ct.Exp.ColumnN())
			return &utils.ReturnType{Value: "nil", Type: utils.NIL}
		}
		env.SetError(fmt.Sprintf("No hay casteo de \"%s\" a \"Float\"", ct.getType(value.Type)), ct.Exp.LineN(), ct.Exp.ColumnN())
		return &utils.ReturnType{Value: "nil", Type: utils.NIL}
	}
	if ct.Destiny == utils.STRING {
		if value.Type == utils.INT {
			return &utils.ReturnType{Value: fmt.Sprintf("%v", value.Value), Type: ct.Destiny}
		}
		if value.Type == utils.FLOAT {
			return &utils.ReturnType{Value: ct.cutDecimals(fmt.Sprintf("%v", value.Value)), Type: ct.Destiny}
		}
		if value.Type == utils.BOOLEAN {
			return &utils.ReturnType{Value: fmt.Sprintf("%v", value.Value), Type: ct.Destiny}
		}
		env.SetError(fmt.Sprintf("No hay casteo de \"%s\" a \"String\"", ct.getType(value.Type)), ct.Exp.LineN(), ct.Exp.ColumnN())
		return &utils.ReturnType{Value: "nil", Type: utils.NIL}
	}
	env.SetError("Error: No hay casteo a \"Bool\" y \"Character\"", ct.Line, ct.Column)
	return &utils.ReturnType{Value: "nil", Type: utils.NIL}
}

func (ct *Cast) isValidNumericalString(num string) bool {
	expresionRegular := `^\d+(\.\d+)?$`
	regexpValidador := regexp.MustCompile(expresionRegular)
	return regexpValidador.MatchString(num)
}

func (ct *Cast) cutDecimals(num string) string {
	parts := strings.Split(num, ".")
	if len(parts) == 2 {
		if len(parts[1]) > 5 {
			return parts[0] + "." + parts[1][:5]
		}
	} else if len(parts) == 1 {
		return num + ".0"
	}
	return num
}

func (ct *Cast) getType(Type utils.Type) string {
	switch Type {
	case utils.INT:
		return "Int"
	case utils.FLOAT:
		return "Float"
	case utils.STRING:
		return "String"
	case utils.BOOLEAN:
		return "Bool"
	case utils.CHAR:
		return "Character"
	default:
		return "nil"
	}
}
