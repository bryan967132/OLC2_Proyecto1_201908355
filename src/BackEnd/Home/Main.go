package main

import (
	interfaces "TSwift/Classes/Interfaces"
	listener "TSwift/Language"
	parser "TSwift/Language/Parser"
	"fmt"
	"io/ioutil"
	"sort"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	filePath := "../../../Inputs/Input5.swift"
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
	Code := listener.Code
	sort.Slice(Code, func(i, j int) bool { return true })
	for _, inst := range Code {
		inst.(interfaces.Instruction).Exec(nil)
	}
}
