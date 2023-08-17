package expressions

import (
	env "TSwift/Classes/Env"
	interfaces "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
	"fmt"
	"strconv"
)

type Arithmetic struct {
	Line    int
	Column  int
	TypeExp utils.TypeExp
	Exp1    interfaces.Expression
	Sign    string
	Exp2    interfaces.Expression
	Type    utils.Type
}

func NewArithmetic(line, column int, exp1 interfaces.Expression, sign string, exp2 interfaces.Expression) *Arithmetic {
	return &Arithmetic{line, column, utils.ARITHMETIC_OP, exp1, sign, exp2, utils.NIL}
}

func (ar *Arithmetic) LineN() int {
	return ar.Line
}

func (ar *Arithmetic) ColumnN() int {
	return ar.Column
}

func (ar *Arithmetic) Exec(env *env.Env) *utils.ReturnType {
	switch ar.Sign {
	case "+":
		return ar.plus(env)
	case "-":
		if ar.Exp1 != nil {
			return ar.minus(env)
		}
		return ar.uminus(env)
	case "*":
		return ar.mult(env)
	case "/":
		return ar.div(env)
	case "%":
		return ar.mod(env)
	default:
		return &utils.ReturnType{Value: "nil", Type: utils.NIL}
	}
}

func (ar *Arithmetic) plus(env *env.Env) *utils.ReturnType {
	value1 := ar.Exp1.Exec(env)
	value2 := ar.Exp2.Exec(env)
	if int(value1.Type) < len(utils.Plus) && int(value2.Type) < len(utils.Plus[0]) {
		ar.Type = utils.Plus[value1.Type][value2.Type]
	} else {
		ar.Type = utils.NIL
	}
	if ar.Type != utils.NIL {
		if ar.Type == utils.INT {
			return &utils.ReturnType{Value: value1.Value.(int) + value2.Value.(int), Type: ar.Type}
		}
		if ar.Type == utils.FLOAT {
			floatValue1, _ := strconv.ParseFloat(fmt.Sprintf("%v", value1.Value), 64)
			floatValue2, _ := strconv.ParseFloat(fmt.Sprintf("%v", value2.Value), 64)
			return &utils.ReturnType{Value: floatValue1 + floatValue2, Type: ar.Type}
		}
		if ar.Type == utils.STRING {
			return &utils.ReturnType{Value: fmt.Sprintf("%v%v", value1.Value, value2.Value), Type: ar.Type}
		}
	}
	env.SetError(fmt.Sprintf("%s %v:%v", "Los tipos no coinciden para operaciones aritméticas.", ar.Exp1.LineN(), ar.Exp1.ColumnN()))
	return &utils.ReturnType{Value: "nil", Type: ar.Type}
}

func (ar *Arithmetic) minus(env *env.Env) *utils.ReturnType {
	value1 := ar.Exp1.Exec(env)
	value2 := ar.Exp2.Exec(env)
	if int(value1.Type) < len(utils.Minus) && int(value2.Type) < len(utils.Minus[0]) {
		ar.Type = utils.Minus[value1.Type][value2.Type]
	} else {
		ar.Type = utils.NIL
	}
	if ar.Type != utils.NIL {
		if ar.Type == utils.INT {
			return &utils.ReturnType{Value: value1.Value.(int) - value2.Value.(int), Type: ar.Type}
		}
		if ar.Type == utils.FLOAT {
			floatValue1, _ := strconv.ParseFloat(fmt.Sprintf("%v", value1.Value), 64)
			floatValue2, _ := strconv.ParseFloat(fmt.Sprintf("%v", value2.Value), 64)
			return &utils.ReturnType{Value: floatValue1 - floatValue2, Type: ar.Type}
		}
	}
	env.SetError(fmt.Sprintf("%s %v:%v", "Los tipos no coinciden para operaciones aritméticas.", ar.Exp1.LineN(), ar.Exp1.ColumnN()))
	return &utils.ReturnType{Value: "nil", Type: ar.Type}
}

func (ar *Arithmetic) uminus(env *env.Env) *utils.ReturnType {
	value2 := ar.Exp2.Exec(env)
	ar.Type = value2.Type
	if ar.Type == utils.INT {
		return &utils.ReturnType{Value: -value2.Value.(int), Type: ar.Type}
	}
	if ar.Type == utils.FLOAT {
		floatValue2, _ := strconv.ParseFloat(fmt.Sprintf("%v", value2.Value), 64)
		return &utils.ReturnType{Value: -floatValue2, Type: ar.Type}
	}
	env.SetError(fmt.Sprintf("%s %v:%v", "Los tipos no coinciden para operaciones aritméticas.", ar.Exp2.LineN(), ar.Exp2.ColumnN()))
	return &utils.ReturnType{Value: "nil", Type: ar.Type}
}

