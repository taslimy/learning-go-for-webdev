# Learning Go for Web Dev.
### Purpose
Gave up find a job doing html,css,js,node,react. So learning something new and maybe it pays off.

## Go Compiling
- *go run filename.go* compiles and immediately runs Go Programs.
- *go build filename.go* complaies a bunch of Go source files and it complies packages and dependencies. If you run *go build* it will complite the files in the current directory and will produce an executable file with the name of the current working directory.
- *go build -o app* will produce an ececutable file called app.
### Compiling for Specific platforms
- Compiling for Windows: `GOOS=windows GOARCH=amd64 go build -o winapp.exe`
- Compiling for Linux: `GOOS=linux GOARCH=amd64 go build -o linuxapp`
- Compiling for Mac: `GOOS=darwin GOARCH=amd64 go build -o macapp`

## Formatting in GO
- gofmt which comes from *golang formatter*
- **gofmt** is built into the language runtime and it formats Go code according to a set of stable, well-understood language rules.
- example: `gofmt -w filename.go` formats according to Go Style.