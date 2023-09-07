package expressions

import (
	env "TSwift/Classes/Env"
	interfaces "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
	vector "TSwift/Classes/Vector"
	"fmt"
)

type CallFunction struct {
	Line    int
	Column  int
	TypeExp utils.TypeExp
	ID      string
	Args    []utils.Arg
}

func NewCallFunction(line, column int, id string, args []utils.Arg) *CallFunction {
	return &CallFunction{line, column, utils.CALL_FUNC, id, args}
}

func (c *CallFunction) LineN() int {
	return c.Line
}

func (c *CallFunction) ColumnN() int {
	return c.Column
}

func (c *CallFunction) Exec(Env *env.Env) *utils.ReturnType {
	global := Env.GetGlobal()
	function := (*global.GetFunction(c.ID, c.Line, c.Column)).(interfaces.Function)
	if function != nil {
		if len(function.GetParams()) == len(c.Args) {
			envFunc := env.NewEnv(global, "Funcion "+c.ID)
			for i := 0; i < len(function.GetParams()); i++ {
				if function.GetParams()[i].ExternID == c.Args[i].ExternID {
					if !function.GetParams()[i].IsInout && !c.Args[i].IsInout {
						// POR VALOR
						var value *utils.ReturnType = c.Args[i].Exp.(interfaces.Expression).Exec(Env)
						var param utils.Param = function.GetParams()[i]
						if param.IsPrimitive {
							if param.Type.Value.(utils.Type) == value.Type {
								if _, exists := (*envFunc.Ids)[param.ID]; !exists {
									(*envFunc.Ids)[param.ID] = &env.Symbol{IsVariable: true, IsPrimitive: true, Value: value, Id: param.ID, Type: value.Type}
									continue
								}
								global.SetError(fmt.Sprintf("No puede haber parámetros distintos con el mismo nombre. %v:%v", c.Line, c.Column))
								return nil
							}
							global.SetError(fmt.Sprintf("El tipo de dato del parámetro no es el esperado. %v:%v", c.Line, c.Column))
							return nil
						} else if param.IsVector {
							if _, okID := c.Args[i].Exp.(interfaces.Expression).(*AccessID); okID {
								var accessReference *AccessID = c.Args[i].Exp.(interfaces.Expression).(*AccessID)
								var typeVec utils.VectorType = param.Type.Value.(utils.VectorType)
								symbolReference := Env.GetValueID(accessReference.Id, c.Line, c.Column)
								if typeVec.Type.Value.(utils.Type) == symbolReference.ArrType {
									if _, exists := (*envFunc.Ids)[param.ID]; !exists {
										(*envFunc.Ids)[param.ID] = &env.Symbol{IsVariable: true, IsPrimitive: true, Value: symbolReference.Value, Id: param.ID, Type: utils.VECTOR, ArrType: symbolReference.ArrType}
										continue
									}
									global.SetError(fmt.Sprintf("No puede haber parámetros distintos con el mismo nombre. %v:%v", c.Line, c.Column))
									return nil
								}
								global.SetError(fmt.Sprintf("El tipo del vector no es el esperado. %v:%v", c.Line, c.Column))
								return nil
							} else if _, okV := c.Args[i].Exp.(interfaces.Expression).(*AccessArray); okV {
								var accessReference *AccessArray = c.Args[i].Exp.(interfaces.Expression).(*AccessArray)
								var typeVec utils.VectorType = param.Type.Value.(utils.VectorType)
								symbolReference := Env.GetValueID(accessReference.Id, c.Line, c.Column)
								getArray := accessReference.Exec(Env)
								if typeVec.Type.Value.(utils.Type) == symbolReference.ArrType {
									if _, exists := (*envFunc.Ids)[param.ID]; !exists {
										(*envFunc.Ids)[param.ID] = &env.Symbol{IsVariable: true, IsPrimitive: true, Value: getArray.Value, Id: param.ID, Type: utils.VECTOR, ArrType: symbolReference.ArrType}
										continue
									}
									global.SetError(fmt.Sprintf("No puede haber parámetros distintos con el mismo nombre. %v:%v", c.Line, c.Column))
									return nil
								}
								global.SetError(fmt.Sprintf("El tipo del vector no es el esperado. %v:%v", c.Line, c.Column))
								return nil
							}
							global.SetError(fmt.Sprintf("No pueden enviarse arreglos directamente. %v:%v", c.Line, c.Column))
							return nil
						}
						global.SetError(fmt.Sprintf("El tipo de dato no se reconoce. %v:%v", c.Line, c.Column))
						return nil
					}
					if function.GetParams()[i].IsInout && c.Args[i].IsInout {
						// POR REFERENCIA
						var param utils.Param = function.GetParams()[i]
						if param.IsPrimitive {
							var accessReference *AccessID = c.Args[i].Exp.(interfaces.Expression).(*AccessID)
							symbolReference := Env.GetValueID(accessReference.Id, c.Line, c.Column)
							if _, okID := c.Args[i].Exp.(interfaces.Instruction).(*AccessID); okID {
								if param.Type.Value.(utils.Type) == symbolReference.Type {
									if _, exists := (*envFunc.Ids)[param.ID]; !exists {
										(*envFunc.Ids)[param.ID] = symbolReference
										continue
									}
									global.SetError(fmt.Sprintf("No puede haber parámetros distintos con el mismo nombre. %v:%v", c.Line, c.Column))
									return nil
								}
								global.SetError(fmt.Sprintf("El tipo de dato del parámetro no es el esperado. %v:%v", c.Line, c.Column))
								return nil
							}
							global.SetError(fmt.Sprintf("El tipo de dato no se reconoce. %v:%v", c.Line, c.Column))
							return nil
						} else if param.IsVector {
							if _, okID := c.Args[i].Exp.(interfaces.Instruction).(*AccessID); okID {
								var accessReference *AccessID = c.Args[i].Exp.(interfaces.Expression).(*AccessID)
								var typeVec utils.VectorType = param.Type.Value.(utils.VectorType)
								symbolReference := Env.GetValueID(accessReference.Id, c.Line, c.Column)
								if typeVec.Type.Value.(utils.Type) == symbolReference.ArrType {
									if _, exists := (*envFunc.Ids)[param.ID]; !exists {
										(*envFunc.Ids)[param.ID] = symbolReference
										continue
									}
									global.SetError(fmt.Sprintf("No puede haber parámetros distintos con el mismo nombre. %v:%v", c.Line, c.Column))
									return nil
								}
								global.SetError(fmt.Sprintf("El tipo del vector no es el esperado. %v:%v", c.Line, c.Column))
								return nil
							} else if _, okV := c.Args[i].Exp.(interfaces.Expression).(*AccessArray); okV {
								var accessReference *AccessArray = c.Args[i].Exp.(interfaces.Expression).(*AccessArray)
								var typeVec utils.VectorType = param.Type.Value.(utils.VectorType)
								symbolReference := Env.GetValueID(accessReference.Id, c.Line, c.Column)
								if typeVec.Type.Value.(utils.Type) == symbolReference.ArrType {
									if _, exists := (*envFunc.Ids)[param.ID]; !exists {
										indexs := []int{}
										for _, in := range accessReference.Index {
											indexs = append(indexs, in.Exec(Env).Value.(int))
										}
										v := symbolReference.Value.(*vector.Vector).GetPosition(Env, indexs, accessReference.Line, accessReference.Column)
										(*envFunc.Ids)[param.ID] = &env.Symbol{IsVariable: true, IsPrimitive: true, Value: v, Id: param.ID, Type: utils.VECTOR, ArrType: symbolReference.ArrType}
										continue
									}
									global.SetError(fmt.Sprintf("No puede haber parámetros distintos con el mismo nombre. %v:%v", c.Line, c.Column))
									return nil
								}
								global.SetError(fmt.Sprintf("El tipo del vector no es el esperado. %v:%v", c.Line, c.Column))
								return nil
							}
							global.SetError(fmt.Sprintf("No pueden enviarse arreglos directamente. %v:%v", c.Line, c.Column))
							return nil
						}
					}
					if function.GetParams()[i].IsInout && !c.Args[i].IsInout {
						global.SetError(fmt.Sprintf("El parámetro no se envió por referencia. Falta '&'. %v:%v", c.Line, c.Column))
						return nil
					}
					if !function.GetParams()[i].IsInout && c.Args[i].IsInout {
						global.SetError(fmt.Sprintf("El parámetro no se envió por valor. No hace falta '&'. %v:%v", c.Line, c.Column))
						return nil
					}
				}
				global.SetError(fmt.Sprintf("Mal uso de nombre externo. %v:%v", c.Line, c.Column))
				return nil
			}
			var execute *utils.ReturnType = function.GetBlock().Exec(envFunc)
			if execute != nil {
				if execute.Value == utils.RETURN {
					return nil
				}
				return execute
			}
			return nil
		}
		global.SetError(fmt.Sprintf("Cantidad errónea de parámetros enviados. %v:%v", c.Line, c.Column))
		return nil
	}
	return nil
}
