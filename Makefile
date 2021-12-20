BINARY_NAME=file-renamer

build:
	go build -o bin/${BINARY_NAME} .

clean:
	go clean
	rm -f ./bin/${BINARY_NAME}

install:
	cp -f bin/${BINARY_NAME} /usr/local/bin/
