build:
	@go build -o bin/curly
run: build
	@./bin/curly
build-linux:
	@GOOS=linux GOARCH=amd64 go build -o bin/curly
build-win:
	@GOOS=windows GOARCH=amd64 go build -o bin/curly.exe
build-mac:
	@GOOS=darwin GOARCH=amd64 go build -o bin/curly
