package interfaces

import (
	env "TSwift/Classes/Env"
)

type Instruction interface {
	Exec(env *env.Env) interface{}
}
