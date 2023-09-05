package instructions

import (
	env "TSwift/Classes/Env"
	utils "TSwift/Classes/Utils"
	vector "TSwift/Classes/Vector"
	"fmt"
)

type RemoveLast struct {
	Line   int
	Column int
	ID     string
}

func NewRemoveLast(line, column int, id string) *RemoveLast {
	return &RemoveLast{line, column, id}
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
			vec.Value.(*vector.Vector).Generate(env, vec.ArrType)
		}
		return nil
	}
	env.SetError(fmt.Sprintf("El método 'removeLast' es exclusivo de vectores. %v:%v", r.Line, r.Column))
	return nil
}
