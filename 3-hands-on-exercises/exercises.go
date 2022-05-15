package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(i)
	}

	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}

	a := 1997
	for a < 2023 {
		fmt.Println(a)
		a++
	}

	b := 1997
	for {
		fmt.Println(b)
		b++
		if b > 2022 {
			break
		}
	}

	for i := 10; i <= 100; i++ {
		fmt.Printf("When %v is divided by 4, the answer is %v", i, i%4)
	}
}
