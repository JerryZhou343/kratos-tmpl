# kratos-tmpl
基于哔哩哔哩 kratos 建立的项目模板脚手架模板代码风格有MVC 和 DDD 两种风格，传输方式都为GRPC方式。

## 帮助信息: 
```shell script
kratos-tmpl -h
Usage:
   [flags]

Flags:
  -g, --group name string     setting server name (default "demo")
  -h, --help                  help for this command
  -p, --project name string   setting project name (default "demo")
  -s, --site name string      setting site name (default "demo")
      --style string          setting project layout style. opt:ddd,mvc (default "ddd")
```

##  使用示例
```shell script
kratos-tmpl -s github.com -g jerryzhou -p demo
cd demo
make grpc
go mod tidy
make wire
make build
```

# 周边
[proto 编译辅助工具](https://github.com/JerryZhou343/prototool)
