package main

import "fmt"

func main(){
	fmt.Println("Data types in Go Lang")
	fmt.Println("Integer Data Type")
	/*	
		Data type		Size
		int/uint		either 32 bits (4 bytes) or 64 bits (8 bytes)
		int16/uint16	16 bits (2 bytes)
		int8/uint8		8 bits (1 byte)
		int64/uint64	64 bits ( 8 bytes)
		int32/uint32	32 bits (4 bytes)
	*/
	var id,age int = 4,25
	fmt.Printf("id=%d age=%d\n",id,age)

	fmt.Println("Float Data Type")
	var salary float32 =50000.503882901
	fmt.Printf("Salary=%f\n",salary)
	/*
	Formatted Output:
	fmt.Printf()
		fmt.Printf allows you to specify how variables should be formatted, such as controlling the number of decimal places for floating-point numbers or aligning text.
		For example, you can print floating-point numbers with a specific number of digits after the decimal using %.2f (2 decimal places) or specify other formats like %d for integers or %s for strings.
	*/

	var scidata float64 =  50000.503882901
	fmt.Printf("scidata=%f\n",scidata)
	/*
		Data Type	Size
		float32		32 bits (4 bytes)
		float64		64 bits (8 bytes)
	*/
}