package interfaces

import (
	env "TSwift/Classes/Env"
)

type Instruction interface {
	LineN() int
	ColumnN() int
	Exec(env *env.Env) interface{}
}
