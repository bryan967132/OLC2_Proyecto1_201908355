package expressions

import (
	env "TSwift/Classes/Env"
	utils "TSwift/Classes/Utils"
)

type AccessID struct {
	Line    int
	Column  int
	TypeExp utils.TypeExp
	Id      string
}

func NewAccessID(line, column int, id string) *AccessID {
	return &AccessID{line, column, utils.ACCESS_ID, id}
}

func (ac *AccessID) LineN() int {
	return ac.Line
}

func (ac *AccessID) ColumnN() int {
	return ac.Column
}

func (ac *AccessID) Exec(env *env.Env) *utils.ReturnType {
	value := env.GetValueID(ac.Id, ac.Line, ac.Column)
	if value != nil {
		if value.Type == utils.VECTOR {
			return &utils.ReturnType{Value: value.Value, Type: utils.VECTOR}
		}
		return &utils.ReturnType{Value: value.Value.(*utils.ReturnType).Value, Type: value.Value.(*utils.ReturnType).Type}
	}
	return &utils.ReturnType{Value: "nil", Type: utils.NIL}
}
