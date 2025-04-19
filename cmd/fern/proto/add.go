package proto

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

type TemplateData struct {
	PackageName        string
	ServiceName        string
	AdditionalMessages []string
}

var (
	addCmd = &cobra.Command{
		Use:   "add [proto file path]",
		Short: "Generate a new proto file template",
		Long:  `Generate a new proto file with basic template at the specified path.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			filePath := args[0]
			// 准备模板数据
			serviceName := strings.TrimSuffix(filepath.Base(filePath), ".proto")
			data := TemplateData{
				PackageName: serviceName,
				ServiceName: serviceName,
			}
			// 模板编译
			tmpl, err := template.ParseFiles("templates/proto.tpl")
			if err != nil {
				log.Fatal(err)
			}
			// 写入文件
			var fullpath string
			if filepath.Dir(filePath) == "." {
				fullpath = filepath.Join("api", filePath)
			} else {
				fullpath = filePath
			}
			file, err := os.Create(fullpath)
			if err != nil {
				log.Fatal(err)

			}
			defer file.Close()

			if err := tmpl.Execute(file, data); err != nil {
				log.Fatal(err)
			}
			os.Exit(0)
		},
	}
)

func init() {
	ProtoCmd.AddCommand(addCmd)
}
