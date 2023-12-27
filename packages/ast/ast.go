package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"
	"reflect"
	"strings"
)

func main() {
	// 解析 Go 源码文件
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "packages/ast/interface_definition.go", nil, parser.ParseComments)
	if err != nil {
		log.Panicln(err.Error())
	}

	outFile, err := parser.ParseFile(fset, "packages/ast/out.go", "package "+file.Name.Name, parser.ParseComments)
	// 遍历 AST，寻找接口
	ast.Inspect(file, func(node ast.Node) bool {
		if genDecl, ok := node.(*ast.GenDecl); ok && genDecl.Tok == token.TYPE {
			for _, spec := range genDecl.Specs {
				if typeSpec, ok := spec.(*ast.TypeSpec); ok {
					if intf, ok := typeSpec.Type.(*ast.InterfaceType); ok {
						GenInterfaceDefaultImpl(typeSpec.Name.Name, intf, outFile)
					}
				}
			}
		}
		return true
	})

	if err := SaveASTToFile("packages/ast/out.go", fset, outFile); err != nil {
		log.Panicln(err.Error())
	}
}

// GenInterfaceDefaultImpl 生成接口的默认实现
func GenInterfaceDefaultImpl(name string, intf *ast.InterfaceType, file *ast.File) {
	structName := name + "DefaultImpl"
	// 创建结构体的类型声明
	typeSpec := &ast.TypeSpec{
		Name: ast.NewIdent(structName),
		Type: &ast.StructType{
			Fields: &ast.FieldList{},
		},
	}
	// 创建结构体的声明
	decl := &ast.GenDecl{
		Tok:   token.TYPE,
		Specs: []ast.Spec{typeSpec},
	}
	// 将结构体的声明添加到文件的声明列表中
	file.Decls = append(file.Decls, decl)
	for _, method := range intf.Methods.List {
		// 方法名称
		var methodName string
		for _, name := range method.Names {
			methodName = name.Name
		}
		// TO DO：支持接口继承
		if methodName == "" {
			log.Panicln("no method name found")
		}
		methodType, ok := method.Type.(*ast.FuncType)
		if !ok {
			continue
		}
		// 创建方法的接收者
		recv := &ast.FieldList{
			List: []*ast.Field{
				{
					Names: []*ast.Ident{ast.NewIdent(GetFirstLower(structName))},
					Type:  &ast.StarExpr{X: ast.NewIdent(structName)},
				},
			},
		}
		// 创建方法的参数
		params := &ast.FieldList{
			List: methodType.Params.List,
		}
		// 创建方法的输出结果
		results := &ast.FieldList{
			List: nil,
		}
		if methodType.Results != nil {
			results = &ast.FieldList{
				List: methodType.Results.List,
			}
		}
		// 创建方法的函数体
		body := &ast.BlockStmt{
			List: []ast.Stmt{
				CreateReturnZeroStmt(results),
			},
		}
		logArgsStmts := CreateLogArgsStatements(params)
		// 在函数体的开头插入打印参数的语句
		body.List = append(logArgsStmts, body.List...)
		// 创建方法声明
		funcDecl := &ast.FuncDecl{
			Name: ast.NewIdent(methodName),
			Recv: recv,
			Type: &ast.FuncType{
				Params:  params,
				Results: results,
			},
			Body: body,
		}

		// 将方法的声明添加到文件的声明列表中
		file.Decls = append(file.Decls, funcDecl)

		// 导入使用的包
		if !HasImportedPackage(file, "log") {
			AddImport(file, "log")
		}
	}
}

// CreateReturnZeroStmt 创建返回零值的语句
func CreateReturnZeroStmt(results *ast.FieldList) *ast.ReturnStmt {
	// 创建返回值列表
	var resultsExpr []ast.Expr
	for _, field := range results.List {
		switch ident := reflect.ValueOf(field.Type).Elem().Interface().(type) {
		case ast.Ident:
			resultsExpr = append(resultsExpr, CreateZeroExpr(ident.Name))
		default:
			resultsExpr = append(resultsExpr, &ast.Ident{Name: "nil"})
		}
	}
	// 创建返回语句
	return &ast.ReturnStmt{Results: resultsExpr}
}

