/*
2. WAP in go language to print recursive sum of digits of given number.
*/

package main
import (
  "fmt"
  )
 
/*
  fmt.Println(123%10)  //3
  fmt.Println(123/10)  //12
  fmt.Println(12%10)   //2
  fmt.Println(1/10)    //0
  fmt.Println(1%10)    //1
  
  fmt.Println(456%10) //6
  fmt.Println(456/10) //45
  fmt.Println(45%10) //5
  fmt.Println(45/10) //4
  fmt.Println(4%10) //4
  
  fmt.Println(456%10) //6
  fmt.Println(456/10) //45
  fmt.Println(45%10) //5
  fmt.Println(45/10) //4
  fmt.Println(4%10) //4
  fmt.Println(4/10) //0
*/

func recsum(num int)int{
    fmt.Println("num=",num)
    if(num==0){
      return 0
    }
   
    return num%10+recsum(num/10)
}

func main(){
  var num int
  fmt.Println("Enter num:")
  fmt.Scan(&num) 
  sum:=recsum(num)
  fmt.Println("sum=",sum)
}
