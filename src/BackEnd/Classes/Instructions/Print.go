package instructions

import (
	env "TSwift/Classes/Env"
	abstracts "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
	"fmt"
)

type Print struct {
	Line     int
	Column   int
	TypeInst utils.TypeInst
	Exp      abstracts.Expression
}

func (prt *Print) NewPrint(line int, column int, exp abstracts.Expression) *Print {
	return &Print{Line: line, Column: column, Exp: exp, TypeInst: utils.PRINT}
}

func (prt *Print) Exec(env *env.Env) {
	value := ""
	if prt.Exp != nil {
		value1 := prt.Exp.Exec(env)
		value = fmt.Sprintf("%v", value1.Value)
	}
	fmt.Println(value)
}
