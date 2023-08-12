package abstracts

import (
	env "Modules/Classes/Env"
)

type Instruction interface {
	Exec(env *env.Env) interface{}
}
