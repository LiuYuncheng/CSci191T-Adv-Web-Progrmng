package main
import(
"fmt"
"strings"
"math"
"io"
"io/ioutil"
"bufio"
"os"
"encoding/csv"
"strconv"
"crypto/md5"
"path/filepath"
"html/template"
"encoding/hex"
)

/* creat a program that returns a map of the counts of each "word" in a string. (strings.Fields)
WordCount("test test") → map[string]int{"test": 2}
*/

func WordCount (s string) map[string]int{
words := strings.Fields (s)
result := make(map[string]int)

for _, w := range words {
result[w]++   //if key is not present, it returns the value's zero value, which is 0
}

return result
}


/*
create a program that computes the average of a list of numbers, but removes the largest and smallest values.
centeredAverage([]float64{1,3,5,8,99}) → 8
*/


func CenteredAverage (numbers []float64) float64{
if len(numbers) <= 2 {
return 0.0
}
min := numbers[0]
max := numbers[1]

for _, val := range numbers {
min = math.Min(min, val)
max = math.Max(max, val)
}

sum := 0.0
for _, val := range numbers {
if val != max && val != min{
sum += val
}
}

return sum / float64(len(numbers) - 2)
}

/*
Write a program that can swap two integers
x := 4
y := 3
swap(&x, &y)
should give us x=3 and y=4
*/


func swap (x, y *int){
temp := *x
*x = *y
*y = temp
return
}



/*
Say that a "clump" in a list is a series of 2 or more adjacent elements of the same value.
Return the number of clumps in the given list.
countClumps([]int{1,2,2,3,4,4}) → 2 
countClumps([]int{1,1,2,1,1}) → 2 
countClumps([]int{1,1,1,1,1}) → 2
*/



func countClumps (numbers []int) (count int){
clump := false
for i, val := range numbres{
if i > 0{
prev := numbers[i-1]
if prev == val{
clump = true
} else if clump {
clump = false
count++
}
}
}
if clump {
count++
}
return
}

//Create your own version of cat which reads a file and dumps it to stdout.

func cat (filename string){
contents, err := ioutil.ReadFile (filename)
if err != nil{
fmt.Println("error reading file")
} else {
fmt.Println(string(contents))
}
}


//Create a program which opens a file, reads a file, then writes the contents to a new file.

func copy (oldFile string){
contents, err := ioutil.ReadFile(oldFile)
if err != nil{
fmt.Println("error reading file")
} else {
dot := strings.Index(oldFile, ".")
name := oldFile[:dot]
ext := oldFile[dot:]
newFile := name + "-copy" + ext
ioutil.WriteFile(newFile, contents, 0777)
}
}

//Create your own version of cp which reads a file and writes it to another file.

func cp (oldFile, newFile string) {
contents, err := ioutil.ReadFile (oldFile)
if err != nil{
fmt.Println("error reading file")
} else {
ioutil.WriteFile(newFile, contents, 0777)
}
}


/*Q： Why would you use an embedded anonymous unnamed field instead of a normal named field?
A: So you could access the anonymous field's properties as though they belonged to the outer object.
*/

//Create a program which converts the first character of each line in a file to uppercase and writes it to stdout.

func capitalizeLine(filename string){
file, _ := os.Open(filename)
fscanner := bufio.NewScanner(file)
for fscanner.Scan(){
line := fscanner.Text()
first := line[0]
newFirst := strings.ToUpper(string(first))
newline := newFirst + line[1:]
fmt.Println(newline)
}
}


//Create a program which capitalizes the first letter of every word from a text file and writes it to stdout.

func capitalizeWords (filename string){
file, _ : os.Open(filename)
fscanner := bufio.NewScanner(file)
for fscanner.Scan(){
line := fscanner.Text()
words := strings.Fields(line)
for _, w := range words {
newFirst := strings.ToUpper(string(w[0]))
newWord := newFirst + w[1:]
fmt.Print(newWord, " ")
}
fmt.Println()
}
}



