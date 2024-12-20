/*
1. WAP in go language to print addition of two number using function.
*/

package main
import (
  "fmt"
)

func add(x,y int)int {
  return x+y
}

func main(){
  fmt.Println("add=",add(2,5))
}
