package main

import (
	"bufio"
	_ "bufio"
	"errors"
	"fmt"
	"github.com/sheikhusmanshakeel/hello-go/pkg/crawler"
	"github.com/sheikhusmanshakeel/hello-go/pkg/indexer"
	"io"
	_ "io"
	"math"
	"math/rand"
	"os"
	_ "os"
	"reflect"
	_ "time"
)


type myFloat float64

func (f myFloat)CheckSalary() (myFloat, error) {
	if f >= 0.0 {
		return f, nil
	} else{
		return f, errors.New("myFloat: Should have non zero value")

	}
}

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) ScaleWithPointer(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Main function used for stripping and crawling
func main() {
	crawler.Crawl()
	indexer.IndexURL(56)
}

func slices_and_arrays(){
	var dy = make([][]uint8, 256)
	for y :=0; y < 256; y++	{
		dy[y] = make([]uint8, 256)
		for x:= 0; x <256; x++{
			dy[y][x] = uint8(x * y) + uint8(rand.Int63n(256))
		}
	}
}
func interesting_defer_behaviour() {
	//defer fmt.Println("World")
	//defer fmt.Println("hello")
	//defer fmt.Println("!")
	defer println("!")
	defer println("World ")
	defer println("hello ")

	fmt.Println("Sheldon Cooper")

oz:
	for i := 0; i <= 10; i++ {
		defer println(fmt.Sprintf("the number is %d", i))
		if i == 5 {
			break oz

		}
	}
	println("outside Oz")
}

func sum(a int, b int) int{
	return a+b

}
func garbage_person(){
	Usman := "Usman"
	p:= Person{
		FirstName: Usman,
		LastName:  "Shakeel",
		Hobbies:   nil,
		id:        10,
		Salary:    10,
	}
	fmt.Println(p.getFullName("some_token"));
}


type Person struct {
	FirstName string
	LastName string
	Hobbies []string
	id int64
	Salary float64
}

func (p *Person) getFullName(token string) (string, string) {
	var some string
	p.id = 15
	some = fmt.Sprintf("%s %s %d",p.FirstName,p.LastName,p.id)
	some_with_id := fmt.Sprintf("%d%s", p.id, p.LastName)
	return some, some_with_id
}

func garbage_test(){
	fmt.Println("Print something")

	var kind = reflect.Bool
	fmt.Println(kind)

	myString := ""
	arguments := os.Args
	if len(arguments) == 1{
		myString = "Please give me one argument!"
	}	else {
		myString = arguments[1]
	}
	fmt.Println(myString)
	io.WriteString(os.Stdout,myString)
	io.WriteString(os.Stdout,"\n")

	var f *os.File
	f = os.Stdin
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(">", scanner.Text())
	}


}

