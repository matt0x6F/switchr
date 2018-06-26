VERSION="0.0.2"

build:
	@export GOARCH=amd64

	@export GOOS=darwin
	@go build -o build/switchr-darwin-amd64-v${VERSION} -ldflags "-X github.com/mattouille/switchr/cmd.Version=${VERSION}"

	@export GOOS=linux
	@go build -o build/switchr-linux-amd64-v${VERSION} -ldflags "-X github.com/mattouille/switchr/cmd.Version=${VERSION}"