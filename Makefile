.PHONY: build
build:
	go build -ldflags "-X main.version=$(VERSION)" -o dist/local/act main.go