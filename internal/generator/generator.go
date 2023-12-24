package generator

import (
	"embed"
	"fmt"
	"github.com/JerryZhou343/kratos-tmpl/internal/strs"
	"github.com/pkg/errors"
	"io/fs"
	"os"
	"path"
	"strings"
	"text/template"
)

const (
	StyleDDD = "ddd"
	StyleMVC = "mvc"
)

type Parameter struct {
	GoMod        string
	RelativePath string
}

type Generator struct {
	basePath  string
	style     string
	parameter *Parameter
	prjTmpl   *template.Template
}

func NewGenerator(parameter *Parameter, style, basePath string) *Generator {
	prjTmpl := template.New("ProjectTemplate")
	prjTmpl.Funcs(template.FuncMap{
		"CamelCase": func(s string) string {
			return strs.GoCamelCase(strings.ToLower(s))
		},
	})

	return &Generator{
		basePath:  basePath,
		style:     style,
		parameter: parameter,
		prjTmpl:   prjTmpl,
	}
}

func (g *Generator) GenProject(tmpl embed.FS) (err error) {
	isEmpty, err := g.IsDirEmpty()
	if err != nil {
		return err
	}
	if !isEmpty {
		return fmt.Errorf("dir:[%s] not empty", g.basePath)
	}
	//1. read template directory
	var (
		dirs   []fs.DirEntry
		parent string
	)

	if g.style == StyleDDD {
		parent = "tmpl/ddd"
	} else {
		parent = "tmpl/mvc"
	}

	if dirs, err = tmpl.ReadDir(parent); err != nil {
		return err
	}

	//3. foreach template file or subdirectory
	for _, dir := range dirs {
		if err = g.GenFile(parent, g.basePath, dir, tmpl); err != nil {
			return err
		}
	}

	return
}

func (g *Generator) GenFile(tmplParentPath, targetParentPath string, dir fs.DirEntry, tmpl embed.FS) (err error) {

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
			if err = g.GenFile(tmplPath, targetPath, tmpDir, tmpl); err != nil {
				return err
			}
		}

		return err
	}

	//2. process file
	return g.ParseFile(tmplPath, targetPath, tmpl)
}

func (g *Generator) ParseFile(tmplPath, targetPath string, tmpl embed.FS) (err error) {
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

	// render file
	render, err := g.prjTmpl.Parse(string(tmpContent))
	if err != nil {
		return errors.Wrapf(err, "parse template failed.")
	}

	if err = render.Execute(newFile, g.parameter); err != nil {
		return errors.Wrapf(err, "file %s", targetPath)
	}

	return nil
}

func (g *Generator) IsDirEmpty() (bool, error) {
	dir, err := os.ReadDir(g.basePath)
	if err != nil {
		return true, nil
	}

	if len(dir) == 0 {
		return true, nil
	}

	return false, err
}
