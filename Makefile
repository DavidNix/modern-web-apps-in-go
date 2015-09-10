.PHONY: setup test

default: test

setup:
	go get -t ./...
	go get github.com/smartystreets/goconvey
	go get github.com/mailgun/godebug

test:
	go test -timeout=60s ./...
