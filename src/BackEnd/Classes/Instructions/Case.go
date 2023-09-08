package instructions

import (
	env "TSwift/Classes/Env"
	interfaces "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
	"fmt"
)

type Case struct {
	Line         int
	Column       int
	TypeInst     utils.TypeInst
	Case         interfaces.Expression
	Block        interfaces.Instruction
	CaseEvaluate *utils.ReturnType
	Flag         bool
}

func NewCase(line, column int, case_ interfaces.Expression, block interfaces.Instruction) *Case {
	return &Case{Line: line, Column: column, TypeInst: utils.CASE, Case: case_, Block: block, Flag: false}
}

func (c *Case) LineN() int {
	return c.Line
}

func (c *Case) ColumnN() int {
	return c.Column
}

func (c *Case) SetCase(caseEvaluate *utils.ReturnType) {
	c.CaseEvaluate = caseEvaluate
}

func (c *Case) Exec(Env *env.Env) *utils.ReturnType {
	envCase := env.NewEnv(Env, Env.Name+" Case")
	caseE := c.CaseEvaluate
	case_ := c.Case.Exec(envCase)
	envCase.Name += fmt.Sprintf(" %v", case_.Value)
	if c.compare(*case_, *caseE, envCase) {
		c.Flag = true
		block := c.Block.Exec(envCase)
		if block != nil {
			return block
		}
	}
	return nil
}

func (c *Case) compare(value1, value2 utils.ReturnType, Env *env.Env) bool {
	if value1.Type == value2.Type {
		if value1.Type == utils.INT {
			return value1.Value == value2.Value
		}
		if value1.Type == utils.FLOAT {
			return value1.Value == value2.Value
		}
		if value1.Type == utils.BOOLEAN {
			return fmt.Sprintf("%v", value1.Value) == fmt.Sprintf("%v", value2.Value)
		}
		if value1.Type == utils.CHAR {
			return fmt.Sprintf("%v", value1.Value) == fmt.Sprintf("%v", value2.Value)
		}
		if value1.Type == utils.STRING {
			return fmt.Sprintf("%v", value1.Value) == fmt.Sprintf("%v", value2.Value)
		}
	}
	Env.SetError("El tipo de expresion en case no es compatible con el argumento del switch", c.Line, c.Column)
	return false
}
