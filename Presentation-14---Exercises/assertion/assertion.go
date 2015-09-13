package main
import "fmt"

func main(){
var myInterface interface {}=3
fmt.Println(7+myInterface.(int))
}
