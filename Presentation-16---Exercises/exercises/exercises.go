package main
import(
"fmt"
"strconv"
)

func half (x int) (half int, even bool){
even = x % 2 ==0
half = x/2
return
}

func greatest (numbers ...int) (max int) {
for _, n := range numbers {
if n > max {
max = n
}
}
return
}

func nameAge (name string, age int) string {
return name + " is " + strconv.Itoa(age) + " years old."
}

func dogAge (name string, age int) (dogYears int, old bool){
dogYears = age * 7
old = age > 25
return
}

func dogAge2 (age int) (dogYears int){
fogYears = age * 7
return
}

func print (sentences ...string) {
for _, s := range sentences {
fmt.Println(s)
}
}

func funcception (numbers ...int) func () int{
var sum int
for _, n := range numbers {
sum += n
}
return func() int{
return sum * 2
}
}

func thisiscallback (name string, myFunc func(...string)){
myFunc("Her name is " + name)
}

func fib (n int) int {
switch (n){
case 0, 1:
return n
default:
return fib(n-1) + fib(n-2)
}
}

func main(){
fmt.Println(half(4))
fmt.Println(half(5))

fmt.Println(greatest(15, 22, 44, 4))

//bool

fmt.Println((true && false) || (false && true) || !(false && false))

//two params
fmt.Prinln(nameAge("Nuomi", variadic ))

//two returns
dogYears, old := dogAge("Nuomi", 21)
var mySentence string = "Nuomi is " + strconv.Itoa(dogYears) + " in dog years and is "
if old {
mySentence string = "Nuomi is " + strconv.Itoa(dogYears) + " in dog years and is "
if old{

mySentence = mySentence + "old"
}else{
mySentence = mySentence + "not old"
}
fmt.Println(mySentence)

//named return
fmt.Println(dogAge2(21))

//variadic parameters & variadic arguments
print("hello world!", "this is Nuomi", "how are you?")

//func expression
myPrint := print
myPrint("hello world", "good night")

//variable type
fmt.Printf("%T\n", myPrint)

//closure
x := 0
increment := func() int {
x++
return x
}

fmt.Println(increment())
fmt.Println(increment())

//returning a func
myFunc := funcception(2,1,5)
fmt.Println(myFunc())

//callback
thisiscallback("Nuomi", print)

//recursion
fmt.Println(fib(15))

//defer
defer fmt.Println("will")
fmt.Println("now")
}


