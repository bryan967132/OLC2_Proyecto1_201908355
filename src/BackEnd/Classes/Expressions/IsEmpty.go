package expressions

import (
	env "TSwift/Classes/Env"
	interfaces "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
	vector "TSwift/Classes/Vector"
	"fmt"
)

type IsEmpty struct {
	Line    int
	Column  int
	TypeExp utils.TypeExp
	Exp     interfaces.Expression
}

func NewIsEmpty(line, column int, exp interfaces.Expression) *IsEmpty {
	return &IsEmpty{line, column, utils.ARRAY_ISEMPTY, exp}
}

func (i *IsEmpty) LineN() int {
	return i.Line
}

func (i *IsEmpty) ColumnN() int {
	return i.Column
}

func (i *IsEmpty) Exec(env *env.Env) *utils.ReturnType {
	vec := i.Exp.Exec(env)
	if vec.Type == utils.VECTOR {
		return &utils.ReturnType{Value: vec.Value.(*vector.Vector).Length == 0, Type: utils.BOOLEAN}
	}
	env.SetError(fmt.Sprintf("El atributo 'isEmpty' es exclusivo de vectores. %v:%v", i.Line, i.Column))
	return &utils.ReturnType{Value: "nil", Type: utils.NIL}
}
