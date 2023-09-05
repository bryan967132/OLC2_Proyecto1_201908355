package instructions

import (
	env "TSwift/Classes/Env"
	interfaces "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
	vector "TSwift/Classes/Vector"
	"fmt"
)

type AsignPosArray struct {
	Line     int
	Column   int
	TypeInst utils.TypeInst
	Id       string
	Index    []interfaces.Expression
	NewValue interfaces.Expression
}

func NewAsignPosArray(line, column int, id string, index []interfaces.Expression, newValue interfaces.Expression) *AsignPosArray {
	return &AsignPosArray{line, column, utils.ASIGN_ARRAY, id, index, newValue}
}

func (ac *AsignPosArray) LineN() int {
	return ac.Line
}

func (ac *AsignPosArray) ColumnN() int {
	return ac.Column
}

func (ac *AsignPosArray) Exec(env *env.Env) *utils.ReturnType {
	value := env.GetValueID(ac.Id, ac.Line, ac.Column)
	if value != nil {
		if value.Type == utils.VECTOR {
			vec := value.Value.(*vector.Vector)
			if vec.Dims == len(ac.Index) {
				indexs := []int{}
				var in *utils.ReturnType
				for _, exp := range ac.Index {
					in = exp.Exec(env)
					if in.Type != utils.INT {
						env.SetError(fmt.Sprintf("Él índice no toma una expresión de tipo entero. %v:%v", ac.Line, ac.Column))
						return &utils.ReturnType{Value: "nil", Type: utils.NIL}
					}
					indexs = append(indexs, in.Value.(int))
				}
				vec.SetValuePosition(env, indexs, ac.NewValue, ac.Line, ac.Column)
				return nil
			}
			env.SetError(fmt.Sprintf("Las dimensiones no coinciden con las del vector. %v:%v", ac.Line, ac.Column))
			return nil
		}
		env.SetError(fmt.Sprintf("La variable llamada no es un vector. %v:%v", ac.Line, ac.Column))
	}
	return nil
}
