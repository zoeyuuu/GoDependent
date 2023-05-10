package file

// package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"log"
)

func main() {
	filename := "D:\\projects\\go_projects\\src\\gin\\gin.go"
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, nil, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(f)
}
