package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/gobuffalo/packr/v2"
)

const (
	style_ddd = "ddd"
	style_mvc = "mvc"

	transport_all  = "all"
	transport_grpc = "grpc"
	transport_http = "http"
)

// project project config
type project struct {
	// project name
	Name string
	// mod prefix
	ModName string
	// project dir
	path      string
	style     string
	transport string
	SvcName string
}

var p project

//go:generate packr2
func create() (err error) {
	var box *packr.Box
	if p.style == style_mvc {
		box = packr.New("mvc.all", "./templates/mvc/all")
		if p.transport == transport_http {
			box = packr.New("mvc.http", "./templates/mvc/http")
		} else if p.transport == transport_grpc {
			box = packr.New("mvc.grpc", "./templates/mvc/grpc")
		}
	} else {
		box = packr.New("ddd.all", "./templates/ddd/all")
		if p.transport == transport_http {
			box = packr.New("ddd.http", "./templates/ddd/http")
		} else if p.transport == transport_grpc {
			box = packr.New("ddd.grpc", "./templates/ddd/grpc")
		}
	}
	if err = os.MkdirAll(p.path, 0755); err != nil {
		return
	}
	for _, name := range box.List() {
		tmpl, _ := box.FindString(name)
		i := strings.LastIndex(name, string(os.PathSeparator))
		if i > 0 {
			dir := name[:i]
			if err = os.MkdirAll(filepath.Join(p.path, dir), 0755); err != nil {
				return
			}
		}
		if strings.HasSuffix(name, ".tmpl") {
			name = strings.TrimSuffix(name, ".tmpl")
		}
		if err = write(filepath.Join(p.path, name), tmpl); err != nil {
			return
		}
	}

	return
}

func generate(path string) error {
	cmd := exec.Command("go", "generate", "-x", path)
	cmd.Dir = p.path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func write(path, tpl string) (err error) {
	data, err := parse(tpl)
	if err != nil {
		return
	}
	return ioutil.WriteFile(path, data, 0644)
}

func parse(s string) ([]byte, error) {
	t, err := template.New("").Parse(s)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	if err = t.Execute(&buf, p); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
