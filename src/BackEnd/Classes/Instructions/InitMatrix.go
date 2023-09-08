package instructions

import (
	env "TSwift/Classes/Env"
	utils "TSwift/Classes/Utils"
	vector "TSwift/Classes/Vector"
)

type InitMatrix struct {
	Line       int
	Column     int
	IsVariable bool
	TypeInst   utils.TypeInst
	Id         string
	Type       *utils.VectorType
	Value      *vector.Vector
}

func NewInitMatrix(line, column int, isVar bool, id string, Type *utils.VectorType, value *vector.Vector) *InitMatrix {
	return &InitMatrix{line, column, isVar, utils.INIT_ID, id, Type, value}
}

func (in *InitMatrix) LineN() int {
	return in.Line
}

func (in *InitMatrix) ColumnN() int {
	return in.Column
}

func (in *InitMatrix) Exec(env *env.Env) *utils.ReturnType {
	if in.Type != nil {
		if !in.Value.IsRepeating {
			isGenerated := in.Value.Generate(env, in.Type.Type.Value.(utils.Type))
			if !isGenerated {
				env.SetError("Los tipos no coinciden para el vector", in.Line, in.Column)
				return nil
			}
			if in.Value.Dims == in.Type.Length {
				env.SaveArray(in.IsVariable, in.Id, in.Value, in.Value.Type.Value.(utils.Type), in.Line, in.Column)
				return nil
			}
			env.SetError("Las dimensiones no concuerdan", in.Line, in.Column)
			return nil
		}
		in.Value = in.Value.GenerateRepeating(env, in.Line, in.Column, in.Type.Type.Value.(utils.Type))
		if in.Value != nil {
			env.SaveArray(in.IsVariable, in.Id, in.Value, in.Value.Type.Value.(utils.Type), in.Line, in.Column)
		}
		return nil
	}
	if !in.Value.IsRepeating {
		isGenerated := in.Value.Generate(env, utils.NIL)
		if !isGenerated {
			env.SetError("Los tipos no coinciden para el vector", in.Line, in.Column)
			return nil
		}
		env.SaveArray(in.IsVariable, in.Id, in.Value, in.Value.Type.Value.(utils.Type), in.Line, in.Column)
		return nil
	}
	in.Value = in.Value.GenerateRepeating(env, in.Line, in.Column, utils.NIL)
	if in.Value != nil {
		env.SaveArray(in.IsVariable, in.Id, in.Value, in.Value.Type.Value.(utils.Type), in.Line, in.Column)
	}
	return nil
}

func (in *InitMatrix) getType(Type utils.Type) string {
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
