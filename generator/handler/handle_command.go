package handler

import (
	"github.com/AlphaFoxz/hot-deploy-go-example/generator/utils/logutils"
	"github.com/Xuanwo/gg"

	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"strings"
)

func HandleCommand(sourceDir string) {
	sets := token.NewFileSet()
	srcCode, _ := fs.ReadFile(os.DirFS(sourceDir+"/declare"), "command.go")
	node, err := parser.ParseFile(sets, "", srcCode, 0)
	if err != nil {
		fmt.Printf("err = %s", err)
	}

	var initFuncValues []commandStruct

	//ast.Print(sets, node)
	for _, decl := range node.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}
	Struct:
		for _, spec := range genDecl.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			// 检查类型声明是否为结构体
			if structType, ok := typeSpec.Type.(*ast.StructType); ok {
				//log.Printf("Found a struct declaration: %s\n", typeSpec.Name.Name)
				for _, field := range structType.Fields.List {
					for _, fieldName := range field.Names {
						// 将字段名转换为字符串并打印
						fieldNameStr := fieldName.Name
						if strings.HasSuffix(fieldNameStr, "_") {
							_, typeStr := getPackageAndType(field.Type)
							initFuncValues = append(initFuncValues, commandStruct{CommandType: "*" + typeSpec.Name.Name, RepoType: typeStr})
							continue Struct
						}
					}
				}
			}
		}
	}

	f := gg.NewGroup()
	f.AddPackage("declare")
	for _, fun := range initFuncValues {
		f.NewFunction("Init").WithReceiver("command", fun.CommandType).AddParameter("repo_", fun.RepoType).AddResult("v", fun.CommandType).AddBody(
			gg.String("    command.repo_ = repo_"),
			gg.String("    return command\n"),
		)
	}
	//logutils.LogDarkYellow(f.String())
	err = os.WriteFile(sourceDir+"/declare/command_gen.go", []byte(f.String()), 0666)
	if err != nil {
		logutils.LogRed("command_gen.go generate failed!", err.Error())
		return
	}
	logutils.LogGreen("command_gen.go generated success.")
}
