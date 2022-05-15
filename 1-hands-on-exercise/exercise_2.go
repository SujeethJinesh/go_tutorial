package main

import "fmt"

var x int
var y string
var z bool

func main() {
	fmt.Println(x, y, z)

	x = 42
	y = "James Bond"
	z = true

	s := fmt.Sprintf("%d, %s, %t", x, y, z)

	fmt.Println(s)
}
