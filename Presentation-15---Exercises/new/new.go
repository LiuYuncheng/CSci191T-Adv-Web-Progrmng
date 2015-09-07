package main
import "fmt"

func main(){
//int
var x *int = new (int)
fmt.Println(x)
fmt.Println(*x)
//true

//string
var y +string = new(string)
fmt.Println(y)
fmt.Println(*y)
//true

//bool
var z *bool = new(bool)
fmt.Println(z)
fmt.Println(*z)
//true

//make []int
slice := make([]int, 5, 5)
fmt.Println(slice)
//false

//map
myMap := make(map[int]string)
fmt.Println(myMap)
//false

}
