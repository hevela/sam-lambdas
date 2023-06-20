.PHONY: build

build:
	sam build

build-HelloWorldFunction:
	cd hello-world; GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o hello-world main.go
	mv hello-world $(ARTIFACTS_DIR)
.PHONY: build-HelloWorldFunction