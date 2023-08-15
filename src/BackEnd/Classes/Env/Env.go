package env

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
