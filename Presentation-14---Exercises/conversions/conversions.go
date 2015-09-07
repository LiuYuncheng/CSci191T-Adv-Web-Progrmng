package main
import "fmt"

func main(){
//int to float64
myInt := 4
fmt.Println(int64(myInt))

//float64 to int
myFloat := 6.24
fmt.Println(int(myFloat))

//[]byte to string
myBytes := []byte{'n','i','g','h','t'}
fmt.Println(string(myBytes))

//string to [byte]
myString := "good night"
fmt.Println([]byte(myString))
}
