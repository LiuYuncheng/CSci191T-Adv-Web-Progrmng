package main
import "fmt"

func main(){
var x int =2
var y *int =&x

fmt.Println("x's value is ", x)
fmt.Println("x's address is ", &x)
fmt.Println("y's value is ", y)
fmt.Println("y points to ", *y)
}