//Create a program which capitalizes every other word (capitalizes the entire word) from a text file and writes it to stdout.

func capitalizeOddWords (filename string) {
file, _ := os.Open(filename)
fscanner := bufio.NewScanner(file)
odd := true
for fscanner.Scan(){
line := fscanner.Text()
words := strings.Fields(line)
for _, w := range words {
if odd{
fmt.Print(strings.ToUpper(w), " ")
odd = false
} else {
fmt.Print(w, " ")
odd = true
}
}
fmt.Println()
}
}

//Count how many times the word "whale" is used in Herman Melville's novel, "Moby Dick" (927 pages).
func countWhale () int {
file, _ := os.Open("mobyDick.txt")
fscanner := bufio.NewScanner(file)
count := 0

for fscanner.Scan(){
line := fscanner.Text()
words := strings.Fields(line)
for _, w := range words {
if w == "whale" {
count++
}
}
}
return count
}

//Find the longest word (string of runes not separated by spaces) used in "Moby Dick".
func longestWord () string{
file, _: = os.Open("mobyDick.txt")
fscanner := bufio.NewScanner(file)
max := 0
longestWord := ""

for fscanner.Scan() {
line := fscanner.Text()
words := strings.Fields(line)
for _, w := range words {
if len(w) > max {
max = len(w)
longestWord = w
}
}
}

return longestWord
}


//Create your own version of cat which reads an unlimited number of file and dumps their contents to stdout.
func catMany(files ...string){
for _, filename := range files {
contents, err := ioutil.ReadFile (filename)
if err != nil {
fmt.Println("error reading file")
} else {
fmt.Println(string(contents))
}
}
}



/*
implement a program which parses all of the csv file prints the following fields for each state:
id                     int
name                   string
abbreviation           string
censusRegionName       string
Make sure to use a struct in your program
*/

type State struct {
Id                     int
Name                   string
Abbreviation           string
CensusRegionName       string
}

func printCSV () {
file, _ := os.Open("state_table.csv")
defer file.Close()

reader := csv.NewReader(file)
reader.FieldsPerRecord = -1
data, _ := reader.ReadAll()

var state State
for _, each := range data {
state.Id, _ = strconv.Atoi(each[0])
state.Name = each[1]
state.Abbreviation = each[2]
state.CensusRegionName = each[3]
fmt.Println(state)
}
}

/*
implement a program which parses all of the csv file prints the following fields for each state:
id            int
name          string
abbreviation       string
censusRegionName     string

Make sure to use a struct in your program accept a state abbreviation from the user display only the requested state
*/

func printCSVbyState (abbr string) {
file, _ := os.Open("state_table.csv")
defer file.Close()



reader := csv.NewReader(file)
reader.FieldsPerRecord = -1
data, _ := reader.ReadAll()

var state State
allStates := make(map[string]State)

for _, each := range data {
state.Id, _ = strconv.Atoi(each[0])
state.Name = each[1]
state.Abbreviation = each[2]
state.CensusRegionName = each[3]
allStates[state.Abbreviation] = state
}
fmt.Println(allStates[abbr])
}


/*
implement a program which parses all of the csv file prints the following fields for each state:
id        int
name      string
abbreviation        string
censusRegionName    string
Make sure to use a struct in your program 
accept a state abbreviation from the user
write the requested state into an html file
*/

func printCSVbyStateHTML (abbr string) {
file, _ := os.Open("state_table.csv")
defer file.Close()

reader := csv.NewReader(file)
reader.FieldsPerRecord = -1
data, _ := reader.ReadAll()


var state State
allStates := make(map[string]State)
for _, each := range data{
state.Id, _ = strconv.Atoi(each[0])
state.Name = each[1]
state.Abbreviation = each[2]
state.CensusRegionName = each[3]
allStates[state.Abbreviation] = state
}

t, _ := template.New("tpl").Parse(`
{{define "State"}}
<table>
<tr>
<td>ID</td>
<td>{{ .Id }}</td>
</tr>
<tr>
<td>Name</td>
<td>{{ .Name }}</td>
</tr>
<tr>
<td>Abbreviation</td>
<td>{{ .Abbreviation }}</td>
</tr>
<tr>

<td>Census Region Name</td>
<td>{{ .CensusRegionName }}</td>
</tr>
</table>
{{end}}
`)

out, _ := os.Create("myHTMLfile.html")
_ = t.ExecuteTemplate(out, "State", allStates[abbr])
}



