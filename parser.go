package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/fatih/structtag"
)

type Model struct {
	Name        string
	PackageName string
	ModuleName  string
	Columns     []Column
}

type Column struct {
	WhereName string
	Name      string
	Type      string
}

func Parse(file string) (Model, error) {
	m := Model{}
	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, file, nil, 4)
	if err != nil {
		return m, err
	}

	ast.Inspect(f, func(n ast.Node) bool {
		switch n := n.(type) {
		case *ast.File:
			m.PackageName = n.Name.Name
		case *ast.TypeSpec:
			if m.Name == "" {
				m.Name = n.Name.Name

				ast.Inspect(n, func(n ast.Node) bool {
					switch n := n.(type) {
					case *ast.Field:
						tags, _ := structtag.Parse(n.Tag.Value[1 : len(n.Tag.Value)-1])

						if tags.Tags()[0].Name != "-" {
							var t string
							switch _t := n.Type.(type) {
							case *ast.SelectorExpr:
								t = fmt.Sprintf("%s.%s", _t.X, _t.Sel)
							case *ast.Ident:
								t = _t.String()
							}
							m.Columns = append(m.Columns, Column{
								WhereName: n.Names[0].Name,
								Name:      tags.Tags()[0].Name,
								Type:      t,
							})
						}
						return false
					}
					return true
				})
			}
		}
		return true
	})

	return m, nil
}
