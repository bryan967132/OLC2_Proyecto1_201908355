package instructions

import (
	env "TSwift/Classes/Env"
	interfaces "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
	vector "TSwift/Classes/Vector"
)

type Remove struct {
	Line     int
	Column   int
	TypeInst utils.TypeInst
	ID       string
	Exp      interfaces.Expression
}

func NewRemove(line, column int, id string, exp interfaces.Expression) *Remove {
	return &Remove{line, column, utils.ARRAY_REMOVE, id, exp}
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
					vec.Value.(*vector.Vector).Values = append(vec.Value.(*vector.Vector).Values[:index.Value.(int)], vec.Value.(*vector.Vector).Values[index.Value.(int)+1:]...)
					vec.Value.(*vector.Vector).Length -= 1
					return nil
				}
				env.SetError("Índice fuera de rango", r.Line, r.Column)
				return nil
			}
			env.SetError("Ya no hay elementos en el vector", r.Line, r.Column)
			return nil
		}
		env.SetError("El paramétro enviado no es un dato numérico entero", r.Line, r.Column)
		return nil
	}
	env.SetError("El método 'remove' es exclusivo de vectores", r.Line, r.Column)
	return nil
}
