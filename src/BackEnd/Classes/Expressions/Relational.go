package expressions

import (
	env "TSwift/Classes/Env"
	interfaces "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
	"fmt"
	"strconv"
)

type Relational struct {
	Line    int
	Column  int
	TypeExp utils.TypeExp
	Exp1    interfaces.Expression
	Sign    string
	Exp2    interfaces.Expression
}

func NewRelational(line, column int, exp1 interfaces.Expression, sign string, exp2 interfaces.Expression) *Relational {
	return &Relational{line, column, utils.RELATIONAL_OP, exp1, sign, exp2}
}

func (rl *Relational) LineN() int {
	return rl.Line
}

func (rl *Relational) ColumnN() int {
	return rl.Column
}

func (rl *Relational) Exec(env *env.Env) *utils.ReturnType {
	switch rl.Sign {
	case "==":
		return rl.equal(env)
	case "!=":
		return rl.notEqual(env)
	case ">=":
		return rl.moreEqual(env)
	case "<=":
		return rl.lessEqual(env)
	case ">":
		return rl.more(env)
	case "<":
		return rl.less(env)
	default:
		return &utils.ReturnType{Value: "nil", Type: utils.NIL}
	}
}

func (rl *Relational) equal(env *env.Env) *utils.ReturnType {
	value1 := rl.Exp1.Exec(env)
	value2 := rl.Exp2.Exec(env)
	if value1.Type == utils.INT && value2.Type == utils.INT {
		return &utils.ReturnType{Value: value1.Value.(int) == value2.Value.(int), Type: utils.BOOLEAN}
	}
	if value1.Type == utils.FLOAT && value2.Type == utils.FLOAT {
		floatValue1, _ := strconv.ParseFloat(fmt.Sprintf("%v", value1.Value), 64)
		floatValue2, _ := strconv.ParseFloat(fmt.Sprintf("%v", value2.Value), 64)
		return &utils.ReturnType{Value: floatValue1 == floatValue2, Type: utils.BOOLEAN}
	}
	if value1.Type == utils.BOOLEAN && value2.Type == utils.BOOLEAN {
		return &utils.ReturnType{Value: value1.Value.(bool) == value2.Value.(bool), Type: utils.BOOLEAN}
	}
	if value1.Type == utils.STRING && value2.Type == utils.STRING {
		return &utils.ReturnType{Value: fmt.Sprintf("%v", value1.Value) == fmt.Sprintf("%v", value2.Value), Type: utils.BOOLEAN}
	}
	if value1.Type == utils.CHAR && value2.Type == utils.CHAR {
		return &utils.ReturnType{Value: fmt.Sprintf("%v", value1.Value) == fmt.Sprintf("%v", value2.Value), Type: utils.BOOLEAN}
	}
	env.SetError(fmt.Sprintf("%s %v:%v", "Los tipos no coinciden para operaciones relacionales.", rl.Exp1.LineN(), rl.Exp1.ColumnN()))
	return &utils.ReturnType{Value: "nil", Type: utils.NIL}
}

func (rl *Relational) notEqual(env *env.Env) *utils.ReturnType {
	value1 := rl.Exp1.Exec(env)
	value2 := rl.Exp2.Exec(env)
	if value1.Type == utils.INT && value2.Type == utils.INT {
		return &utils.ReturnType{Value: value1.Value.(int) != value2.Value.(int), Type: utils.BOOLEAN}
	}
	if value1.Type == utils.FLOAT && value2.Type == utils.FLOAT {
		floatValue1, _ := strconv.ParseFloat(fmt.Sprintf("%v", value1.Value), 64)
		floatValue2, _ := strconv.ParseFloat(fmt.Sprintf("%v", value2.Value), 64)
		return &utils.ReturnType{Value: floatValue1 != floatValue2, Type: utils.BOOLEAN}
	}
	if value1.Type == utils.BOOLEAN && value2.Type == utils.BOOLEAN {
		return &utils.ReturnType{Value: value1.Value.(bool) != value2.Value.(bool), Type: utils.BOOLEAN}
	}
	if value1.Type == utils.STRING && value2.Type == utils.STRING {
		return &utils.ReturnType{Value: fmt.Sprintf("%v", value1.Value) != fmt.Sprintf("%v", value2.Value), Type: utils.BOOLEAN}
	}
	if value1.Type == utils.CHAR && value2.Type == utils.CHAR {
		return &utils.ReturnType{Value: fmt.Sprintf("%v", value1.Value) != fmt.Sprintf("%v", value2.Value), Type: utils.BOOLEAN}
	}
	env.SetError(fmt.Sprintf("%s %v:%v", "Los tipos no coinciden para operaciones relacionales.", rl.Exp1.LineN(), rl.Exp1.ColumnN()))
	return &utils.ReturnType{Value: "nil", Type: utils.NIL}
}

