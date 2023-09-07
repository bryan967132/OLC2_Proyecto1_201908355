package instructions

import (
	env "TSwift/Classes/Env"
	interfaces "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
	"reflect"
)

type Switch struct {
	Line     int
	Column   int
	TypeInst utils.TypeInst
	Arg      interfaces.Expression
	Cases    interface{}
	Default  interfaces.Instruction
}

func NewSwitch(line, column int, arg interfaces.Expression, cases interface{}, _default interfaces.Instruction) *Switch {
	return &Switch{line, column, utils.SWITCH, arg, cases, _default}
}

func (s *Switch) LineN() int {
	return s.Line
}

func (s *Switch) ColumnN() int {
	return s.Column
}

func (s *Switch) Exec(Env *env.Env) *utils.ReturnType {
	envSwitch := env.NewEnv(Env, Env.Name+"Switch")
	if s.Cases != nil {
		arg := s.Arg.Exec(Env)
		for _, case_ := range s.ToSlice(s.Cases) {
			case_.(*Case).SetCase(arg)
			case_exec := case_.(interfaces.Instruction).Exec(envSwitch)
			if case_exec != nil {
				if case_exec.Value == utils.RETURN {
					return nil
				}
				if case_exec.Value == utils.BREAK {
					return nil
				}
				return case_exec
			}
			if case_.(*Case).Flag {
				return nil
			}
		}
	}
	if s.Default != nil {
		default_ := s.Default.Exec(envSwitch)
		if default_ != nil {
			if default_.Value == utils.RETURN {
				return nil
			}
			if default_.Value == utils.BREAK {
				return nil
			}
			return default_
		}
	}
	return nil
}

func (s *Switch) ToSlice(cases interface{}) []interface{} {
	if reflect.TypeOf(cases).Kind() == reflect.Slice {
		value := reflect.ValueOf(cases)
		if value.Type().Elem().Kind() == reflect.Interface {
			interfaceSlice := make([]interface{}, value.Len())
			for i := 0; i < value.Len(); i++ {
				interfaceSlice[i] = value.Index(i).Interface()
			}
			return interfaceSlice
		}
	}
	return nil
}