/*
Grab historical financial data from Yahoo as a csv file
http://finance.yahoo.com/q/hp?s=GOOG+Historical+Prices) 
read that file 
print content from each record to standard out
*/

type Price struct{
Date string
Open float64
High float64
Low float64
Close float64
Volume int
AdjClose float64
}

func printCSV2 (){
file, _ := os.Open("table.csv")
defer file.Close()

reader := csv.NewReader(file)
reader.FieldsPerRecord = -1
data, _: reader.ReadAll()

var price Price
for _, each := range data{
price.Date = each[0]
price.Open, _ = strconv.ParseFloat(each[1], 64)
price.High, _ = strconv.ParseFloat(each[2], 64)
price.Low, _ = strconv.ParseFloat(each[3], 64)
price.Close, _ = strconv.ParseFloat(each[4], 64)
price.Volume, _ = strconv.Atoi(each[5])
price.AdjClose, _ = strconv.ParseFloat(each[6], 64)
fmt.Println(price)
}
}



/*
Grab historical financial data from Yahoo as a csv file 
http://finance.yahoo.com/q/hp?s=GOOG+Historical+Prices) 
read that file
print content from each record to an html table
*/
func printCSVbyStateHTML2 () {
file, _ := os.Open("table.csv")
defer file.Close()

reader := csv.NewReader(file)
reader.FieldsPerRecord = -1
data, _ : reader.ReadAll()

var price Price
var prices []Price
for _, each := range data {
price.Date = each[0]
price.Open, _ = strconv.ParseFloat(each[1], 64)
price.High, _ = strconv.ParseFloat(each[2], 64)
price.Low, _ = strconv.ParseFloat(each[3], 64)
price.Close, _ = strconv.ParseFloat(each[4], 64)
price.Volume, _ = strconv.Atoi(each[5])
price.AdjClose, _ = strconv.ParseFloat(each[6], 64)
prices = append(prices, price)
}

t, _ := template.New("tp12").Parse(`
{{define "Price"}}
<table>
{{ range . }}
<tr>
<td>Date</td>
<td>{{ .Date }}</td>
</tr>
<tr>
<td>Open</td>
<td>{{ .Open }}</td>
</tr>
<tr>
<td>High</td>
<td>{{ .High }}</td>
</tr>
<tr>
<td>Low</td>
<td>{{ .Low }}</td>
</tr>
<tr>
<td>Volume</td>
<td>{{ .Volume }}</td>
</tr>
<tr>
<td>AdjClose</td>
<td>{{ .AdjClose }}</td>
</tr>
{{ end }}
</table>
{{end}}
`)

out, _ := os.Create("myHTMLfile2.html")
_ = t.ExecuteTemplate(out, "price", prices[1:])
}

/*Create a program which finds the md5 checksum of a file.
at the terminal
go install
yourProgramName yourFileNameToCheckSum
should return a large hexadecimal number like 'd47c2bbc28298ca9befdfbc5d3aa4e65'

*/

func hashFile (filename string) string {
contents, err := ioutil.ReadFile (filename)
if err != nil {
panic(err)
}

h := md5.New()
io.WriteString(h, string(contents))
return fmt.Sprintf("%x", h.Sum(nil))
}

//Create a program which finds the md5 checksum of all of the files in a directory.

func hashDir(dirname string){
filepath.Walk(dirname, func (path string, info os.FileInfo, err error) error {
if info.IsDir() {
return nil
}

h := hashFile(path)
fmt.Println(h)
return nil
})
}

