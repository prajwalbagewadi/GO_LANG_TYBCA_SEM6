/*
1. WAP in go language to concatenate two strings using pointers.
*/

package main

import (
	"fmt"
)

func main(){
	fmt.Printf("String Concat using Pointers:\n")
	var str1 string = "Prajwal"
	var str2 string = "Bagewadi"
	fmt.Printf("str1=%s\n",str1)
	fmt.Printf("str2=%s\n",str2)
	//str var addresses are stored in ptr and concated using ptr
	var s1ptr *string=&str1
	fmt.Printf("Address s1ptr=%p\n",s1ptr)
	var s2ptr *string=&str2
	fmt.Printf("Address s2ptr=%p\n",s2ptr)
	var res string = *s1ptr + *s2ptr
	fmt.Printf("res=%s",res)
}
