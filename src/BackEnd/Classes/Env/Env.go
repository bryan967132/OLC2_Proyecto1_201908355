package env

import "fmt"

type Env struct {
	Prints   []string
	Errors   []string
	previous *Env
	name     string
}

func (env *Env) SetError(errorD string) {
	env.Errors = append(env.Errors, errorD)
}

func NewEnv(previous *Env, name string) *Env {
	return &Env{[]string{}, []string{}, previous, name}
}

func (env *Env) PrintErrors() {
	fmt.Println("\nErrores")
	for _, Error := range env.Errors {
		fmt.Println(Error)
	}
}
