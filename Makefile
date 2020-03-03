.PHONY: all clean

OUTPUT=kratos-tmpl

all: clean
	go build  -v -o ./bin/${OUTPUT} main.go new.go project.go

clean:
	rm -f ./bin/${OUTPUT}

