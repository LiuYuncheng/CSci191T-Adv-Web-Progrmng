package main
import "fmt"
import "strconv"

func main(){
myNum := 44
fmt.Println("strconv: ", strconv.Itoa(myNum))

myString := "44"
num, _ := strconv.Atoi(myString)
fmt.Println(num + 1000)
}
