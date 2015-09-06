package main
import "fmt"

func main () {
	//this works
	mySentence := "1 2"
	var one, two int
	fmt.Sscan(mySentence, &one, &two)
	fmt.Println(one, two)

	//this does not. I thought interface{} was "any type"
	//so why is it telling me "can't accept []string as []interface{} ?
	mySentence = "this is a second attempt"
	var myWords []string
	num, err := fmt.Sscan(mySentence, myWords...)
	fmt.Println(num, err)
	fmt.Println(myWords)
}