// CreateZeroExpr 创建零值的 AST 表达式
func CreateZeroExpr(name string) ast.Expr {
	switch name {
	case "bool":
		return &ast.Ident{Name: "false"}
	case "int", "int8", "int16", "int32", "int64":
		return &ast.BasicLit{Kind: token.INT, Value: "0"}
	case "uint", "uint8", "uint16", "uint32", "uint64", "uintptr":
		return &ast.BasicLit{Kind: token.INT, Value: "0"}
	case "float32", "float64":
		return &ast.BasicLit{Kind: token.FLOAT, Value: "0.0"}
	case "complex64", "complex128":
		return &ast.BasicLit{Kind: token.IMAG, Value: "0.0i"}
	case "string":
		return &ast.BasicLit{Kind: token.STRING, Value: `""`}
	default:
		return &ast.Ident{Name: "nil"}
	}
}

// HasImportedPackage 检查是否已导入指定的包
func HasImportedPackage(file *ast.File, packageName string) bool {
	for _, imp := range file.Imports {
		if imp.Path.Value == fmt.Sprintf(`"%s"`, packageName) {
			return true
		}
	}
	return false
}

// AddImport 添加导入语句
func AddImport(file *ast.File, packageName string) {
	// 创建导入的 AST 节点
	importSpec := &ast.ImportSpec{
		Path: &ast.BasicLit{
			Kind:  token.STRING,
			Value: fmt.Sprintf(`"%s"`, packageName),
		},
	}
	// 创建导入声明
	importDecl := &ast.GenDecl{
		Tok:   token.IMPORT,
		Specs: []ast.Spec{importSpec},
	}
	file.Imports = append(file.Imports, importSpec)
	// 将导入声明添加到文件的声明列表
	file.Decls = append([]ast.Decl{importDecl}, file.Decls...)
}

// CreatePrintStatement 创建打印语句
func CreatePrintStatement(paramName string) ast.Stmt {
	// log.Println
	printFunc := &ast.SelectorExpr{
		X:   ast.NewIdent("log"),
		Sel: ast.NewIdent("Println"),
	}
	// "<ParamName> = "，比如 "name = "
	paramNameString := &ast.BasicLit{
		Kind:  token.STRING,
		Value: fmt.Sprintf(`"%s = "`, paramName),
	}
	// <ParamName>，比如 name
	paramNameExpr := ast.NewIdent(paramName)
	// log.Println("<ParamName> = ", <ParamName>)，比如 log.Println("name = ", name)
	callExpr := &ast.CallExpr{
		Fun:  printFunc,
		Args: []ast.Expr{paramNameString, paramNameExpr},
	}
	return &ast.ExprStmt{
		X: callExpr,
	}
}

// CreateLogArgsStatements 创建打印参数的语句
func CreateLogArgsStatements(params *ast.FieldList) []ast.Stmt {
	stmts := make([]ast.Stmt, 0)
	for _, param := range params.List {
		for _, name := range param.Names {
			printStmt := CreatePrintStatement(name.Name)
			stmts = append(stmts, printStmt)
		}
	}
	return stmts
}

// SaveASTToFile 将格式化后的代码保存到文件
func SaveASTToFile(fileName string, fset *token.FileSet, node *ast.File) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// 格式化代码
	var buf bytes.Buffer
	err = format.Node(&buf, fset, node)
	if err != nil {
		return err
	}

	// 将结果写入文件
	_, err = file.WriteString(buf.String())
	return err
}

// GetFirstLower 首先获取给定的字符串的首字母，然后将其转成小写字母
func GetFirstLower(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToLower(string(s[0]))
}
