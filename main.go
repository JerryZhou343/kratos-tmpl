package main

import (
	"embed"
	_ "embed"
	"github.com/JerryZhou343/kratos-tmpl/internal/strs"
	"io/fs"
	"log"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type Parameter struct {
	Site    string
	Group   string
	Project string
}

//go:embed  tmpl/* tmpl/ddd/* tmpl/mvc/*
var tmpl embed.FS

var (
	wkDir     string
	parameter Parameter
)

const (
	StyleDDD = "ddd"
	StyleMVC = "mvc"
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
			err = GenProject(&parameter, style)
			return
		},
	}

	rootCmd.PersistentFlags().StringVarP(&parameter.Site, "site name", "s", "demo", "setting site name")
	rootCmd.PersistentFlags().StringVarP(&parameter.Group, "group name", "g", "demo", "setting server name")
	rootCmd.PersistentFlags().StringVarP(&parameter.Project, "project name", "p", "demo", "setting project name")
	rootCmd.PersistentFlags().StringVarP(&style, "style", "", "ddd", "setting project layout style. opt:ddd,mvc")
	if err = rootCmd.Execute(); err != nil {
		log.Printf("error:%v", err)
	}
}

func GenProject(parameter *Parameter, style string) (err error) {
	//1. read template directory
	var (
		dirs   []fs.DirEntry
		parent string
	)

	if style == StyleDDD {
		parent = "tmpl/ddd"
	} else {
		parent = "tmpl/mvc"
	}

	if dirs, err = tmpl.ReadDir(parent); err != nil {
		return err
	}

	//2.create server root directory
	rootDirectory := path.Join(wkDir, parameter.Project)
	if err = os.MkdirAll(rootDirectory, os.ModePerm); err != nil {
		return errors.Wrap(err, "create server root directory failed.")
	}

	//3. foreach template file or subdirectory
	for _, dir := range dirs {
		if err = GenFile(parent, rootDirectory, dir); err != nil {
			return err
		}
	}

	return
}

func GenFile(tmplParentPath, targetParentPath string, dir fs.DirEntry) (err error) {

	tmplPath := path.Join(tmplParentPath, dir.Name())
	targetPath := path.Join(targetParentPath, dir.Name())

	//1. process directory
	if dir.IsDir() {
		if err = os.MkdirAll(targetPath, os.ModePerm); err != nil {
			return err
		}

		subDirs, err := tmpl.ReadDir(tmplPath)
		if err != nil {
			return errors.Wrapf(err, "read sub dir:", tmplPath)
		}

		for _, tmpDir := range subDirs {
			if err = GenFile(tmplPath, targetPath, tmpDir); err != nil {
				return err
			}
		}

		return err
	}

	//2. process file
	return ParseFile(tmplPath, targetPath)
}

func ParseFile(tmplPath, targetPath string) (err error) {
	tmpContent, err := tmpl.ReadFile(tmplPath)
	if err != nil {
		return errors.Wrapf(err, "read file:%s", tmplPath)
	}

	if strings.Contains(targetPath, ".tmpl") {
		targetPath = strings.Replace(targetPath, ".tmpl", "", 1)
	}

	newFile, err := os.OpenFile(targetPath, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return errors.Wrapf(err, "create file:%s", targetPath)
	}

	//3. render file
	tmpl := template.New("template")
	tmpl.Funcs(template.FuncMap{
		//"ToLower": strings.ToLower,
		//"ToUpper": strings.ToUpper,
		//"ToTitle": strings.Title,
		"CamelCase": func(s string) string {
			return strs.GoCamelCase(strings.ToLower(s))
		},
	})
	renderConent, err := tmpl.Parse(string(tmpContent))
	if err != nil {
		return errors.Wrapf(err, "parse template failed.")
	}

	if err = renderConent.Execute(newFile, parameter); err != nil {
		return errors.Wrapf(err, "file %s", targetPath)
	}

	return nil
}
