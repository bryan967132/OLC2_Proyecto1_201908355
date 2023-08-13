package interfaces

import (
	env "TSwift/Classes/Env"
	utils "TSwift/Classes/Utils"
)

type Expression interface {
	Exec(env *env.Env) *utils.ReturnType
}
