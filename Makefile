
PKGS_NO_MOCKS=$(shell go list ./... | grep -v '/mocks')

test-cov:
	go test -cover $(PKGS_NO_MOCKS)

test-covout:
	go test -coverprofile=coverage.out $(PKGS_NO_MOCKS)

test-html:
	go tool cover -html=coverage.out

test-all: 
	make test-cov && make test-covout && make test-html