/*
4. WAP in go Language to print address of a variable.
*/

package main
import ("fmt")

func main() {
    var a int = 72
    fmt.Printf("var a=%d\n",a)
    fmt.Printf("address of var a=%p\n",&a)
}
