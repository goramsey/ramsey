MAKEFLAGS = -s

cypher.go: cypher.y
	go run golang.org/x/tools/cmd/goyacc -o cypher.go cypher.y
	gofmt -w cypher.go
