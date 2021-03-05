parser:
	make -C go/parser

build:
	go build ./go/...

unit-test:
	go test ./go/...
	
all: parser unit-test