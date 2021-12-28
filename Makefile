BINARY_NAME=file-renamer

build:
	go build -o bin/${BINARY_NAME} .

clean:
	go clean
	rm -f ./bin/${BINARY_NAME}

install: build
	cp -f bin/${BINARY_NAME} /usr/local/bin/

test:
	go test
