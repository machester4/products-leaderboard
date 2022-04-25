BINARY=service
build:
	go build ${LDFLAGS} -o ${BINARY} cmd/*.go

test:
	go test -race ./...

web: build
	./${BINARY} -E dev