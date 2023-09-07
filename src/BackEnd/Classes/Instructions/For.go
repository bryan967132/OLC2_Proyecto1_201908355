package instructions

import (
	env "TSwift/Classes/Env"
	interfaces "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
	vector "TSwift/Classes/Vector"
	"fmt"
)

type For struct {
	Line     int
	Column   int
	TypeInst utils.TypeInst
	IDIter   string
	Exp      interfaces.Expression
	LimInf   interfaces.Expression
	LimSup   interfaces.Expression
	Block    interfaces.Instruction
}

func NewFor(line, column int, idIte string, exp, limInf, limSup interfaces.Expression, block interfaces.Instruction) *For {
	return &For{line, column, utils.LOOP_FOR, idIte, exp, limInf, limSup, block}
}

func (f *For) LineN() int {
	return f.Line
}

func (f *For) ColumnN() int {
	return f.Column
}

func (f *For) Exec(Env *env.Env) *utils.ReturnType {
	envFor := env.NewEnv(Env, Env.Name+" For")
	if f.Exp == nil {
		limInf := f.LimInf.Exec(Env)
		if limInf.Type != utils.INT {
			Env.SetError(fmt.Sprintf("Tipo inválido para rango. %v:%v", f.Line, f.Column))
			return nil
		}
		limSup := f.LimSup.Exec(Env)
		if limSup.Type != utils.INT {
			Env.SetError(fmt.Sprintf("Tipo inválido para rango. %v:%v", f.Line, f.Column))
			return nil
		}
		envFor.SaveID(false, f.IDIter, limInf, utils.INT, f.Line, f.Column)
		for i := limInf.Value.(int); i <= limSup.Value.(int); i++ {
			(*(envFor.Ids))[f.IDIter].Value.(*utils.ReturnType).Value = i
			block := f.Block.Exec(envFor)
			if block != nil {
				if block.Value.(utils.TypeInst) == utils.CONTINUE {
					continue
				}
				if block.Value.(utils.TypeInst) == utils.BREAK {
					break
				}
				return block
			}
		}
		return nil
	}
	exp := f.Exp.Exec(Env)
	if exp.Type == utils.STRING {
		envFor.SaveID(false, f.IDIter, &utils.ReturnType{Value: fmt.Sprintf("%v", exp.Value)[0], Type: utils.STRING}, utils.STRING, f.Line, f.Column)
		for _, char := range fmt.Sprintf("%v", exp.Value) {
			(*(envFor.Ids))[f.IDIter].Value.(*utils.ReturnType).Value = fmt.Sprintf("%c", rune(char))
			block := f.Block.Exec(envFor)
			if block != nil {
				if block.Value == utils.CONTINUE {
					continue
				}
				if block.Value == utils.BREAK {
					break
				}
				return block
			}
		}
		return nil
	}
	if exp.Type == utils.VECTOR {
		vec := exp.Value.(*vector.Vector)
		if vec.IsMatrix {
			envFor.SaveID(false, f.IDIter, &utils.ReturnType{Value: vec.Vectors[0], Type: utils.VECTOR}, utils.VECTOR, f.Line, f.Column)
			for _, element := range vec.Vectors {
				(*envFor.Ids)[f.IDIter].Value.(*utils.ReturnType).Value = element
				block := f.Block.Exec(envFor)
				if block != nil {
					if block.Value == utils.CONTINUE {
						continue
					}
					if block.Value == utils.BREAK {
						break
					}
					return block
				}
			}
		} else {
			if len(exp.Value.(*vector.Vector).Values) == 0 {
				return nil
			}
			envFor.SaveID(false, f.IDIter, &utils.ReturnType{Value: vec.Values[0].Value, Type: vec.Values[0].Type}, vec.Values[0].Type, f.Line, f.Column)
			for _, element := range vec.Values {
				(*envFor.Ids)[f.IDIter].Value.(*utils.ReturnType).Value = element.Value
				block := f.Block.Exec(envFor)
				if block != nil {
					if block.Value == utils.CONTINUE {
						continue
					}
					if block.Value == utils.BREAK {
						break
					}
					return block
				}
			}
		}
		return nil
	}
	Env.SetError(fmt.Sprintf("No se puede iterar en el valor. %v:%v", f.Line, f.Column))
	return nil
}
