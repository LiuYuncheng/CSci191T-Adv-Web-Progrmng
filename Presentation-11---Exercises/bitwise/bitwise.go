package main
import "fmt"

func main(){
var x uint = 10
var y uint = 10
fmt.Println("x is ", x)
fmt.Println("y is ", y)
x = x << `
y = y >> 1
fmt.Println("x now ", x)
fmt.Println("y now ", y)
}
// 10: 000 1010
//10<<1(shift one bit left): 0001 0100 (2^4 + 2^2 = 20)
// 10>>1(shift one bit right: 0000 0101) (2^2 + 2^0 = 5)
