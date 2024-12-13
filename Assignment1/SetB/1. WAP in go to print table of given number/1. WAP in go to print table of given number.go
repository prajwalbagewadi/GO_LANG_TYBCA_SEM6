/*
1. WAP in go to print table of given number.
*/

package main

import (
	"fmt"
)

func main(){
    var tab int
    fmt.Println("Enter a num to get table num:")
    fmt.Scanln(&tab)
    for i:=1;i<=10;i++{
        fmt.Printf("%d*%d=%d\n",tab,i,tab*i)
    }
}
