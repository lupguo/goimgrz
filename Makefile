# Make golang program
VERSION = 0.1.0
BINARY	= goimgrz
BINDIR	= ./bin/
BUILD	= `date +%FT%T%z`
PLATFORMS = darwin linux windows
ARCHITECTURES = 386 amd64
GOOS = `go env GOOS`

# Setup linker flags option for build that interoperate with variable names in src code
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"


# Compile only for the current platform
defualt: build
build:
	@for b in ${BINARY}; do \
		echo "build program into ./bin/$$b"; \
		go build -o ./bin/$$b ./cmd/$$b; \
	done
	@echo "------------------------" 
	@echo "see more using :" 
	@for b in ${BINARY}; do  \
		echo "  ./bin/$$b --help"; \
	done 

# Compile and install only for the current platform
install: 
	go get -v -u ./cmd/...

# Support for cross-platform compilation
build_all: 
	for GOOS in ${PLATFORMS}; do \
		for GOARCH in ${ARCHITECTURES}; do \
			echo "build $${GOOS}-$${GOARCH}"; \
			export GOOS=$${GOOS} GOARCH=$${GOARCH} \
				GO111MODULE=on GOPROXY=https://goproxy.io; \
		   	go build -v -o ${BINDIR}${BINARY}-$${GOOS}-$${GOARCH} ./cmd/...; \
			echo ${BINDIR}${BINARY}-$${GOOS}-$${GOARCH};\
		done \
	done

list:
	go list ./...
fmt:
	go fmt ./...
test:
	@echo "Nothing to test..."
	go test -v ./...
clean:
	@find `go env GOPATH`/bin `go env GOBIN` ./bin \
		-name '${BINARY}*' -exec echo 'rm {}' \; 
	@find `go env GOPATH`/bin `go env GOBIN` ./bin \
		-name '${BINARY}*' -delete

.PHONY: deafult build build_all list fmt install test clean
