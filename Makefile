PATH := ${CURDIR}/bin:$(PATH)

.PHONY: build
build:
	mkdir -p bin/out
	go build \
		-o bin/out/cmd \
		./cmd/...

.PHONY: test
test:
	go test -v -race ./...

.PHONY: task
task: build
	./bin/out/cmd

.PHONY: codegen
codegen: bin/mockgen$(go_exe)
	go generate ./...

bin/mockgen$(go_exe): go.mod
	go build -o $@ go.uber.org/mock/mockgen
