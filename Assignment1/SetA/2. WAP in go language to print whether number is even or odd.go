/*
2. WAP in go language to print whether number is even or odd.
*/

package main
import ("fmt")

func main() {
    var num int = 72
    if(num%2==0){
        fmt.Printf("num=%d is even num.\n",num)
    }else{
        fmt.Printf("num=%d is odd num.\n",num)
    }
    
}
