package main

import (
	env "TSwift/Classes/Env"
	interfaces "TSwift/Classes/Interfaces"
	listener "TSwift/Language"
	parser "TSwift/Language/Parser"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	filePath := "../../../Inputs/Input16.swift"
	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error al leer el archivo:", err)
		return
	}
	input := string(fileData)
	inputStream := antlr.NewInputStream(input)
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

	//fmt.Println(TreeDot(tree, parser.RuleNames))

	global := env.NewEnv(nil, "Global")
	for _, fail := range parserErrors.Errors {
		global.SetError(fmt.Sprintf("%v. %v:%v", Replaces(fail.Msg), fail.Line, fail.Column))
	}

	execute := listener.Code
	for _, instruction := range execute {
		defer func() {
			if r := recover(); r != nil {
				global.SetError(fmt.Sprintf("%v", r))
				global.PrintPrints()
				global.PrintErrors()
			}
		}()
		instruction.(interfaces.Instruction).Exec(global)
	}
	global.PrintPrints()
	global.PrintErrors()
}

func Replaces(msg string) string {
	replaces := [][]string{{"mismatched input", "Inesperado:"}, {"expecting", ". Se esperaba:"}, {"missing", "Inesperado:"}, {"at", "en"}}
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
	res := "\n\tnodo_" + id + "[label=\"" + strings.Replace(s, "\"", "\\\"", -1) + "\"];"
	for i := 0; i < c; i++ {
		res += NodesTree(fmt.Sprintf("%s_%v", id, i), tree.GetChild(i), ruleNames)
		res += "\n\tnodo_" + id + " -> " + "nodo_" + fmt.Sprintf("%s_%v", id, i) + ";"
	}
	return res
}
