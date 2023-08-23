
CURRENT_DIR=$(shell pwd)
PACKAGE_DIRS=`go list -e ./... | egrep -v "binary_output_dir|.git|mocks"`
# Get current git commit number
LABEL=$(shell git log -1 --format=%h)
test: test-report 
test-package-dirs:
	go test $(PACKAGE_DIRS) -race -coverprofile=cover.out -covermode=atomic

test-report: test-package-dirs
	@echo 'Generating test coverage report...'
	 go tool cover -html=cover.out -o cover.html

	@echo 'Generating overall test coverage ...'
	go tool cover -func cover.out | grep total:

lint:
	@golangci-lint run -v

docker-build:
	@docker build -t qisstpay/nadra:$(LABEL) .

security-check:
	@echo 'Doing static security checks'
	gosec -include=G101,G102,G103,G104,G105, G106,G107,G108, G109, G110, G111, G112  ./... 
	gosec -include=G201,G202, G203, G204 ./...
	gosec -include=G301, G302, G303, G304, G305, G306, G307 ./...
	gosec -include=G401, G402, G403, G404 ./...
	gosec -include=G501, G503, G503, G504, G505 ./...
	gosec -include=G601 ./...

.PHONY: test test-package-dirs test-report