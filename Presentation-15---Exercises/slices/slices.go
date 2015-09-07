package main
import "fmt"

func main(){
//slice of ints
mySlice := []int{2,3,4}
fmt.Println(mySlice)

//slice of strings
mySecondSlice := []string{"good", "night"}
fmt.Println(mySecondSlice)

//slice of ints using make
var myThirdSlice []int = make([]int, 5, 10)
myThirdSlice[0] = 1
myThirdSlice[1] = 2
myThirdSlice[2] = 3
myThirdSlice[3] = 4
myThirdSlice[4] = 5
fmt.Println(myThirdSlice)
fmt.Println(len(myThirdSlice))
fmt.Println(cap(myThirdSlice))

//append a slice to a slice
myNewSlice := []int{3,4,5}
myOtherSlice := []int{4,5,6}
fmt.Println(append(myNewSlice, myOtherSlice...))

//delete an element from a slice
myNewSlice = append(myNewSlice[:3], myNewSlice[4:]...)
fmt.Println(myNewSlice)

//index out of range error
myLastSlice := make([]int, 2, 3)
myLastSlice[2]=7
}

