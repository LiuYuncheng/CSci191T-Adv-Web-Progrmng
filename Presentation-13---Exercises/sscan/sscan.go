package main
import "fmt"

func main(){
	mySentence := "2 4"
	var two, four int
	fmt.Sscan(mySentence, &two, &four)
	fmt.Println(two, four)
	//{} can be any type but [] string can't accept as []interface{}
	mySentence = "a second attempt"
	var myWords []string
	num, err := fmt.Sscan(mySentence, myWords...)
	fmt.Println(num, err)
	fmt.Println(myWords)
}
	
	
