package main
import "fmt"

func myFunc (x ...int){
fmt.Println(x)
}
func main(){
myFunc(1,2,3)
var x =[] int {1,2,3}
myFunc (x...)
}
