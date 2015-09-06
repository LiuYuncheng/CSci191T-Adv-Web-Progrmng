package main
import "fmt"

func main(){
fmt.Println("Please enter two numbers, the first larger than the second")
var dividend, divisor int
fmt.Scan(&dividend)
fmt.Scan(&divisor)
fmt.Println("The remainder is ", dividend & divisor)
}
