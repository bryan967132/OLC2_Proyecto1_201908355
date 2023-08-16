package env

import "fmt"

type Env struct {
	Prints   []string
	Errors   []string
	previous *Env
	name     string
}

func NewEnv(previous *Env, name string) *Env {
	return &Env{[]string{}, []string{}, previous, name}
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
