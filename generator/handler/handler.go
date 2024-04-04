package handler

import (
	"fmt"
	"go/ast"
	"reflect"
	"strings"
)

var moduleName = ""

type commandStruct struct {
	CommandType string
	RepoType    string
	fields      *[]commandField
}

type commandField struct {
	Name string
	Type string
}

func (item commandStruct) CommandTypeName() string {
	if strings.HasPrefix(item.CommandType, "*") {
		return item.CommandType[1:]
	}
	return item.CommandType
}

type TypeStr = string
type PackageStr = string

func getPackageAndType(expr ast.Expr) (PackageStr, TypeStr) {
	switch typ := expr.(type) {
	case *ast.Ident:
		return "", typ.Name
	case *ast.ArrayType:
		p, t := getPackageAndType(typ.Elt)
		return "[]" + t, p
	case *ast.StarExpr:
		_, t := getPackageAndType(typ.X)
		//return p, "*" + t
		return t, "*" + t
	case *ast.SelectorExpr:
		p, t := getPackageAndType(typ.X)
		p = t
		t += "." + typ.Sel.Name
		return p, t
	case *ast.MapType:
		p1, t1 := getPackageAndType(typ.Key)
		p2, t2 := getPackageAndType(typ.Key)
		var p string
		if p1 != "" && p2 != "" {
			p = p1 + "," + p2
		} else {
			p = p1 + p2
		}
		return p, fmt.Sprintf("map[%s]%s", t1, t2)
	// 处理更多类型...
	default:
		return "", fmt.Sprintf("%s", reflect.TypeOf(typ).String())
	}
}

func getModuleName() string {
	if len(moduleName) == 0 {
		moduleName = "github.com/AlphaFoxz/hot-deploy-go-example"
	}
	return moduleName
}
