package expressions

import (
	env "TSwift/Classes/Env"
	interfaces "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
)

type Logic struct {
	Line    int
	Column  int
	TypeExp utils.TypeExp
	Exp1    interfaces.Expression
	Sign    string
	Exp2    interfaces.Expression
}

func NewLogic(line, column int, exp1 interfaces.Expression, sign string, exp2 interfaces.Expression) *Logic {
	return &Logic{line, column, utils.LOGIC_OP, exp1, sign, exp2}
}

func (lg *Logic) LineN() int {
	return lg.Line
}

func (lg *Logic) ColumnN() int {
	return lg.Column
}

func (lg *Logic) Exec(env *env.Env) *utils.ReturnType {
	switch lg.Sign {
	case "&&":
		return lg.and(env)
	case "||":
		return lg.or(env)
	case "!":
		return lg.not(env)
	default:
		return &utils.ReturnType{Value: "nil", Type: utils.NIL}
	}
}

func (lg *Logic) and(env *env.Env) *utils.ReturnType {
	value1 := lg.Exp1.Exec(env)
	if value1.Type != utils.BOOLEAN {
		env.SetError("Los tipos no son válidos para operaciones lógicas", lg.Exp1.LineN(), lg.Exp1.ColumnN())
		return &utils.ReturnType{Value: "nil", Type: utils.NIL}
	}
	value2 := lg.Exp2.Exec(env)
	if value2.Type == utils.BOOLEAN {
		return &utils.ReturnType{Value: value1.Value.(bool) && value2.Value.(bool), Type: utils.BOOLEAN}
	}
	env.SetError("Los tipos no son válidos para operaciones lógicas", lg.Exp2.LineN(), lg.Exp2.ColumnN())
	return &utils.ReturnType{Value: "nil", Type: utils.NIL}
}

func (lg *Logic) or(env *env.Env) *utils.ReturnType {
	value1 := lg.Exp1.Exec(env)
	if value1.Type != utils.BOOLEAN {
		env.SetError("Los tipos no son válidos para operaciones lógicas", lg.Exp1.LineN(), lg.Exp1.ColumnN())
		return &utils.ReturnType{Value: "nil", Type: utils.NIL}
	}
	value2 := lg.Exp2.Exec(env)
	if value2.Type == utils.BOOLEAN {
		return &utils.ReturnType{Value: value1.Value.(bool) || value2.Value.(bool), Type: utils.BOOLEAN}
	}
	env.SetError("Los tipos no son válidos para operaciones lógicas", lg.Exp2.LineN(), lg.Exp2.ColumnN())
	return &utils.ReturnType{Value: "nil", Type: utils.NIL}
}

func (lg *Logic) not(env *env.Env) *utils.ReturnType {
	value2 := lg.Exp2.Exec(env)
	if value2.Type == utils.BOOLEAN {
		return &utils.ReturnType{Value: !value2.Value.(bool), Type: utils.BOOLEAN}
	}
	env.SetError("Los tipos no son válidos para operaciones lógicas", lg.Exp2.LineN(), lg.Exp2.ColumnN())
	return &utils.ReturnType{Value: "nil", Type: utils.NIL}
}
