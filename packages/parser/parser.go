package parser

import (
	"fmt"
	"go/parser"
	"go/token"
	"log"
)

// 读取import行号
func Parser() {
	// 要解析的文件路径
	filepath := "example.go"

	// 创建一个 token 文件集合
	fset := token.NewFileSet()

	// 解析源代码文件
	f, err := parser.ParseFile(fset, filepath, nil, parser.ImportsOnly)
	if err != nil {
		log.Fatalf("解析文件失败: %v", err)
	}

	// 打印所有导入语句的行号
	for _, s := range f.Imports {
		fmt.Printf("导入语句 %q 在第 %d 行\n", s.Path.Value, fset.Position(s.Pos()).Line)
	}
}
