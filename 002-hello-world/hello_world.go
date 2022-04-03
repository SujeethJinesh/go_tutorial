package main

import "fmt"

func main() {
	fmt.Println("starting program")

	foo(2)

	fmt.Println("finished foo")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	fmt.Println("finished main")
}

func foo(stuff int) {
	fmt.Printf("in foo, %v", stuff)
}
