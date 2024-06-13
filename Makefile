.PHONY: build clean deploy

GO_BUILD := env GOARCH=arm64 GOOS=linux go build -ldflags="-s -w"

build:
	export GO111MODULE=on
	export CGO_ENABLED=1

	${GO_BUILD} -o bin/handler/bootstrap ./handler/main.go
	chmod +x bin/handler/bootstrap
	zip -j handler.zip bin/handler/bootstrap

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose --stage $(STAGE)
