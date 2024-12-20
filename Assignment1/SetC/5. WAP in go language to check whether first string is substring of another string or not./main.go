/*
5. WAP in go language to check whether first string is substring of another string or not.
*/
package main

import (
	"fmt"
	"strings"
)

func main(){
	var s1,s2 string
	fmt.Println("enter string1:")
	fmt.Scan(&s1)
	fmt.Println("enter string2:")
	fmt.Scan(&s2)
	if(strings.Contains(s1,s2)){
		fmt.Println(s1," is a substring of ",s2)
	}else{
		fmt.Println(s1," is a not substring of ",s2)
	}
}
