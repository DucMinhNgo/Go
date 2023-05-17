package main

import "fmt"
import "strconv"

var (
	n int = 10
	m int = 20
	h string = "abc"
)

func main() {
	// i := 11

	var n int = 1000
	// In dòng dữ liệu
	// fmt.Println("Hello World!", i)
	// In theo format (%v: giá trị, %T: kiểu dữ liệu)
	// fmt.Printf("%v, %T \n", i, i)
	fmt.Println("Global Scope", n, m, h)

	var i int = 10
	fmt.Printf("%v, %T \n", i, i)

	var j string = strconv.Itoa(i)
	fmt.Printf("%v, %T \n", j, j) 
}