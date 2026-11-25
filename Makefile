.PHONY: build test vet fmt clean

build:
	go build -o bin/daroga ./main.go

fmt:
	gofmt -w $$(find . -name '*.go' -not -path './.git/*')

vet:
	go vet ./...

test:
	go test -race ./...

clean:
	rm -rf bin
