VERSION = $(shell cat ../VERSION)

build:
	@go build -i -ldflags "-X main.Version=${VERSION}" -o neo-local .