package example

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestTraverseAST(t *testing.T) {
	fset := token.NewFileSet()                       // 创建file set
	f, err := parser.ParseFile(fset, "", srcCode, 0) // 创建ast.File
	if err != nil {
		panic(err)
	}
	ast.Inspect(f, func(node ast.Node) bool {
		ast.Print(fset, node) // 读到叶子节点就会输出  nil
		return true
	})
}
