package env

import (
	utils "TSwift/Classes/Utils"
	"fmt"
)

type Env struct {
	Ids      *map[string]*Symbol
	previous *Env
	Name     string
}

func NewEnv(previous *Env, name string) *Env {
	return &Env{&map[string]*Symbol{}, previous, name}
}

func (env *Env) SaveID(isVariable bool, id string, value *utils.ReturnType, Type utils.Type, line, column int) bool {
	if _, exists := (*env.Ids)[id]; !exists {
		(*env.Ids)[id] = &Symbol{IsVariable: isVariable, Value: value, Id: id, Type: Type}
		return true
	}
	env.SetError(fmt.Sprintf("Declaración de variable existente. %v:%v", line, column))
	return false
}

func (env *Env) GetValueID(id string, line, column int) *Symbol {
	var current *Env = env
	for current != nil {
		if _, exists := (*current.Ids)[id]; exists {
			return (*current.Ids)[id]
		}
		current = current.previous
	}
	current.SetError(fmt.Sprintf("Acceso a variable inexistente. %v:%v", line, column))
	return nil
}

func (env *Env) ReasignID(id string, value *utils.ReturnType, line, column int) bool {
	var current *Env = env
	for current != nil {
		if _, exists := (*current.Ids)[id]; exists {
			if (*current.Ids)[id].IsVariable {
				if (*current.Ids)[id].Type == value.Type || (*current.Ids)[id].Type == utils.STRING && value.Type == utils.CHAR || (*current.Ids)[id].Type == utils.FLOAT && value.Type == utils.INT {
					(*current.Ids)[id].Value = value
					return true
				}
				current.SetError(fmt.Sprintf("Los tipos no coinciden en la asignación. Intenta asignar un \"%v\" a un \"%v\". %v:%v", current.getType(value.Type), current.getType((*current.Ids)[id].Type), line, column))
				return false
			}
			current.SetError(fmt.Sprintf("Asignación de valor a constante. %v:%v", line, column))
			return false
		}
		current = current.previous
	}
	current.SetError(fmt.Sprintf("Asignación de valor a variable inexistente. %v:%v", line, column))
	return false
}

func (env *Env) SetPrints(print string) {
	utils.PrintConsole = append(utils.PrintConsole, print)
}

func (env *Env) PrintPrints() {
	fmt.Println("\nTSwift:")
	if len(utils.PrintConsole) > 0 {
		for _, Print := range utils.PrintConsole {
			fmt.Println(Print)
		}
	}
}

func (env *Env) SetError(errorD string) {
	utils.Errors = append(utils.Errors, errorD)
}

func (env *Env) PrintErrors() {
	if len(utils.Errors) > 0 {
		fmt.Println("\nERRORES:")
		for _, Error := range utils.Errors {
			fmt.Println(Error)
		}
	}
}

func (env *Env) getType(Type utils.Type) string {
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
