package instructions

import (
	env "TSwift/Classes/Env"
	interfaces "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
	vector "TSwift/Classes/Vector"
	"fmt"
)

type Remove struct {
	Line   int
	Column int
	ID     string
	Exp    interfaces.Expression
}

func NewRemove(line, column int, id string, exp interfaces.Expression) *Remove {
	return &Remove{line, column, id, exp}
}

func (r *Remove) LineN() int {
	return r.Line
}

func (r *Remove) ColumnN() int {
	return r.Column
}

func (r *Remove) Exec(env *env.Env) *utils.ReturnType {
	vec := env.GetValueID(r.ID, r.Line, r.Column)
	if vec.Type == utils.VECTOR {
		index := r.Exp.Exec(env)
		if index.Type == utils.INT {
			if len(vec.Value.(*vector.Vector).Elements) > 0 {
				if index.Value.(int) >= 0 && index.Value.(int) < len(vec.Value.(*vector.Vector).Elements) {
					vec.Value.(*vector.Vector).Elements = append(vec.Value.(*vector.Vector).Elements[:index.Value.(int)], vec.Value.(*vector.Vector).Elements[index.Value.(int)+1:]...)
					vec.Value.(*vector.Vector).Generate(env, vec.ArrType)
					return nil
				}
				env.SetError(fmt.Sprintf("Índice fuera de rango. %v:%v", r.Line, r.Column))
				return nil
			}
			env.SetError(fmt.Sprintf("Ya no hay elementos en el vector. %v:%v", r.Line, r.Column))
			return nil
		}
		env.SetError(fmt.Sprintf("El paramétro enviado no es un dato numérico entero. %v:%v", r.Line, r.Column))
		return nil
	}
	env.SetError(fmt.Sprintf("El método 'remove' es exclusivo de vectores. %v:%v", r.Line, r.Column))
	return nil
}
