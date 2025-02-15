package main

import "fmt"

// declare inside function
var a int
var b float32
var c string
var d bool

func main() {

	a = 1
	b = 3.16
	c = "Hello, Ozone Digitech"
	d = true
	e := 1 //this sign only can use inside the function

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}
