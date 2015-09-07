package main
import "fmt"

func main(){
myMap := map[string]string{
"yuncheng":"liu",
"Todd":"McLeod",
"Mark":"Zuckerberg",
}

myMap["Navmit"] = "Danjeer"
myMap["Navmit"] = "Something Else"
delete(myMap, "Navmit")
for key, val := range myMap{
fmt.Println(key, " - ", val)
}

fmt.Println(len(myMap))
if val, ok := myMap["Mark"]; ok{
fmt.Println(val)
}
}
