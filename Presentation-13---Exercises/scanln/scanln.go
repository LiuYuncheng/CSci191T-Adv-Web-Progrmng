ppackage main
import "fmt"

func main(){
var num int
fmt.Print("Enter a number: ")
fmt.Scanln (&num)
newNum := num * 5
fmt.Println(num, "x 5 =", newNum)
}
