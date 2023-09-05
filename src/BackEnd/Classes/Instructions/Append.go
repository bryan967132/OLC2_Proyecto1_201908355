package instructions

import (
	env "TSwift/Classes/Env"
	interfaces "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
	vector "TSwift/Classes/Vector"
	"fmt"
)

type Append struct {
	Line     int
	Column   int
	TypeInst utils.TypeInst
	ID       string
	Exp      interfaces.Expression
}

func NewAppend(line, column int, id string, exp interfaces.Expression) *Append {
	return &Append{line, column, utils.ARRAY_APPEND, id, exp}
}

func (a *Append) LineN() int {
	return a.Line
}

func (a *Append) ColumnN() int {
	return a.Column
}

func (a *Append) Exec(env *env.Env) *utils.ReturnType {
	vec := env.GetValueID(a.ID, a.Line, a.Column)
	if vec.Type == utils.VECTOR {
		newValue := a.Exp.Exec(env)
		if newValue.Type == vec.ArrType {
			vec.Value.(*vector.Vector).Elements = append(vec.Value.(*vector.Vector).Elements, a.Exp)
			vec.Value.(*vector.Vector).Generate(env, newValue.Type)
			return nil
		}
		env.SetError(fmt.Sprintf("El tipo de dato del nuevo valor no es el mismo que almacena el vector. %v:%v", a.Line, a.Column))
		return nil
	}
	env.SetError(fmt.Sprintf("El método 'append' es exclusivo de vectores. %v:%v", a.Line, a.Column))
	return nil
}
