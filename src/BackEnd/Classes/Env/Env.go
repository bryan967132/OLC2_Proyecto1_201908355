package env

import (
	utils "TSwift/Classes/Utils"
	"fmt"
)

type Env struct {
	Ids      *map[string]*Symbol
	Prints   []string
	Errors   []string
	previous *Env
	name     string
}

func NewEnv(previous *Env, name string) *Env {
	return &Env{&map[string]*Symbol{}, []string{}, []string{}, previous, name}
}

func (env *Env) SaveID(id string, value *utils.ReturnType, Type utils.Type, line, column int) bool {
	if _, exists := (*env.Ids)[id]; !exists {
		(*env.Ids)[id] = &Symbol{Value: value, Id: id, Type: Type}
		return true
	}
	env.SetError(fmt.Sprintf("Declaración de variable existente. %v:%v", line, column))
	return false
}

func (env *Env) GetValueID(id string, line, column int) *Symbol {
	var current *Env = env
	for current != nil {
		if _, exists := (*env.Ids)[id]; exists {
			return (*env.Ids)[id]
		}
		current = current.previous
	}
	env.SetError(fmt.Sprintf("Acceso a variable inexistente. %v:%v", line, column))
	return nil
}

func (env *Env) ReasignID(id string, value *utils.ReturnType, line, column int) bool {
	var current *Env = env
	for current != nil {
		if _, exists := (*env.Ids)[id]; exists {
			if (*env.Ids)[id].Type == value.Type {
				(*env.Ids)[id].Value = value
				return true
			}
			env.SetError(fmt.Sprintf("Los tipos no coinciden en la asignación. %v:%v", line, column))
			return false
		}
		current = current.previous
	}
	env.SetError(fmt.Sprintf("Asignación de valor a variable inexistente. %v:%v", line, column))
	return false
}

func (env *Env) SetPrints(print string) {
	env.Prints = append(env.Prints, print)
}

func (env *Env) PrintPrints() {
	fmt.Println("\nTSwift:")
	if len(env.Prints) > 0 {
		for _, Print := range env.Prints {
			fmt.Println(Print)
		}
	}
}

func (env *Env) SetError(errorD string) {
	env.Errors = append(env.Errors, errorD)
}

func (env *Env) PrintErrors() {
	if len(env.Errors) > 0 {
		fmt.Println("\nERRORES:")
		for _, Error := range env.Errors {
			fmt.Println(Error)
		}
	}
}
