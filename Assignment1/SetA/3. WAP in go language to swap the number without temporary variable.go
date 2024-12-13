/*
3. WAP in go language to swap the number without temporary variable.
*/

package main
import ("fmt")

func main() {
    var a int = 72
    var b int = 120
    //add substact method
    fmt.Println("add substact method:")
    fmt.Printf("a=%d b=%d\n",a,b)
    fmt.Printf("%d=%d+%d\n",a+b,a,b)
    a=a+b
    fmt.Printf("%d=%d-%d\n",a-b,a,b)
    b=a-b
    fmt.Printf("%d=%d-%d\n",a-b,a,b)
    a=a-b
    fmt.Printf("a=%d b=%d\n",a,b)
    //multiply div method
    fmt.Println("multiply div method:")
    a=72
    b=120
    fmt.Printf("a=%d b=%d\n",a,b)
    fmt.Printf("%d=%d*%d\n",a*b,a,b)
    a=a*b
    fmt.Printf("%d=%d/%d\n",a/b,a,b)
    b=a/b
    fmt.Printf("%d=%d/%d\n",a/b,a,b)
    a=a/b
    fmt.Printf("a=%d b=%d\n",a,b)
}