/*
Create a program which you run from the command line reads user information (at least name and emal) from the command line (os.Args[1])
creates an HTML page generates a profile image from gravatar.com
gravatar hash; gravatar image; part of your solution should include this code:
h := md5.New()
io.WriteString(h, email)
finalBytes := h.Sum(nil)
finalString := hex.EncodeToString(finalBytes)
*/

type Person struct{
	Name string
	Email string
}
func gravatar(){
	var name, email string
	fmt.Println("Please enter your first name")
	fmt.Scanln(&name)
	
	fmt.Println("Enter your first name: ")
	fmt.Scanln(&name)
	
	fmt.Println("Enter your email: ")
	fmt.Scanln(&email)
	
	email = strings.ToLower(email)
	
	h := md5.New()
	io.WriteString(h, email)
	finalBytes := h.Sum(nil)
	finalString := hex.EncodeToString(finalBytes)
	
	//generate page
	t, _ := template.New("tp13").Parse(`
	{{define "Gravatar"}}
	<h1>{{.Name}}</h1>
	<ima src="http://www.gravatar.com/avatar/{{.Email}}"/>
	{{end}}
	`)
	out, _ := os.Create("myGravatar.html")
	_ = t.ExecuteTemplate(out, "Gravatar", Person{name, finalString})
}
	
	
func main(){
	fmt.Println("Word Count")
	fmt.Println(WordCount("test test"))
	fmt.Println()
	
	fmt.Println("Centered Average")
	fmt.Println(CenteredAverage([]float64{1,3,5,8,99}))
	fmt.Println()
	
	fmt.Println("Swap")
	x := 4
	y := 3
	swap(&x, &y)
	fmt.Println(x,y)
	fmt.Println()
	
	fmt.Println("Count Clumps")
	fmt.Println(countClumps([]int{1,2,2,3,4,4}))
	fmt.Println(countClumps([]int{1,1,2,1,1}))
	fmt.Println(countClumps([]int{1,1,1,1,1}))
	fmt.Println()
	
	fmt.Println("Cat")
	cat("test.txt")
	fmt.Println()
	
	fmt.Println("Copy")
	copy("test.txt")
	cat("test-copy.txt")
	fmt.Println()
	
	fmt.Println("cp")
	cp("test.txt", "newFile.txt")
	cat("newFile.txt")
	fmt.Println()
	
	fmt.Println("capitalizeLine")
	capitalizeLine("test.txt")
	fmt.Println()
	
	fmt.Println("capitalizeWords")
	capitalizeWords("test.txt")
	fmt.Println()
	
	fmt.Println("capitalizeOddWords")
	capitalizeOddWords("test.txt")
	fmt.Println()
	
	fmt.Println("Whale Count in Moby Dick")
	fmt.Println(countWhale())
	fmt.Println()
	fmt.Println("Longest Word in Moby Dick: ")
	fmt.Println(longestWord())
	fmt.Println()
	
	fmt.Println("Cat Many")
	catMany("test.txt", "newFile.txt", "test-copy.txt")
	fmt.Println()
	
	fmt.Println("Print CSV")
	printCSV()
	fmt.Println()
	
	fmt.Println("Print CSV by state abbr")
	printCSVbyState("CA")
	fmt.Println()
	
	fmt.Println("Print CSV2")
	printCSV2()
	fmt.Println()
	
	fmt.Println("Hash file")
	fmt.Println(hashFile("test.txt"))
	fmt.Println()
	
	fmt.Println("Hash Dir")
	hashDir(".")
	fmt.Println()
	
	fmt.Println("Print CSV by state abbr to HTML")
	printCSVbyStateHTML("WA")
	fmt.Println()
	
	fmt.Println("Print CSV prices to HTML")
	printCSVbyStateHTML2()
	fmt.Println()
	
	fmt.Println("Gravatar")
	gravatar()
	fmt.Println()
}
