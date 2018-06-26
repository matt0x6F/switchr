VERSION="0.0.2"

build:
	@go build -ldflags "-X github.com/mattouille/switchr/cmd.Version=${VERSION}"