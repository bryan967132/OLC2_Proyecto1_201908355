package instructions

import (
	env "TSwift/Classes/Env"
	expressions "TSwift/Classes/Expressions"
	interfaces "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
	"fmt"
)

type Guard struct {
	Line      int
	Column    int
	TypeInst  utils.TypeInst
	Condition interfaces.Expression
	Block     interfaces.Instruction
}

func NewGuard(line, column int, condition interfaces.Expression, block interfaces.Instruction) *Guard {
	return &Guard{line, column, utils.GUARD, condition, block}
}

func (g *Guard) LineN() int {
	return g.Line
}

func (g *Guard) ColumnN() int {
	return g.Column
}

func (g *Guard) Exec(Env *env.Env) *utils.ReturnType {
	condition := g.Condition.Exec(Env)
	if condition.Type == utils.BOOLEAN {
		if g.validateInstruction(Env) {
			if fmt.Sprintf("%v", condition.Value) == "false" {
				envGuard := env.NewEnv(Env, Env.Name+" Guard")
				block := g.Block.Exec(envGuard)
				if block != nil {
					return block
				}
				return nil
			}
			return nil
		}
		Env.SetError("No se encuentra la instrucción de transferencia final", g.Line, g.Column)
		return nil
	}
	Env.SetError("No se evalúa una expresión lógica o relacional como condicion", g.Line, g.Column)
	return nil
}

func (g *Guard) validateInstruction(env *env.Env) bool {
	instructions := g.Block.(*Block).Instructions
	if len(instructions) > 0 {
		if _, ok := instructions[len(instructions)-1].(*expressions.Return); ok {
			return true
		}
		if _, ok := instructions[len(instructions)-1].(*Continue); ok {
			return true
		}
		if _, ok := instructions[len(instructions)-1].(*Break); ok {
			return true
		}
	}
	return false
}
