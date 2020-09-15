# kratos-tmpl
基于哔哩哔哩 kratos 建立的项目模板脚手架模板代码风格有MVC 和 DDD 两种风格，传输方式都为GRPC方式。

帮助信息: 
```shell script
kratos-tmpl -h
模板项目生成工具，
如果没有填输出目录参数，则生成到当前目录下，目录名为项目名。
go mod 前缀为必填参数,项目名为必填参数

Usage:
   [flags]

Flags:
  -h, --help           help for this command
  -m, --mod string     go mod 前缀,必填
  -n, --name string    项目名称,必填
  -p, --path string    项目生成目录 (default ".")
  -s, --style string   模板风格[ddd,mvc] (default "ddd")
      --version        version for this command

```

# 周边
[proto 编译辅助工具](https://github.com/JerryZhou343/prototool)
