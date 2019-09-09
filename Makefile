# Make Golang Program
PACKAGE	= github.com/tkstorm/goimgrz
BINARY	= `ls ./cmd/`
BUILD	= `date +%FT%T%z`

build:
	@echo "build on time: ${BUILD}"
	@for b in ${BINARY}; do \
		echo "build program '$$b'"; \
		go build -o ./bin/$$b ./cmd/$$b ; \
	done
	@echo "------------------------"
	@echo "see more using :"
	@for b in ${BINARY}; do \
		echo "  ./bin/$$b --help"; \
	done
list:
	go list ./...
fmt:
	go fmt ./...
install: 
	go get -v -u ./cmd/...
test:
	@echo "Nothing to test..."
clean:
	@for b in ${BINARY}; do \
		find "`go env GOPATH`/bin" `go env GOBIN` ./bin \
			\( -name $$b -or -name $${b}".exe" \) \
			-exec rm -f {} \; ; \
	done
