package main
import "fmt"

type student struct{
name string
age int
}

func main(){
alicia := student{name: "Alicia", age:18}
angelina := student{name:"Angelina"}

fmt.Println(alicia.name, alicia.age)
fmt.Println(angelina.name)

angelina.age=21
fmt.Println(angelina.age)
}
