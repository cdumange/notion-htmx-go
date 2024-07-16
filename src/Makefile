# the name of the binary when built
BINARY_NAME=notion-htmx-go

.PHONY: all clean build-debug dev templates server

# remove any binaries that are built
clean:
	rm -f ./bin/$(BINARY_NAME)*

build-debug: clean
	CGO_ENABLED=0 go build -gcflags=all="-N -l" -o bin/$(BINARY_NAME)-debug main.go\

dev: templates server
templates:
	templ generate --watch &
server:
	air &