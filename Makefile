.PHONY: all clean build deploy test

all: test

build/main: cmd/main.go $(wildcard pkg/**/*)
	GO111MODULE=on GOOS=linux GOARCH=amd64 go build -o build/main cmd/main.go

build/main.zip: build/main
	zip -FS -j build/main.zip build/main

build: build/main

clean: 
	rm -rf build/

deploy: build/main.zip
	aws lambda update-function-code \
		--function-name=alexa-ov \
		--zip-file=fileb://build/main.zip

test: build/main
	sam local invoke AlexaOVFunction -e test/start-session.json