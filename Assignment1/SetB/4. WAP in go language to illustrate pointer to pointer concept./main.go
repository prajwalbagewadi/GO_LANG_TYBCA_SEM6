/*
4. WAP in go language to illustrate pointer to pointer concept.
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println("pointer to pointer:")
	var num int = 10
	fmt.Printf("num=%d\n",num)
	fmt.Printf("num address=%p\n",&num)
	var ptr_to_num *int = &num
	fmt.Printf("ptr_to_num=%p\n",ptr_to_num)
	fmt.Printf("val stored at the addr stored in ptr_to_num=%d\n",*ptr_to_num)
	var ptr_to_ptr **int = &ptr_to_num
	fmt.Printf("ptr_to_ptr=%p\n",ptr_to_num)
	fmt.Printf("val stored at the addr stored in ptr_to_ptr=%d\n",**ptr_to_ptr)
}
