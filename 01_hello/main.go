/*
Every program must start with the package declaration. In Go language, packages are used to organize and reuse the code. In Go, there are two types of program available one is executable and another one is the library. The executable programs are those programs that we can run directly from the terminal and Libraries are the package of programs that we can reuse them in our program.
Here, the package main tells the compiler that the package must compile in the executable program rather than a shared library.
It is the starting point of the program and also contains the main function in it.
*/
package main

/*
Here, import keyword is used to import packages in your program and fmt package is used to implement formatted Input/Output with functions.
*/
import "fmt"

/*
func: It is used to create a function in Go language.
main: It is the main function in Go language, which doesn’t contain the parameter, doesn’t return anything, and call when you execute your program.
Println(): This method is present in fmt package and it is used to display “!… Hello World …!” string. Here, Println means Print line.
*/
func main(){
	fmt.Println("hello World")
}

/*
installation:
download go:
https://go.dev/dl/

install and add path:

verify:
C:\Users\Admin>go version
go version go1.23.4 windows/amd64

on vscode add:
The VS Code Go extension provides rich language support for the Go programming language.
You can install additional tools and update them by using "Go: Install/Update Tools".
run>Go: Install/Update Tools

terminal:
(initialize the go project)
PS C:\Users\Admin\Downloads\GO\mygolanglcm\01_hello> go mod init hello
go: creating new go.mod: module hello
go: to add module requirements and sums:
go: to add module requirements and sums:
        go mod tidy
(run project)
PS C:\Users\Admin\Downloads\GO\mygolanglcm\01_hello> go run .\main.go 
hello World
*/
