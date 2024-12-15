/*
2. WAP in go language to accept two strings and compare them.
*/
package main

import (
	"fmt"
	"strings"
)

func main(){
	var str1,str2 string
	fmt.Println("Enter the First String:")
	fmt.Scanln(&str1)	
	fmt.Println("Enter the Second String:")
	fmt.Scanln(&str2)
	fmt.Printf("str1=%s\n",str1)
	fmt.Printf("str2=%s\n",str2)
	var flag bool = strings.Contains(str1,str2)
	//fmt.Printf("flag=%v",flag)
	if(flag==true){
		fmt.Printf("str1=%s and str2=%s are same.\n",str1,str2)
	}else{
		fmt.Printf("str1=%s and str2=%s are not same.\n",str1,str2)
	}
}
