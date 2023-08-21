package instructions

import (
	env "TSwift/Classes/Env"
	interfaces "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
	"fmt"
)

type While struct {
	Line      int
	Column    int
	TypeInst  utils.TypeInst
	Condition interfaces.Expression
	Block     interfaces.Instruction
}

func NewWhile(line, column int, condition interfaces.Expression, block interfaces.Instruction) *While {
	return &While{line, column, utils.LOOP_WHILE, condition, block}
}

func (w *While) LineN() int {
	return w.Line
}

func (w *While) ColumnN() int {
	return w.Column
}

func (w *While) Exec(Env *env.Env) interface{} {
	envWhile := env.NewEnv(Env, Env.Name+" While")
	condition := w.Condition.Exec(Env)
	for fmt.Sprintf("%v", condition.Value) == "true" {
		block := w.Block.Exec(envWhile)
		if block != nil {
			if block.(utils.ReturnType).Value == utils.CONTINUE {
				condition = w.Condition.Exec(Env)
				continue
			}
			if block.(utils.ReturnType).Value == utils.BREAK {
				break
			}
			return block
		}
		condition = w.Condition.Exec(Env)
	}
	return nil
}
