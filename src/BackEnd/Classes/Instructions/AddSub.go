package instructions

import (
	env "TSwift/Classes/Env"
	expressions "TSwift/Classes/Expressions"
	interfaces "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
	"fmt"
)

type AddSub struct {
	Line     int
	Column   int
	TypeInst utils.TypeInst
	Id       string
	Sign     string
	Exp      interfaces.Expression
}

func NewAddSub(line, column int, id, sign string, exp interfaces.Expression) *AddSub {
	if id == "+=" {
		return &AddSub{line, column, utils.ADD, id, sign, exp}
	}
	return &AddSub{line, column, utils.SUB, id, sign, exp}
}

func (a *AddSub) LineN() int {
	return a.Line
}

func (a *AddSub) ColumnN() int {
	return a.Column
}

func (a *AddSub) Exec(env *env.Env) *utils.ReturnType {
	value := env.GetValueID(a.Id, a.Line, a.Column)
	switch a.Sign {
	case "+=":
		env.ReasignID(
			a.Id,
			expressions.NewArithmetic(a.Line, a.Column,
				expressions.NewPrimitive(a.Line, a.Column, fmt.Sprintf("%v", value.Value.(*utils.ReturnType).Value), value.Type),
				"+",
				a.Exp).Exec(env),
			a.Line,
			a.Column)
	case "-=":
		env.ReasignID(
			a.Id,
			expressions.NewArithmetic(a.Line, a.Column,
				expressions.NewPrimitive(a.Line, a.Column, fmt.Sprintf("%v", value.Value.(*utils.ReturnType).Value), value.Type),
				"-",
				a.Exp).Exec(env),
			a.Line,
			a.Column)
	}
	return nil
}