func (rl *Relational) moreEqual(env *env.Env) *utils.ReturnType {
	value1 := rl.Exp1.Exec(env)
	value2 := rl.Exp2.Exec(env)
	if value1.Type == utils.INT && value2.Type == utils.INT {
		return &utils.ReturnType{Value: value1.Value.(int) >= value2.Value.(int), Type: utils.BOOLEAN}
	}
	if value1.Type == utils.FLOAT && value2.Type == utils.FLOAT {
		floatValue1, _ := strconv.ParseFloat(fmt.Sprintf("%v", value1.Value), 64)
		floatValue2, _ := strconv.ParseFloat(fmt.Sprintf("%v", value2.Value), 64)
		return &utils.ReturnType{Value: floatValue1 >= floatValue2, Type: utils.BOOLEAN}
	}
	if value1.Type == utils.STRING && value2.Type == utils.STRING {
		return &utils.ReturnType{Value: fmt.Sprintf("%v", value1.Value) >= fmt.Sprintf("%v", value2.Value), Type: utils.BOOLEAN}
	}
	if value1.Type == utils.CHAR && value2.Type == utils.CHAR {
		return &utils.ReturnType{Value: fmt.Sprintf("%v", value1.Value) >= fmt.Sprintf("%v", value2.Value), Type: utils.BOOLEAN}
	}
	env.SetError(fmt.Sprintf("%s %v:%v", "Los tipos no coinciden para operaciones relacionales.", rl.Exp1.LineN(), rl.Exp1.ColumnN()))
	return &utils.ReturnType{Value: "nil", Type: utils.NIL}
}

func (rl *Relational) lessEqual(env *env.Env) *utils.ReturnType {
	value1 := rl.Exp1.Exec(env)
	value2 := rl.Exp2.Exec(env)
	if value1.Type == utils.INT && value2.Type == utils.INT {
		return &utils.ReturnType{Value: value1.Value.(int) <= value2.Value.(int), Type: utils.BOOLEAN}
	}
	if value1.Type == utils.FLOAT && value2.Type == utils.FLOAT {
		floatValue1, _ := strconv.ParseFloat(fmt.Sprintf("%v", value1.Value), 64)
		floatValue2, _ := strconv.ParseFloat(fmt.Sprintf("%v", value2.Value), 64)
		return &utils.ReturnType{Value: floatValue1 <= floatValue2, Type: utils.BOOLEAN}
	}
	if value1.Type == utils.STRING && value2.Type == utils.STRING {
		return &utils.ReturnType{Value: fmt.Sprintf("%v", value1.Value) <= fmt.Sprintf("%v", value2.Value), Type: utils.BOOLEAN}
	}
	if value1.Type == utils.CHAR && value2.Type == utils.CHAR {
		return &utils.ReturnType{Value: fmt.Sprintf("%v", value1.Value) <= fmt.Sprintf("%v", value2.Value), Type: utils.BOOLEAN}
	}
	env.SetError(fmt.Sprintf("%s %v:%v", "Los tipos no coinciden para operaciones relacionales.", rl.Exp1.LineN(), rl.Exp1.ColumnN()))
	return &utils.ReturnType{Value: "nil", Type: utils.NIL}
}

func (rl *Relational) more(env *env.Env) *utils.ReturnType {
	value1 := rl.Exp1.Exec(env)
	value2 := rl.Exp2.Exec(env)
	if value1.Type == utils.INT && value2.Type == utils.INT {
		return &utils.ReturnType{Value: value1.Value.(int) > value2.Value.(int), Type: utils.BOOLEAN}
	}
	if value1.Type == utils.FLOAT && value2.Type == utils.FLOAT {
		floatValue1, _ := strconv.ParseFloat(fmt.Sprintf("%v", value1.Value), 64)
		floatValue2, _ := strconv.ParseFloat(fmt.Sprintf("%v", value2.Value), 64)
		return &utils.ReturnType{Value: floatValue1 > floatValue2, Type: utils.BOOLEAN}
	}
	if value1.Type == utils.STRING && value2.Type == utils.STRING {
		return &utils.ReturnType{Value: fmt.Sprintf("%v", value1.Value) > fmt.Sprintf("%v", value2.Value), Type: utils.BOOLEAN}
	}
	if value1.Type == utils.CHAR && value2.Type == utils.CHAR {
		return &utils.ReturnType{Value: fmt.Sprintf("%v", value1.Value) > fmt.Sprintf("%v", value2.Value), Type: utils.BOOLEAN}
	}
	env.SetError(fmt.Sprintf("%s %v:%v", "Los tipos no coinciden para operaciones relacionales.", rl.Exp1.LineN(), rl.Exp1.ColumnN()))
	return &utils.ReturnType{Value: "nil", Type: utils.NIL}
}

func (rl *Relational) less(env *env.Env) *utils.ReturnType {
	value1 := rl.Exp1.Exec(env)
	value2 := rl.Exp2.Exec(env)
	if value1.Type == utils.INT && value2.Type == utils.INT {
		return &utils.ReturnType{Value: value1.Value.(int) < value2.Value.(int), Type: utils.BOOLEAN}
	}
	if value1.Type == utils.FLOAT && value2.Type == utils.FLOAT {
		floatValue1, _ := strconv.ParseFloat(fmt.Sprintf("%v", value1.Value), 64)
		floatValue2, _ := strconv.ParseFloat(fmt.Sprintf("%v", value2.Value), 64)
		return &utils.ReturnType{Value: floatValue1 < floatValue2, Type: utils.BOOLEAN}
	}
	if value1.Type == utils.STRING && value2.Type == utils.STRING {
		return &utils.ReturnType{Value: fmt.Sprintf("%v", value1.Value) < fmt.Sprintf("%v", value2.Value), Type: utils.BOOLEAN}
	}
	if value1.Type == utils.CHAR && value2.Type == utils.CHAR {
		return &utils.ReturnType{Value: fmt.Sprintf("%v", value1.Value) < fmt.Sprintf("%v", value2.Value), Type: utils.BOOLEAN}
	}
	env.SetError(fmt.Sprintf("%s %v:%v", "Los tipos no coinciden para operaciones relacionales.", rl.Exp1.LineN(), rl.Exp1.ColumnN()))
	return &utils.ReturnType{Value: "nil", Type: utils.NIL}
}
