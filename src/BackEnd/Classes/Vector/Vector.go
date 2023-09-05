package vector

import (
	env "TSwift/Classes/Env"
	interfaces "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
	"fmt"
)

type Vector struct {
	IsMatrix  bool
	IsReuseID bool
	Length    int
	Dims      int
	Type      *utils.AttribsType
	Vectors   []*Vector
	Elements  []interfaces.Expression
	Values    []*utils.ReturnType
	ID        string
}

func NewMatrix(Type *utils.AttribsType, Vectors []*Vector) *Vector {
	return &Vector{IsMatrix: true, IsReuseID: false, Type: Type, Vectors: Vectors}
}

func NewVector(Type *utils.AttribsType, Elements []interfaces.Expression) *Vector {
	return &Vector{IsMatrix: false, IsReuseID: false, Type: Type, Elements: Elements}
}

func NewReuseVector(id string) *Vector {
	return &Vector{IsMatrix: false, IsReuseID: true, ID: id}
}

func (v *Vector) Generate(env *env.Env, Type utils.Type) bool {
	if v.IsMatrix {
		v.Vectors = []*Vector{}
		for i := 0; i < len(v.Vectors); i++ {
			v.Vectors[i].Generate(env, Type)
		}
		v.Length = len(v.Vectors)
		v.Dims = v.Vectors[0].Dims + 1
		return true
	}
	v.Values = []*utils.ReturnType{}
	for i := 0; i < len(v.Elements); i++ {
		v.Values = append(v.Values, v.Elements[i].Exec(env))
		if &Type == nil {
			Type = v.Values[i].Type
		}
		if v.Values[i].Type != Type {
			return false
		}
	}
	v.Length = len(v.Values)
	v.Dims = 1
	return true
}

func (v *Vector) GetPosition(env *env.Env, indexs []int, line, column int) interface{} {
	if len(indexs) > v.Dims {
		env.SetError(fmt.Sprintf("Las dimensiones no coinciden con las del vector. %v:%v", line, column))
		return nil
	}
	if len(indexs) == 1 {
		if v.IsMatrix {
			if indexs[0] >= 0 && indexs[0] < len(v.Vectors) {
				return v.Vectors[indexs[0]]
			}
			env.SetError(fmt.Sprintf("Índices fuera de rango. %v:%v", line, column))
			return nil
		}
		if indexs[0] >= 0 && indexs[0] < len(v.Values) {
			return v.Values[indexs[0]]
		}
		env.SetError(fmt.Sprintf("Índices fuera de rango. %v:%v", line, column))
		return nil
	}
	if indexs[0] >= 0 && indexs[0] < len(v.Vectors) {
		return v.Vectors[indexs[0]].GetPosition(env, indexs[1:], line, column)
	}
	env.SetError(fmt.Sprintf("Índices fuera de rango. %v:%v", line, column))
	return nil
}

func (v *Vector) SetValuePosition(env *env.Env, indexs []int, newValue interfaces.Expression, line, column int) bool {
	if len(indexs) > v.Dims {
		env.SetError(fmt.Sprintf("Las dimensiones no coinciden con las del vector. %v:%v", line, column))
		return false
	}
	if len(indexs) == 1 {
		if v.IsMatrix {
			if indexs[0] >= 0 && indexs[0] < len(v.Vectors) {
				env.SetError(fmt.Sprintf("Solo puede modificarse un elemento a la vez en un vector. %v:%v", line, column))
				return false
			}
			env.SetError(fmt.Sprintf("Índices fuera de rango. %v:%v", line, column))
			return false
		}
		if indexs[0] >= 0 && indexs[0] < len(v.Values) {
			if v.Values[indexs[0]].Type == newValue.Exec(env).Type {
				v.Elements[indexs[0]] = newValue
				v.Values[indexs[0]] = newValue.Exec(env)
				return true
			}
			env.SetError(fmt.Sprintf("El tipo del nuevo valor no coincide con el que almacena el vector. %v:%v", line, column))
			return false
		}
		env.SetError(fmt.Sprintf("Índices fuera de rango. %v:%v", line, column))
		return false
	}
	if indexs[0] >= 0 && indexs[0] < len(v.Vectors) {
		return v.Vectors[indexs[0]].SetValuePosition(env, indexs[1:], newValue, line, column)
	}
	env.SetError(fmt.Sprintf("Índices fuera de rango. %v:%v", line, column))
	return false
}
