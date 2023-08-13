package instructions

import (
	env "TSwift/Classes/Env"
	abstracts "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
	"fmt"
	"sort"
)

type Print struct {
	Line     int
	Column   int
	TypeInst utils.TypeInst
	Exps     []abstracts.Expression
}

func NewPrint(line int, column int, exps []abstracts.Expression) *Print {
	return &Print{line, column, utils.PRINT, exps}
}

func (prt *Print) Exec(env *env.Env) interface{} {
	value := ""
	if prt.Exps != nil {
		sort.Slice(prt.Exps, func(i, j int) bool { return true })
		for i := 0; i < len(prt.Exps); i++ {
			value1 := prt.Exps[i].Exec(env)
			if value != "" {
				value = fmt.Sprintf("%s %v", value, value1.Value)
			} else {
				value = fmt.Sprintf("%v", value1.Value)
			}
		}
	}
	fmt.Println(value)
	return nil
}
