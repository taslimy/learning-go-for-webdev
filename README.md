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

## Comments

### How to write comments

```
// Single comment

/*

Block comment like how you do it in CSS.

*/
```

## Naming Conventions - Important for readability, maintainability

- names start with a **letter** or an underscore (\_).
- _case matters:_ quickSort and QuickSort are differnt variables.
- Go [keyword](https://go101.org/article/keywords-and-identifiers.html) Go has only 25 keywords that cannot be used as a name for a variable.
- Sitck to camelCasing! just like in JS.
- Go does not provide automatic setup for Gettings and Setters.
- best pratice is to use the abbreviated letters of words:

```
var mv int = 100//mv -> max value
var max_value int = 100 // NOT OK
```

- use fewer letters in smaller scopes and complete word in larger scopes

```h
// DO NOT USE UNDERSCORES
var packetsRecieved int // NOT OK, name to long.
var n int = 500 // OK ->
var taskDone bool // OK in larger scopes.

var Max = 100 // Can be used globally by other packages.
var max = 100 // local to your current package.

// camelCasing = Recommended.
maxValues := 100 // Recommended.
max_values := // NOT Recommended.

writeToDB := true // ok, idiomatic.
writeToDb := false // not ok.
```

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

### Code Example - Differnt ways to use variables.

```h
ppackage main

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

	// Declaring multiple variables.
	car, cost := "BMW", 60000
	fmt.Println(car, cost)
	car, year := "Toyota", 2018
	_ = year

	opened := false
	opened, file := true, "a.txt"

	_, _ = opened, file

	// Clearer way to write multiple variables - better readability
	var (
		salary    float64
		firstName string
		gender    bool
	)

	fmt.Println(salary, firstName, gender)

	// multiple same type
	var a, b, c int
	fmt.Println(a, b, c)

	// multiple assignment
	var d, e int
	d, e = 5, 8

	d, e = e, d // Swapping variables
	println(d, e)

	sum := 6 + 2.3
	fmt.Println(sum)
}
```

## Zero Values (default values I guess)

### Go Zero values:

- numeric types: 0
- bool type: false
- string type: "" (empty string)
- poiner type: nil
