package expressions

import (
	env "TSwift/Classes/Env"
	interfaces "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
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
						var value *utils.ReturnType = c.Args[i].Exp.(interfaces.Expression).Exec(Env).GetCopy()
						var param utils.Param = function.GetParams()[i]
						if param.IsPrimitive {
							if param.Type.Value.(utils.Type) == value.Type {
								if _, exists := (*envFunc.Ids)[param.ID]; !exists {
									(*envFunc.Ids)[param.ID] = &env.Symbol{IsVariable: true, IsPrimitive: true, Value: value, Id: param.ID, Type: value.Type}
									continue
								}
								global.SetError(fmt.Sprintf("No puede haber parámetros distintos con el mismo nombre. %v:%v", c.Line, c.Column))
								return &utils.ReturnType{Value: "nil", Type: utils.NIL}
							}
							global.SetError(fmt.Sprintf("El tipo de dato del parámetro no es el esperado. %v:%v", c.Line, c.Column))
							return &utils.ReturnType{Value: "nil", Type: utils.NIL}
						} else {

						}
						continue
					}
					if function.GetParams()[i].IsInout && c.Args[i].IsInout {
						// POR REFERENCIA
						if _, ok := c.Args[i].Exp.(interfaces.Instruction).(*AccessID); ok {
							var accessReference *AccessID = c.Args[i].Exp.(interfaces.Expression).(*AccessID)
							symbolReference := Env.GetValueID(accessReference.Id, c.Line, c.Column)
							if symbolReference != nil && symbolReference.IsVariable {
								var value *utils.ReturnType = c.Args[i].Exp.(interfaces.Expression).Exec(Env)
								var param utils.Param = function.GetParams()[i]
								if param.IsPrimitive {
									if param.Type.Value.(utils.Type) == value.Type {
										if _, exists := (*envFunc.Ids)[param.ID]; !exists {
											(*envFunc.Ids)[param.ID] = symbolReference
											continue
										}
										global.SetError(fmt.Sprintf("No puede haber parámetros distintos con el mismo nombre. %v:%v", c.Line, c.Column))
										return &utils.ReturnType{Value: "nil", Type: utils.NIL}
									}
									global.SetError(fmt.Sprintf("El tipo de dato del parámetro no es el esperado. %v:%v", c.Line, c.Column))
									return &utils.ReturnType{Value: "nil", Type: utils.NIL}
								} else {

								}
								continue
							}
							global.SetError(fmt.Sprintf("No se encuentra la variable referenciada. %v:%v", c.Line, c.Column))
							return &utils.ReturnType{Value: "nil", Type: utils.NIL}
						}
						global.SetError(fmt.Sprintf("No pueden enviarse constantes por referencia. %v:%v", c.Line, c.Column))
						return &utils.ReturnType{Value: "nil", Type: utils.NIL}
					}
					if function.GetParams()[i].IsInout && !c.Args[i].IsInout {
						global.SetError(fmt.Sprintf("El parámetro no se envió por referencia. Falta '&'. %v:%v", c.Line, c.Column))
						return &utils.ReturnType{Value: "nil", Type: utils.NIL}
					}
					if !function.GetParams()[i].IsInout && c.Args[i].IsInout {
						global.SetError(fmt.Sprintf("El parámetro no se envió por valor. No hace falta '&'. %v:%v", c.Line, c.Column))
						return &utils.ReturnType{Value: "nil", Type: utils.NIL}
					}
				}
				global.SetError(fmt.Sprintf("Mal uso de nombre externo. %v:%v", c.Line, c.Column))
				return &utils.ReturnType{Value: "nil", Type: utils.NIL}
			}
			var execute *utils.ReturnType = function.GetBlock().Exec(envFunc)
			if execute != nil {
				if execute.Value == utils.RETURN {
					return nil
				}
				return execute
			}
			return &utils.ReturnType{Value: "nil", Type: utils.NIL}
		}
		global.SetError(fmt.Sprintf("Cantidad errónea de parámetros enviados. %v:%v", c.Line, c.Column))
		return &utils.ReturnType{Value: "nil", Type: utils.NIL}
	}
	return &utils.ReturnType{Value: "nil", Type: utils.NIL}
}
