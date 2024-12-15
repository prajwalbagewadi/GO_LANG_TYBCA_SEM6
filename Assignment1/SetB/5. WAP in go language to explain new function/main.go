/*
5. WAP in go language to explain new function
*/
package main

import (
	"fmt"
	"unsafe"
	/* the unsafe package provides functions to find the size of variables.
	We can use the Sizeof function from the unsafe package to find the size of a struct*/)

func main(){
	new_ptr:=new(int)
	*new_ptr=10
	fmt.Printf("memory loc allocated to new_ptr=%p\n",new_ptr)
	fmt.Printf("memeory loc size of new_ptr=%d Bytes for 64bits sys.\n",unsafe.Sizeof(new_ptr))
	fmt.Printf("val stored at that mem loc new_ptr=%d\n",*new_ptr)
}
