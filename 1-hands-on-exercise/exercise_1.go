package main

import "fmt"

// var x = 42
// var y = "James Bond"
// var z = true

func main() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x, y, z)

	fmt.Printf("%d\n", x)
	fmt.Printf("%q\n", y)
	fmt.Printf("%t\n", z)
}
