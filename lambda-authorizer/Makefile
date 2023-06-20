.PHONY: build

build:
	sam build

build-AuthorizerFunction:
	cd authorizer; GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o authorizer main.go
	mv authorizer $(ARTIFACTS_DIR)
.PHONY: build-AuthorizerFunction
