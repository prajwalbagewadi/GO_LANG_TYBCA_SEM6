/*
3. WAP in go language to accept user choice and print answer of using arithmetical operators.
*/

package main

import (
	"fmt"
)

func main(){
    var num1,num2 int
    var opr string
    fmt.Println("Calculator:")
    fmt.Println("enter first number:")
    fmt.Scan(&num1)
    fmt.Println("enter Second number:")
    fmt.Scan(&num2)
    fmt.Println("Select operation:(+,-,*,/,%):")
    fmt.Scan(&opr)
    switch opr{
        case "+":
			fmt.Println("sum=",num1+num2)
		case "-":
			fmt.Println("sub=",num1-num2)
		case "*":
			fmt.Println("mul=",num1*num2)
		case "/":
			fmt.Println("div=",num1/num2)
		case "%":
			fmt.Println("mod=",num1%num2)
		default :
			fmt.Println("invalid operator.")
    }
}
