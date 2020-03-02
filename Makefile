.PHONY: all clean

OUTPUT=kratos-ddd

all: clean wire
	go build  -v -o ./bin/${OUTPUT} cmd/main.go

clean:
	rm -f ./bin/${OUTPUT}

wire:
	wire gen ./di