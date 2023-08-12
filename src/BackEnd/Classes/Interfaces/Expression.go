package abstracts

import (
	env "Modules/Classes/Env"
	utils "Modules/Classes/Utils"
)

type Expression interface {
	Exec(env *env.Env) *utils.ReturnType
}
