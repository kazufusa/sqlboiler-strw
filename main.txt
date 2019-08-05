package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "./measurements.go.test", nil, 4)
	if err != nil {
		fmt.Printf("Failed to parse file\n")
		return
	}

	ast.Print(nil, f)
	// ast.Inspect(f, func(a ast.Node) bool {
	// 	return true
	// })
}

// package main
//
// import (
// 	_ "go/parser"
// 	"io/ioutil"
//
// 	"github.com/hashicorp/hcl/hcl/parser"
// )
//
// func main() {
// 	b, err := ioutil.ReadFile("./measurements.go.test")
// 	if err != nil {
// 		panic(err)
// 	}
// 	sqlboilerCode := string(b)
// 	parser.ParseExpr
// }
