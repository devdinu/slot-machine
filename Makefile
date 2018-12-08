.PHONY: build-all

LOADENV=$(shell cat envfile | xargs)
build-all: install-deps build imports test

install-deps:
	go get -u -v github.com/golang/dep/cmd/dep
	dep version
	dep ensure -v

build:
	cd cmd/server/ && go build ./... && cd ../../

lint:
	golint $(go list ./...) | { grep -Ev 'exported.*should have comment.*' || true; }

vet:
	go vet ./...

imports:
	goimports -l .  | { grep -v vendor || true; }

test:
	go test -race ./...

run: build
	./cmd/server/server
