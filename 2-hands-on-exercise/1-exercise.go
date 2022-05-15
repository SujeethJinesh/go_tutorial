package main

import "fmt"

const (
	c     = 42
	b int = 42
)

func main() {
	a := 42

	fmt.Printf("%d\t%b\t%#x\n", a, a, a)
	fmt.Printf("%d\t%d\n", b, c)

	e := a << 1

	fmt.Printf("%b\n", e)
}