func (ar *Arithmetic) mult(env *env.Env) *utils.ReturnType {
	value1 := ar.Exp1.Exec(env)
	value2 := ar.Exp2.Exec(env)
	if int(value1.Type) < len(utils.Mult) && int(value2.Type) < len(utils.Mult[0]) {
		ar.Type = utils.Mult[value1.Type][value2.Type]
	} else {
		ar.Type = utils.NIL
	}
	if ar.Type != utils.NIL {
		if ar.Type == utils.INT {
			return &utils.ReturnType{Value: value1.Value.(int) * value2.Value.(int), Type: ar.Type}
		}
		if ar.Type == utils.FLOAT {
			floatValue1, _ := strconv.ParseFloat(fmt.Sprintf("%v", value1.Value), 64)
			floatValue2, _ := strconv.ParseFloat(fmt.Sprintf("%v", value2.Value), 64)
			return &utils.ReturnType{Value: floatValue1 * floatValue2, Type: ar.Type}
		}
	}
	env.SetError(fmt.Sprintf("%s %v:%v", "Los tipos no coinciden para operaciones aritméticas.", ar.Exp1.LineN(), ar.Exp1.ColumnN()))
	return &utils.ReturnType{Value: "nil", Type: ar.Type}
}

func (ar *Arithmetic) div(env *env.Env) *utils.ReturnType {
	value1 := ar.Exp1.Exec(env)
	value2 := ar.Exp2.Exec(env)
	if int(value1.Type) < len(utils.Div) && int(value2.Type) < len(utils.Div[0]) {
		ar.Type = utils.Div[value1.Type][value2.Type]
	} else {
		ar.Type = utils.NIL
	}
	if ar.Type != utils.NIL {
		if ar.Type == utils.INT {
			if value2.Value.(int) != 0 {
				return &utils.ReturnType{Value: value1.Value.(int) / value2.Value.(int), Type: ar.Type}
			}
			env.SetError(fmt.Sprintf("%s %v:%v", "División entre cero.", ar.Exp1.LineN(), ar.Exp1.ColumnN()))
			return &utils.ReturnType{Value: "nil", Type: ar.Type}
		}
		if ar.Type == utils.FLOAT {
			floatValue1, _ := strconv.ParseFloat(fmt.Sprintf("%v", value1.Value), 64)
			floatValue2, _ := strconv.ParseFloat(fmt.Sprintf("%v", value2.Value), 64)
			if floatValue2 != 0 {
				return &utils.ReturnType{Value: floatValue1 / floatValue2, Type: ar.Type}
			}
			env.SetError(fmt.Sprintf("%s %v:%v", "División entre cero.", ar.Exp2.LineN(), ar.Exp2.ColumnN()))
			return &utils.ReturnType{Value: "nil", Type: ar.Type}
		}
	}
	env.SetError(fmt.Sprintf("%s %v:%v", "Los tipos no coinciden para operaciones aritméticas.", ar.Exp1.LineN(), ar.Exp1.ColumnN()))
	return &utils.ReturnType{Value: "nil", Type: ar.Type}
}

func (ar *Arithmetic) mod(env *env.Env) *utils.ReturnType {
	value1 := ar.Exp1.Exec(env)
	value2 := ar.Exp2.Exec(env)
	if int(value1.Type) < len(utils.Mod) && int(value2.Type) < len(utils.Mod[0]) {
		ar.Type = utils.Mod[value1.Type][value2.Type]
	} else {
		ar.Type = utils.NIL
	}
	if ar.Type != utils.NIL {
		if ar.Type == utils.INT {
			if value2.Value.(int) != 0 {
				return &utils.ReturnType{Value: value1.Value.(int) / value2.Value.(int), Type: ar.Type}
			}
			env.SetError(fmt.Sprintf("%s %v:%v", "División entre cero.", ar.Exp1.LineN(), ar.Exp1.ColumnN()))
			return &utils.ReturnType{Value: "nil", Type: ar.Type}
		}
	}
	env.SetError(fmt.Sprintf("%s %v:%v", "Los tipos no coinciden para operaciones aritméticas.", ar.Exp1.LineN(), ar.Exp1.ColumnN()))
	return &utils.ReturnType{Value: "nil", Type: ar.Type}
}
