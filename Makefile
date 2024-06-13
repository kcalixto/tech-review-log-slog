.PHONY: build clean deploy

GO_BUILD := env GOARCH=arm64 GOOS=linux go build -ldflags="-s -w"

build:
	export GO111MODULE=on
	export CGO_ENABLED=1

	${GO_BUILD} -o bin/nice-handler/bootstrap ./apis/nice/handler/main.go
	chmod +x bin/nice-handler/bootstrap
	zip -j bin/nice-handler.zip bin/nice-handler/bootstrap

	${GO_BUILD} -o bin/bad-handler/bootstrap ./apis/bad/handler/main.go
	chmod +x bin/bad-handler/bootstrap
	zip -j bin/bad-handler.zip bin/bad-handler/bootstrap

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose --stage $(STAGE)
