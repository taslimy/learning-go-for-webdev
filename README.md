# Learning Go for Web Dev.

### Purpose

Gave up find a job doing html,css,js,node,react. So learning something new and maybe it pays off.

## Go Compiling

- _go run filename.go_ compiles and immediately runs Go Programs.
- _go build filename.go_ complaies a bunch of Go source files and it complies packages and dependencies. If you run _go build_ it will complite the files in the current directory and will produce an executable file with the name of the current working directory.
- _go build -o app_ will produce an ececutable file called app.

### Compiling for Specific platforms

- Compiling for Windows: `GOOS=windows GOARCH=amd64 go build -o winapp.exe`
- Compiling for Linux: `GOOS=linux GOARCH=amd64 go build -o linuxapp`
- Compiling for Mac: `GOOS=darwin GOARCH=amd64 go build -o macapp`

## Formatting in GO

- gofmt which comes from _golang formatter_
- **gofmt** is built into the language runtime and it formats Go code according to a set of stable, well-understood language rules.
- example: `gofmt -w filename.go` formats according to Go Style.
- sets one standard to follow which I like.

## Variables

- A **variable** is a name for a memory location where a value of a specific type is stored.
- In Go a variable blongs and its created at runtime.
- A declated variable **must** be used or we get an error!
- \_ is the **blank indentifer** and mutes the compile-time error returned by unused variables.

### Declaring variables

```h
var x int = 7 break;
var hello string
hello = "Learning GO Homies!"
```

### Code Example

```h
package main

import "fmt"

func main() {
	// To write variable you need a keyword,  type, indentifer
	// I do not need to include the type as go will automatically know.
	var age int = 27
	fmt.Println("Age:", age)

	var name string = "Tas"
	// fmt.Println("My name is:", name)
	_ = name // If i get compile error for unused variables.

	// short declartion operator for writing a variable.
	name1 := "Tas"
	fmt.Println(name1)

	myAge := "old xd"
	fmt.Println(myAge)
}
```
