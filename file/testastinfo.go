package file

import (
	"fmt"
	"go/parser"
	"go/token"
	"log"
)

func ain() {
	filename := "D:\\projects\\go_projects\\src\\github.com\\gin-gonic\\gin\\gin.go"
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, nil, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(f)
}
