.PHONY: build-all

build-all:
	GOOS=linux go build -o build/trigger-func1/main ./trigger-func1
	GOOS=linux go build -o build/trigger-func2/main ./trigger-func2