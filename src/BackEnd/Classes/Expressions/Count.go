package expressions

import (
	env "TSwift/Classes/Env"
	interfaces "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
	vector "TSwift/Classes/Vector"
	"fmt"
)

type Count struct {
	Line   int
	Column int
	Exp    interfaces.Expression
}

func NewCount(line, column int, exp interfaces.Expression) *Count {
	return &Count{line, column, exp}
}

func (c *Count) LineN() int {
	return c.Line
}

func (c *Count) ColumnN() int {
	return c.Column
}

func (c *Count) Exec(env *env.Env) *utils.ReturnType {
	vec := c.Exp.Exec(env)
	if vec.Type == utils.VECTOR {
		return &utils.ReturnType{Value: vec.Value.(*vector.Vector).Length, Type: utils.INT}
	}
	env.SetError(fmt.Sprintf("El atributo 'count' es exclusivo de vectores. %v:%v", c.Line, c.Column))
	return &utils.ReturnType{Value: "nil", Type: utils.NIL}
}
