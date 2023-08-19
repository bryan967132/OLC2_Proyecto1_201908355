package instructions

import (
	env "TSwift/Classes/Env"
	interfaces "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
)

type Block struct {
	Line         int
	Column       int
	TypeInst     utils.TypeInst
	Instructions []interface{}
}

func NewBlock(line, column int, instructions []interface{}) *Block {
	return &Block{line, column, utils.BLOCK_INST, instructions}
}

func (bk *Block) LineN() int {
	return bk.Line
}

func (bk *Block) ColumnN() int {
	return bk.Column
}

func (bk *Block) Exec(Env *env.Env) interface{} {
	newEnv := env.NewEnv(Env, Env.Name)
	var ret interface{}
	for _, instruction := range bk.Instructions {
		ret = (instruction.(interfaces.Instruction)).Exec(newEnv)
		if ret != nil {
			return ret
		}
	}
	return nil
}