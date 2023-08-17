package instructions

import (
	env "TSwift/Classes/Env"
	interfaces "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
	"fmt"
)

type InitID struct {
	Line     int
	Column   int
	TypeInst utils.TypeInst
	Id       string
	Type     utils.Type
	Value    interfaces.Expression
}

func NewInitID(line, column int, id string, Type utils.Type, value interfaces.Expression) *InitID {
	return &InitID{line, column, utils.INIT_ID, id, Type, value}
}

func (in *InitID) LineN() int {
	return in.Line
}

func (in *InitID) ColumnN() int {
	return in.Column
}

func (in *InitID) Exec(env *env.Env) interface{} {
	if in.Value != nil {
		value := in.Value.Exec(env)
		if in.Type != utils.NIL {
			if in.Type == value.Type {
				env.SaveID(in.Id, value.Value, value.Type, in.Line, in.Column)
				return nil
			}
			env.SetError(fmt.Sprintf("Los tipos no coinciden en la declaración. %v:%v", in.Line, in.Column))
			return nil
		}
		env.SaveID(in.Id, value.Value, value.Type, in.Line, in.Column)
	} else {
		env.SaveID(in.Id, &utils.ReturnType{Value: "nil", Type: utils.NIL}, in.Type, in.Line, in.Column)
	}
	return nil
}
