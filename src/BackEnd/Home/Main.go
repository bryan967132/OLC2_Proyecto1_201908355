package main

import (
	parser "TSwift/Language/Parser"
	"fmt"
	"io/ioutil"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	filePath := "../../../Inputs/Input4.swift"
	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error al leer el archivo:", err)
		return
	}
	input := string(fileData)
	fmt.Println(input)
	inputStream := antlr.NewInputStream(input)
	scanner := parser.NewScanner(inputStream)
	stream := antlr.NewCommonTokenStream(scanner, antlr.TokenDefaultChannel)
	parser := parser.NewParserParser(stream)
	antlr.TreesDescendants(parser.Init())
}
