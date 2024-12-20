/*
4. WAP in go language to check whether accepted number is single digit or not.
*/
package main

import "fmt"

func main(){
	var num int
	fmt.Println("Enter the num:")
	fmt.Scan(&num)
	if(num>=-9 && num<=9){
		fmt.Println(num,"is a single digit")
	}else{
		fmt.Println(num,"is not single digit")
	}
}
