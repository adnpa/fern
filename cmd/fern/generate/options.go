package generate

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"html/template"
	"log"
	"os"
	"strings"
)

func GenerateOption() {
	// 解析源文件
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "options.go", nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	// 创建 types.Config 用于类型检查
	conf := types.Config{
		Importer: nil, // 使用默认导入器
	}

	// 创建空的文件集用于类型检查
	pkg, err := conf.Check("", fset, []*ast.File{f}, nil)
	if err != nil {
		log.Fatal(err)
	}

	// 收集字段信息
	var fields []struct {
		Name      string
		ParamName string
		Type      string
	}

	ast.Inspect(f, func(n ast.Node) bool {
		// 查找类型定义
		typeSpec, ok := n.(*ast.TypeSpec)
		if !ok || typeSpec.Name.Name != "options" {
			return true
		}

		// 获取结构体类型信息
		obj := pkg.Scope().Lookup(typeSpec.Name.Name)
		if obj == nil {
			return true
		}

		structType, ok := obj.Type().Underlying().(*types.Struct)
		if !ok {
			return true
		}

		// 遍历结构体字段
		for i := 0; i < structType.NumFields(); i++ {
			field := structType.Field(i)
			fields = append(fields, struct {
				Name      string
				ParamName string
				Type      string
			}{
				Name:      field.Name(),
				ParamName: strings.ToLower(field.Name()),
				Type:      field.Type().String(),
			})
		}
		return false
	})

	// 生成代码
	t := template.Must(template.New("with").Parse(OPTION_TMPL))
	err = t.Execute(os.Stdout, fields)
	if err != nil {
		log.Fatal(err)
	}
}
