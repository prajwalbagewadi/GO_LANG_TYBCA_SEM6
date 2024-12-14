/*
2. WAP in go language to print PASCALS triangle.
*/
package main

import (
	"fmt"
)

func main(){
	fmt.Println("Pascals Triangle.")
	var n int = 5
	var arr[5][5] int
	for i:=0;i<n;i++ {
		for j:=0;j<=i;j++ {
			if(j==0||j==i){
				arr[i][j]=1;
			}else{
				arr[i][j]=arr[i-1][j-1]+arr[i-1][j]
			}
		}
	}
	for i:=0;i<n;i++ {
		for j:=0;j<n;j++ {
			fmt.Printf("%d\t",arr[i][j])
		}
		fmt.Printf("\n")
	}
}
