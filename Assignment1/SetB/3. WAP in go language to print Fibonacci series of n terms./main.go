/*
3. WAP in go language to print Fibonacci series of n terms.
*/

package main

import (
	"fmt"
)

func main(){
	fmt.Println("Fibonacci series")
	var n int = 10;
	n++
	//slice := make([]T, length, capacity)
	//var fibo []int(array)
	fibo:=make([]int,n,n)
	for i:=0;i<n;i++ {
		if(i==0||i==1) {
			fibo[i]=i
			fmt.Printf("%d ",fibo[i])
		}else {
			fibo[i]=fibo[i-2]+fibo[i-1]
			fmt.Printf("%d ",fibo[i])
		}
	}
}
