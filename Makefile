.PHONY: all clean

OUTPUT=kratos-tmpl

all: clean
	go build  -v -o ./bin/${OUTPUT} main.go

clean:
	rm -f ./bin/${OUTPUT}

