package instructions

import (
	env "TSwift/Classes/Env"
	interfaces "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
	"fmt"
)

type Print struct {
	Line     int
	Column   int
	TypeInst utils.TypeInst
	Exps     []interfaces.Expression
}

func NewPrint(line int, column int, exps []interfaces.Expression) *Print {
	return &Print{line, column, utils.PRINT, exps}
}

func (prt *Print) LineN() int {
	return prt.Line
}

func (prt *Print) ColumnN() int {
	return prt.Column
}

func (prt *Print) Exec(env *env.Env) *utils.ReturnType {
	value := ""
	if prt.Exps != nil {
		for _, exp := range prt.Exps {
			value1 := exp.Exec(env)
			if value != "" {
				value = fmt.Sprintf("%s %v", value, value1.Value)
			} else {
				value = fmt.Sprintf("%v", value1.Value)
			}
		}
	}
	env.SetPrints(value)
	return nil
}
