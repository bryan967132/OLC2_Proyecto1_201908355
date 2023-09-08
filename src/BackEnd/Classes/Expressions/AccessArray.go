package expressions

import (
	env "TSwift/Classes/Env"
	interfaces "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
	vector "TSwift/Classes/Vector"
)

type AccessArray struct {
	Line    int
	Column  int
	TypeExp utils.TypeExp
	Id      string
	Index   []interfaces.Expression
}

func NewAccessArray(line, column int, id string, index []interfaces.Expression) *AccessArray {
	return &AccessArray{line, column, utils.ACCESS_ARRAY, id, index}
}

func (ac *AccessArray) LineN() int {
	return ac.Line
}

func (ac *AccessArray) ColumnN() int {
	return ac.Column
}

func (ac *AccessArray) Exec(env *env.Env) *utils.ReturnType {
	value := env.GetValueID(ac.Id, ac.Line, ac.Column)
	if value != nil {
		if value.Type == utils.VECTOR {
			vec := value.Value.(*vector.Vector)
			if vec.Dims >= len(ac.Index) {
				indexs := []int{}
				var in *utils.ReturnType
				for _, exp := range ac.Index {
					in = exp.Exec(env)
					if in.Type != utils.INT {
						env.SetError("Él índice no toma una expresión de tipo entero", ac.Line, ac.Column)
						return &utils.ReturnType{Value: "nil", Type: utils.NIL}
					}
					indexs = append(indexs, in.Value.(int))
				}
				value := vec.GetPosition(env, indexs, ac.Line, ac.Column)
				if value == nil {
					return &utils.ReturnType{Value: "nil", Type: utils.NIL}
				}
				switch value.(type) {
				case *vector.Vector:
					return &utils.ReturnType{Value: value.(*vector.Vector), Type: utils.VECTOR}
				case *utils.ReturnType:
					return value.(*utils.ReturnType)
				}
			}
			env.SetError("Las dimensiones no coinciden con las del vector", ac.Line, ac.Column)
			return &utils.ReturnType{Value: "nil", Type: utils.NIL}
		}
		env.SetError("La variable llamada no es un vector", ac.Line, ac.Column)
	}
	return &utils.ReturnType{Value: "nil", Type: utils.NIL}
}
