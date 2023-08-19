package instructions

import (
	env "TSwift/Classes/Env"
	interfaces "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
	"fmt"
)

type If struct {
	Line      int
	Column    int
	TypeInst  utils.TypeInst
	Condition interfaces.Expression
	Block     interfaces.Instruction
	Except    interfaces.Instruction
}

func NewIf(line, column int, condition interfaces.Expression, block interfaces.Instruction, except interfaces.Instruction) *If {
	return &If{Line: line, Column: column, TypeInst: utils.IF, Condition: condition, Block: block, Except: except}
}

func (i *If) LineN() int {
	return i.Line
}

func (i *If) ColumnN() int {
	return i.Column
}

func (i *If) Exec(Env *env.Env) interface{} {
	condition := i.Condition.Exec(Env)
	if condition.Type != utils.BOOLEAN {
		Env.SetError(fmt.Sprintf("No se evalúa una expresión lógica o relacional como condicion. %v:%v", i.Line, i.Column))
		return nil
	}
	if fmt.Sprintf("%v", condition.Value) == "true" {
		var envIf *env.Env = env.NewEnv(Env, Env.Name+" If")
		block := i.Block.Exec(envIf)
		if block != nil {
			return block
		}
		return nil
	}
	// else
	if i.Except != nil {
		except := i.Except.Exec(Env)
		if except != nil {
			return except
		}
	}
	return nil
}
