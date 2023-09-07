package controller

import (
	env "TSwift/Classes/Env"
	instructions "TSwift/Classes/Instructions"
	interfaces "TSwift/Classes/Interfaces"
	utils "TSwift/Classes/Utils"
	listener "TSwift/Language"
	parser "TSwift/Language/Parser"
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/gofiber/fiber/v2"
)

type Controller struct{}

type Parser struct {
	Code string `json:"code"`
}

func (c Controller) Running(ctx *fiber.Ctx) error {
	return ctx.SendString("Interpreter is running!!!")
}

func (c Controller) Parser(ctx *fiber.Ctx) error {
	var reqBody Parser
	if err := ctx.BodyParser(&reqBody); err != nil {
		return ctx.JSON(fiber.Map{
			"console": "¡Ha ocurrido un error!",
		})
	}

	inputStream := antlr.NewInputStream(reqBody.Code)
	scanner := parser.NewScanner(inputStream)
	tokens := antlr.NewCommonTokenStream(scanner, antlr.TokenDefaultChannel)
	parser := parser.NewParserParser(tokens)
	parser.RemoveErrorListeners()
	parserErrors := listener.NewTSwfitErrorListener()
	parser.AddErrorListener(parserErrors)

	parser.BuildParseTrees = true
	tree := parser.Init()
	var listener *listener.TSwfitListener = listener.NewTSwfitListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	utils.PrintConsole = []string{}
	utils.Errors = []string{}

	global := env.NewEnv(nil, "Global")
	for _, fail := range parserErrors.Errors {
		global.SetError(fmt.Sprintf("%v. %v:%v", Replaces(fail.Msg), fail.Line, fail.Column))
	}

	execute := listener.Code
	for _, instruction := range execute {
		defer func() {
			if r := recover(); r != nil {
				global.SetError(fmt.Sprintf("%v %v:%v", r, instruction.(interfaces.Instruction).LineN(), instruction.(interfaces.Instruction).ColumnN()))
			}
		}()
		if _, ok := instruction.(interfaces.Instruction).(*instructions.Function); ok {
			instruction.(interfaces.Instruction).Exec(global)
		}
	}
	for _, instruction := range execute {
		defer func() {
			if r := recover(); r != nil {
				global.SetError(fmt.Sprintf("%v %v:%v", r, instruction.(interfaces.Instruction).LineN(), instruction.(interfaces.Instruction).ColumnN()))
			}
		}()
		if _, ok := instruction.(interfaces.Instruction).(*instructions.Function); !ok {
			instruction.(interfaces.Instruction).Exec(global)
		}
	}
	return ctx.JSON(fiber.Map{
		"console": utils.GetStringOuts(),
	})
}

func (c Controller) GetAST(ctx *fiber.Ctx) error {
	var reqBody struct {
		Code string `json:"code"`
	}

	if err := ctx.BodyParser(&reqBody); err != nil {
		return err
	}

	// Tu lógica de obtención de AST aquí

	// Ejemplo de respuesta

	return ctx.JSON(fiber.Map{
		"AST": "Tu AST aquí",
	})
}

func (c Controller) GetSymbolsTable(ctx *fiber.Ctx) error {
	// Tu lógica de obtención de tabla de símbolos aquí

	return ctx.JSON(fiber.Map{
		"Table": "Tu AST aquí",
	})
}

func (c Controller) GetErrors(ctx *fiber.Ctx) error {
	// Tu lógica de obtención de errores aquí

	// Ejemplo de respuesta

	return ctx.JSON(fiber.Map{
		"Errors": "Tus errores aquí",
	})
}

func Replaces(msg string) string {
	replaces := [][]string{
		{"mismatched input", "Inesperado:"},
		{"expecting", ". Se esperaba:"},
		{"missing", "Inesperado:"},
		{"extraneous input", "Entrada extraña: "},
		{"input", "entrada"},
		{"at", "en"},
		{"no viable alternenive", "Sin recuperar"},
	}
	for _, m := range replaces {
		msg = strings.Replace(msg, m[0], m[1], -1)
	}
	return msg
}

func TreeDot(tree antlr.Tree, ruleNames []string) string {
	dot := "digraph Tree {"
	dot += "\n\tgraph[fontname=\"Arial\" labelloc=\"t\" bgcolor=\"#252526\"];"
	dot += "\n\tnode[fontname=\"Arial\" fontsize=\"8\" width=\"0\" height=\"0\" color=\"white\" fontcolor=\"white\"];"
	dot += "\n\tedge[fontname=\"Arial\" color=\"white\" dir=none];"
	dot += NodesTree("i", tree, ruleNames)
	dot += "\n}"
	return dot
}

func NodesTree(id string, tree antlr.Tree, ruleNames []string) string {
	s := antlr.TreesGetNodeText(tree, ruleNames, nil)
	s = antlr.EscapeWhitespace(s, false)
	c := tree.GetChildCount()
	res := "\n\tn" + id + "[label=\"" + strings.Replace(s, "\"", "\\\"", -1) + "\"];"
	for i := 0; i < c; i++ {
		res += NodesTree(fmt.Sprintf("%s%v", id, i), tree.GetChild(i), ruleNames)
		res += "\n\tn" + id + " -> " + "n" + fmt.Sprintf("%s%v", id, i) + ";"
	}
	return res
}
