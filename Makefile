include Makefile.site

main: main.go
	@$(MAKE) build

api:
	sam local start-api --debug

run:
	go run main.go

fmt:
	go fmt *.go

build:
	env GOOS=linux GOARCH=amd64 go build -o main main.go

package:
	@zip main.zip.1 main
	@mv main.zip.1 main.zip

clean:
	rm -rf main.zip main
