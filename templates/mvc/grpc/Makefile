.PHONY: all clean wire

OUTPUT={{.Name}}

all: clean wire
	go build  -v -o ./bin/${OUTPUT} cmd/main.go

clean:
	rm -f ./bin/${OUTPUT}

wire:
	wire gen ./di