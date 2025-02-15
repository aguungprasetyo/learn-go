package main

import "fmt"

// declare inside function
var a1 int
var b1 float32
var c1 string
var d1 bool

func outside() {

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
