build:
	go build -o bin/aaf cmd/aaf.go

all: build

clean:
	rm -f bin/**
