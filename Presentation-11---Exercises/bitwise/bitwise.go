package main
import "fmt"

func main(){
var x uint = 1
var y uint = 2
fmt.Println("x is ", x)
fmt.Println("y is ", y)
x = 1 << x
y = y >> 1
fmt.Println("x now ", x)
fmt.Println("y now ", y)
}
