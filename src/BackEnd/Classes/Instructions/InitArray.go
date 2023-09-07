package instructions

import (
	env "TSwift/Classes/Env"
	utils "TSwift/Classes/Utils"
	vector "TSwift/Classes/Vector"
	"fmt"
)

type InitVector struct {
	Line       int
	Column     int
	IsVariable bool
	TypeInst   utils.TypeInst
	Id         string
	Type       *utils.AttribsType
	Value      *vector.Vector
}

func NewInitVector(line, column int, isVar bool, id string, Type *utils.AttribsType, value *vector.Vector) *InitVector {
	return &InitVector{line, column, isVar, utils.INIT_ID, id, Type, value}
}

func (in *InitVector) LineN() int {
	return in.Line
}

func (in *InitVector) ColumnN() int {
	return in.Column
}

func (in *InitVector) Exec(env *env.Env) *utils.ReturnType {
	if in.Type != nil {
		if in.Type.IsPrimitive {
			if in.Value.Type != nil {
				if !in.Value.Type.IsPrimitive {
					env.SetError(fmt.Sprintf("Los tipos no coinciden para el vector. %v:%v", in.Line, in.Column))
					return nil
				}
				if in.Type.Value.(utils.Type) != in.Value.Type.Value.(utils.Type) {
					env.SetError(fmt.Sprintf("Los tipos no coinciden para el vector. %v:%v", in.Line, in.Column))
					return nil
				}
			}
			var values *vector.Vector
			if !in.Value.IsReuseID {
				isGenerated := in.Value.Generate(env, in.Type.Value.(utils.Type))
				if !isGenerated {
					env.SetError(fmt.Sprintf("Los tipos no coinciden para el vector. %v:%v", in.Line, in.Column))
					return nil
				}
				values = in.Value
			} else {
				valuesID := env.GetValueID(in.Value.ID, in.Line, in.Column)
				if valuesID == nil {
					return nil
				}
				if valuesID.Type != utils.VECTOR {
					env.SetError(fmt.Sprintf("El ID no almacena un vector. %v:%v", in.Line, in.Column))
					return nil
				}
				reuseVector := *valuesID.Value.(*vector.Vector)
				values = (&reuseVector).GetCopy()
			}
			env.SaveArray(in.IsVariable, in.Id, values, values.Type.Value.(utils.Type), in.Line, in.Column)
			return nil
		} else {
			// vector con tipo Struct
			return nil
		}
	}
	if in.Value.Type != nil {
		if in.Value.Type.IsPrimitive {
			var values *vector.Vector
			if !in.Value.IsReuseID {
				isGenerated := in.Value.Generate(env, in.Value.Type.Value.(utils.Type))
				if !isGenerated {
					env.SetError(fmt.Sprintf("Los tipos no coinciden para el vector. %v:%v", in.Line, in.Column))
					return nil
				}
				values = in.Value
			} else {
				valuesID := env.GetValueID(in.Value.ID, in.Line, in.Column)
				if valuesID.Type != utils.VECTOR {
					env.SetError(fmt.Sprintf("El ID no almacena un vector. %v:%v", in.Line, in.Column))
					return nil
				}
				reuseVector := *valuesID.Value.(*vector.Vector)
				values = &reuseVector
			}
			env.SaveArray(in.IsVariable, in.Id, values, values.Type.Value.(utils.Type), in.Line, in.Column)
			return nil
		} else {
			// vector con tipo Struct
		}
	} else {
		var values *vector.Vector
		if !in.Value.IsReuseID {
			isGenerated := in.Value.Generate(env, utils.NIL)
			if !isGenerated {
				env.SetError(fmt.Sprintf("Los tipos no coinciden para el vector. %v:%v", in.Line, in.Column))
				return nil
			}
			values = in.Value
		} else {
			valuesID := env.GetValueID(in.Value.ID, in.Line, in.Column)
			if valuesID.Type != utils.VECTOR {
				env.SetError(fmt.Sprintf("El ID no almacena un vector. %v:%v", in.Line, in.Column))
				return nil
			}
			reuseVector := *valuesID.Value.(*vector.Vector)
			values = &reuseVector
		}
		env.SaveArray(in.IsVariable, in.Id, values, values.Type.Value.(utils.Type), in.Line, in.Column)
		return nil
	}
	return nil
}

func (in *InitVector) getType(Type utils.Type) string {
	switch Type {
	case utils.INT:
		return "Int"
	case utils.FLOAT:
		return "Float"
	case utils.STRING:
		return "String"
	case utils.BOOLEAN:
		return "Bool"
	case utils.CHAR:
		return "Character"
	default:
		return "nil"
	}
}
