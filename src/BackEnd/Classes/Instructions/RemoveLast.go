package instructions

import (
	env "TSwift/Classes/Env"
	utils "TSwift/Classes/Utils"
	vector "TSwift/Classes/Vector"
)

type RemoveLast struct {
	Line     int
	Column   int
	TypeInst utils.TypeInst
	ID       string
}

func NewRemoveLast(line, column int, id string) *RemoveLast {
	return &RemoveLast{line, column, utils.ARRAY_REMOVELAST, id}
}

func (r *RemoveLast) LineN() int {
	return r.Line
}

func (r *RemoveLast) ColumnN() int {
	return r.Column
}

func (r *RemoveLast) Exec(env *env.Env) *utils.ReturnType {
	vec := env.GetValueID(r.ID, r.Line, r.Column)
	if vec.Type == utils.VECTOR {
		if len(vec.Value.(*vector.Vector).Elements) > 0 {
			vec.Value.(*vector.Vector).Elements = vec.Value.(*vector.Vector).Elements[:len(vec.Value.(*vector.Vector).Elements)-1]
			vec.Value.(*vector.Vector).Values = vec.Value.(*vector.Vector).Values[:len(vec.Value.(*vector.Vector).Values)-1]
			vec.Value.(*vector.Vector).Length -= 1
		}
		return nil
	}
	env.SetError("El método 'removeLast' es exclusivo de vectores", r.Line, r.Column)
	return nil
}
