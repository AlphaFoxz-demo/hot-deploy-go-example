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

func HandleApi(sourceDir string) {
	sets := token.NewFileSet()
	srcCode, _ := fs.ReadFile(os.DirFS(sourceDir+"/declare"), "command.go")
	node, err := parser.ParseFile(sets, "", srcCode, 0)
	if err != nil {
		fmt.Printf("err = %s", err)
	}

	var initFuncValues []commandStruct
	importPackages := make(map[string]bool)

	//ast.Print(sets, node)
	for _, decl := range node.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}
		for _, spec := range genDecl.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			// 检查类型声明是否为结构体
			if structType, ok := typeSpec.Type.(*ast.StructType); ok {
				var fieldValues = []commandField{}
				//log.Printf("Found a struct declaration: %s\n", typeSpec.Name.Name)
				for _, field := range structType.Fields.List {
					for _, fieldName := range field.Names {
						// 将字段名转换为字符串并打印
						fieldNameStr := fieldName.Name
						if strings.HasSuffix(fieldNameStr, "_") {
							p, typeStr := getPackageAndType(field.Type)
							if len(p) > 0 {
								for _, v := range strings.Split(p, ",") {
									importPackages[v] = true
								}
							}
							initFuncValues = append(initFuncValues, commandStruct{
								CommandType: "*" + typeSpec.Name.Name,
								RepoType:    typeStr,
								fields:      &fieldValues,
							})

						} else {
							p, typeStr := getPackageAndType(field.Type)
							if len(p) > 0 {
								for _, v := range strings.Split(p, ",") {
									importPackages[v] = true
								}
							}
							fieldValues = append(fieldValues, commandField{Name: fieldNameStr, Type: typeStr})
						}
					}
				}
			}
		}
	}

	f := gg.NewGroup()
	f.AddPackage("api")
	i := f.NewImport()
	i.AddDot(getModuleName() + "/domain/declare")
	for k := range importPackages {
		i.AddPath(k)
	}
	a := f.NewStruct("api")
	repos := make(map[string]bool)
	for _, command := range initFuncValues {
		if !repos[command.RepoType] {
			a.AddField("repo_", command.RepoType)
			repos[command.RepoType] = true
		}
	}
	fun := f.NewFunction("New")
	for name := range repos {
		fun.AddParameter("repo_", name).AddBody(
			gg.String("    return api{"),
			gg.String("        repo_: repo_,"),
			gg.String("    }\n"),
		)
	}
	fun.AddResult("v", "api")

	for _, command := range initFuncValues {
		fun := f.NewFunction("New" + command.CommandTypeName())
		fieldsStr := []string{}
		for _, para := range *command.fields {
			if !strings.HasSuffix(para.Name, "_") {
				fieldsStr = append(fieldsStr, fmt.Sprintf("        %s: %s,", para.Name, para.Name))
			}
			fun.AddParameter(para.Name, para.Type)
		}
		fun.WithReceiver("inst_", "api")
		fun.AddResult("v", command.CommandTypeName())
		fun.AddBody(
			gg.String("    command := %s{", command.CommandTypeName()),
			gg.String(strings.Join(fieldsStr, "\n")),
			gg.String("    }"),
			gg.String("    return *command.Init(inst_.repo_)\n"),
		)
	}
	//logutils.LogDarkYellow(f.String())
	err = os.WriteFile(sourceDir+"/api/api_gen.go", []byte(f.String()), 0666)
	if err != nil {
		logutils.LogRed("api_gen.go generate failed!", err.Error())
		return
	}
	logutils.LogGreen("api_gen.go generated success.")
}
