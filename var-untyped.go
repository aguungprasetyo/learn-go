package main

import "fmt"

// untyped variable
var f = 1
var g = 3.16
var h = "Hello, Ozone Digitech"
var i = true

func untyped() {

	j := 1 //this := sign only can use inside the function

	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)
}
