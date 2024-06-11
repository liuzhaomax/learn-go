package example

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

var srcCode = `
package hello

import "fmt"

func greet() {
    var msg = "Hello World!"
    fmt.Println(msg)
}
`

func TestReadAST(t *testing.T) {
	fset := token.NewFileSet()                       // 创建file set
	f, err := parser.ParseFile(fset, "", srcCode, 0) // 创建ast.File
	if err != nil {
		panic(err)
	}
	ast.Print(fset, f)
}

// 输出结果
/*
0  *ast.File {
1  .  Package: 2:1
2  .  Name: *ast.Ident {
3  .  .  NamePos: 2:9
4  .  .  Name: "hello"
5  .  }
6  .  Decls: []ast.Decl (len = 2) {
7  .  .  0: *ast.GenDecl {
8  .  .  .  TokPos: 4:1
9  .  .  .  Tok: import
10  .  .  .  Lparen: -
11  .  .  .  Specs: []ast.Spec (len = 1) {
12  .  .  .  .  0: *ast.ImportSpec {
13  .  .  .  .  .  Path: *ast.BasicLit {
14  .  .  .  .  .  .  ValuePos: 4:8
15  .  .  .  .  .  .  Kind: STRING
16  .  .  .  .  .  .  Value: "\"fmt\""
17  .  .  .  .  .  }
18  .  .  .  .  .  EndPos: -
19  .  .  .  .  }
20  .  .  .  }
21  .  .  .  Rparen: -
22  .  .  }
23  .  .  1: *ast.FuncDecl {
24  .  .  .  Name: *ast.Ident {
25  .  .  .  .  NamePos: 6:6
26  .  .  .  .  Name: "greet"
27  .  .  .  .  Obj: *ast.Object {
28  .  .  .  .  .  Kind: func
29  .  .  .  .  .  Name: "greet"
30  .  .  .  .  .  Decl: *(obj @ 23)
31  .  .  .  .  }
32  .  .  .  }
33  .  .  .  Type: *ast.FuncType {
34  .  .  .  .  Func: 6:1
35  .  .  .  .  Params: *ast.FieldList {
36  .  .  .  .  .  Opening: 6:11
37  .  .  .  .  .  Closing: 6:12
38  .  .  .  .  }
39  .  .  .  }
40  .  .  .  Body: *ast.BlockStmt {
41  .  .  .  .  Lbrace: 6:14
42  .  .  .  .  List: []ast.Stmt (len = 2) {
43  .  .  .  .  .  0: *ast.DeclStmt {
44  .  .  .  .  .  .  Decl: *ast.GenDecl {
45  .  .  .  .  .  .  .  TokPos: 7:5
46  .  .  .  .  .  .  .  Tok: var
47  .  .  .  .  .  .  .  Lparen: -
48  .  .  .  .  .  .  .  Specs: []ast.Spec (len = 1) {
49  .  .  .  .  .  .  .  .  0: *ast.ValueSpec {
50  .  .  .  .  .  .  .  .  .  Names: []*ast.Ident (len = 1) {
51  .  .  .  .  .  .  .  .  .  .  0: *ast.Ident {
52  .  .  .  .  .  .  .  .  .  .  .  NamePos: 7:9
53  .  .  .  .  .  .  .  .  .  .  .  Name: "msg"
54  .  .  .  .  .  .  .  .  .  .  .  Obj: *ast.Object {
55  .  .  .  .  .  .  .  .  .  .  .  .  Kind: var
56  .  .  .  .  .  .  .  .  .  .  .  .  Name: "msg"
57  .  .  .  .  .  .  .  .  .  .  .  .  Decl: *(obj @ 49)
58  .  .  .  .  .  .  .  .  .  .  .  .  Data: 0
59  .  .  .  .  .  .  .  .  .  .  .  }
60  .  .  .  .  .  .  .  .  .  .  }
61  .  .  .  .  .  .  .  .  .  }
62  .  .  .  .  .  .  .  .  .  Values: []ast.Expr (len = 1) {

63  .  .  .  .  .  .  .  .  .  .  0: *ast.BasicLit {
64  .  .  .  .  .  .  .  .  .  .  .  ValuePos: 7:15
65  .  .  .  .  .  .  .  .  .  .  .  Kind: STRING
66  .  .  .  .  .  .  .  .  .  .  .  Value: "\"Hello World!\""
67  .  .  .  .  .  .  .  .  .  .  }
68  .  .  .  .  .  .  .  .  .  }
69  .  .  .  .  .  .  .  .  }
70  .  .  .  .  .  .  .  }
71  .  .  .  .  .  .  .  Rparen: -
72  .  .  .  .  .  .  }
73  .  .  .  .  .  }
74  .  .  .  .  .  1: *ast.ExprStmt {
75  .  .  .  .  .  .  X: *ast.CallExpr {
76  .  .  .  .  .  .  .  Fun: *ast.SelectorExpr {
77  .  .  .  .  .  .  .  .  X: *ast.Ident {
78  .  .  .  .  .  .  .  .  .  NamePos: 8:5
79  .  .  .  .  .  .  .  .  .  Name: "fmt"
80  .  .  .  .  .  .  .  .  }
81  .  .  .  .  .  .  .  .  Sel: *ast.Ident {
82  .  .  .  .  .  .  .  .  .  NamePos: 8:9
83  .  .  .  .  .  .  .  .  .  Name: "Println"
84  .  .  .  .  .  .  .  .  }
85  .  .  .  .  .  .  .  }
86  .  .  .  .  .  .  .  Lparen: 8:16
87  .  .  .  .  .  .  .  Args: []ast.Expr (len = 1) {
88  .  .  .  .  .  .  .  .  0: *ast.Ident {
89  .  .  .  .  .  .  .  .  .  NamePos: 8:17
90  .  .  .  .  .  .  .  .  .  Name: "msg"
91  .  .  .  .  .  .  .  .  .  Obj: *(obj @ 54)
92  .  .  .  .  .  .  .  .  }
93  .  .  .  .  .  .  .  }
94  .  .  .  .  .  .  .  Ellipsis: -
95  .  .  .  .  .  .  .  Rparen: 8:20
96  .  .  .  .  .  .  }
97  .  .  .  .  .  }
98  .  .  .  .  }
99  .  .  .  .  Rbrace: 9:1
100  .  .  .  }
101  .  .  }
102  .  }
103  .  FileStart: 1:1
104  .  FileEnd: 9:3
105  .  Scope: *ast.Scope {
106  .  .  Objects: map[string]*ast.Object (len = 1) {
107  .  .  .  "greet": *(obj @ 27)
108  .  .  }
109  .  }
110  .  Imports: []*ast.ImportSpec (len = 1) {
111  .  .  0: *(obj @ 12)
112  .  }
113  .  Unresolved: []*ast.Ident (len = 1) {
114  .  .  0: *(obj @ 77)
115  .  }
116  .  GoVersion: ""
117  }

*/
