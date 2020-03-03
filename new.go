package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	common "github.com/bilibili/kratos/tool/pkg"
)

func runNew() (err error) {
	if p.path != "" {
		if p.path, err = filepath.Abs(p.path); err != nil {
			return
		}
		p.path = filepath.Join(p.path, p.Name)
	} else {
		pwd, _ := os.Getwd()
		p.path = filepath.Join(pwd, p.Name)
	}
	p.ModName = modPath(p.path)
	// creata a project
	if err := create(); err != nil {
		return err
	}
	fmt.Printf("Project: %s\n", p.Name)
	fmt.Printf("Style: %s\n", p.style)
	fmt.Printf("Trasnport: %s\n", p.transport)
	fmt.Printf("Directory: %s\n\n", p.path)
	fmt.Println("项目创建成功.")
	return nil
}

func modPath(p string) string {
	dir := filepath.Dir(p)
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			content, _ := ioutil.ReadFile(filepath.Join(dir, "go.mod"))
			mod := common.RegexpReplace(`module\s+(?P<name>[\S]+)`, string(content), "$name")
			name := strings.TrimPrefix(filepath.Dir(p), dir)
			name = strings.TrimPrefix(name, string(os.PathSeparator))
			if name == "" {
				return fmt.Sprintf("%s/", mod)
			}
			return fmt.Sprintf("%s/%s/", mod, name)
		}
		parent := filepath.Dir(dir)
		if dir == parent {
			return ""
		}
		dir = parent
	}
}
