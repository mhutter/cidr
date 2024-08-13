SOURCES := $(shell find . -name '*.go')

test: $(SOURCES)
	go test -v -race -cover ./...

lint: $(SOURCES)
	staticcheck ./...
