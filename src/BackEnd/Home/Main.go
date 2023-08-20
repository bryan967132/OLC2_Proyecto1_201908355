package main

import (
	env "TSwift/Classes/Env"
	interfaces "TSwift/Classes/Interfaces"
	listener "TSwift/Language"
	parser "TSwift/Language/Parser"
	"fmt"
	"io/ioutil"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	filePath := "../../../Inputs/Input10.swift"
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
	parser.BuildParseTrees = true
	tree := parser.Init()
	var listener *listener.TSwfitListener = listener.NewTSwfitListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	execute := listener.Code
	global := env.NewEnv(nil, "Global")
	for _, instruction := range execute {
		instruction.(interfaces.Instruction).Exec(global)
	}
	global.PrintPrints()
	global.PrintErrors()
}
