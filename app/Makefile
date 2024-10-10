.PHONY: build
build: clean
  GOOS=linux GOARCH=amd64 go build -o bucket-eventing main.go

.PHONY: buildmac
buildmac: clean
  GOOS=darwin GOARCH=arm64 go build -o bucket-eventing-arm64 main.go

.PHONY: run
run:
  go run main.go

.PHONY: clean
clean:
  rm -rf ./bucket-eventing ./bucket-eventing-arm64
