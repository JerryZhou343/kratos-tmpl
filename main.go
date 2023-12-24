package main

import (
	"embed"
	_ "embed"
	"github.com/JerryZhou343/kratos-tmpl/internal/generator"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"log"
	"os"
	"path/filepath"
)

//go:embed  tmpl/* tmpl/ddd/* tmpl/mvc/*
var tmpl embed.FS

var (
	wkDir              string
	parameter          = &generator.Parameter{}
	workspaceDirectory string
)

// PrintFlags logs the flags in the flagset
func PrintFlags(flags *pflag.FlagSet) {
	flags.VisitAll(func(flag *pflag.Flag) {
		log.Printf("FLAG: --%s=%q", flag.Name, flag.Value)
	})
}

func main() {
	var (
		err   error
		style string
	)
	wkDir, err = os.Getwd()
	if err != nil {
		log.Fatalf("get current work dir failed.err:[%v]", err)
	}

	var rootCmd = &cobra.Command{
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			PrintFlags(cmd.Flags())

			if workspaceDirectory == "" {
				workspaceDirectory = wkDir
			}

			basePath := filepath.Join(workspaceDirectory, parameter.RelativePath)

			err = os.MkdirAll(basePath, os.ModePerm)
			if err != nil {
				return err
			}

			g := generator.NewGenerator(parameter, style, basePath)
			err = g.GenProject(tmpl)

			return
		},
	}

	rootCmd.PersistentFlags().StringVarP(&parameter.GoMod, "gomod", "m", "github.com/demo/demo", "go mod")
	rootCmd.PersistentFlags().StringVarP(&parameter.RelativePath, "relativepath", "r", "", "the path relative to workspace folder. use to calculate path for go import")
	rootCmd.PersistentFlags().StringVarP(&workspaceDirectory, "workspace", "w", "", "the workspace path,default is current directory")

	rootCmd.PersistentFlags().StringVarP(&style, "style", "", "ddd", "setting project layout style. opt:ddd,mvc")
	if err = rootCmd.Execute(); err != nil {
		log.Printf("error:%v", err)
	}
}
