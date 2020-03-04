package main

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

func main() {
	var (
		rootCmd = &cobra.Command{
			Version: "v0.0.0",
			Use:     "",
			Short:   "模板项目生成工具",
			Long: `模板项目生成工具，
如果没有填输出目录参数，则生成到当前目录下，目录名为项目名。
go mod 前缀为必填参数,项目名为必填参数`,
			Args: cobra.NoArgs,
			RunE: func(cmd *cobra.Command, args []string) error {
				//flag 检查
				if p.Name == "" || p.ModName == "" {
					return errors.New("缺少必填标志(flag)")
				}

				//todo:检测服务名
				p.SvcName = strings.ToUpper(string(p.Name[0]))+string(p.Name[1:])
				return runNew()
			},
		}
	)
	//bind flag
	rootCmd.Flags().StringVarP(&p.path, "path", "p", ".", "项目生成目录")
	rootCmd.Flags().StringVarP(&p.ModName, "mod", "m", "", "go mod 前缀,必填")
	rootCmd.Flags().StringVarP(&p.Name, "name", "n", "", "项目名称,必填")
	rootCmd.Flags().StringVarP(&p.transport, "transport", "t", transport_grpc,
		fmt.Sprintf("传输方式[%s,%s,%s]", transport_all, transport_grpc, transport_http))
	rootCmd.Flags().StringVarP(&p.style, "style", "s", style_ddd,
		fmt.Sprintf("模板风格[%s,%s]", style_ddd, style_mvc))
	rootCmd.Execute()
}
