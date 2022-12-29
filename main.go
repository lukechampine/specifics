package main

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(&analysis.Analyzer{
		Name:             "specifics",
		Doc:              "ban generics",
		RunDespiteErrors: true,
		Run: func(pass *analysis.Pass) (interface{}, error) {
			for _, file := range pass.Files {
				ast.Inspect(file, func(n ast.Node) bool {
					switch n := n.(type) {
					case *ast.FuncDecl:
						if n.Type != nil {
							if n.Type.TypeParams != nil {
								pass.Report(analysis.Diagnostic{
									Pos:     n.Pos(),
									Message: "generic function",
								})
								return false
							}
						}
					case *ast.TypeSpec:
						if n.TypeParams != nil {
							pass.Report(analysis.Diagnostic{
								Pos:     n.Pos(),
								Message: "generic type",
							})
							return false
						}
					}
					return true
				})
			}
			return nil, nil
		},
	})
}
