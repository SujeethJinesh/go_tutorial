package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("GOARCH\t\t", runtime.GOARCH)
	fmt.Println("NumCPU\t\t", runtime.NumCPU())
	fmt.Println("NumGoroutine\t", runtime.NumGoroutine())

	wg.Add(2)
	go foo()
	go foo()

	wg.Wait()
	fmt.Println("NumCPU\t\t", runtime.NumCPU())
	fmt.Println("NumGoroutine\t", runtime.NumGoroutine())
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar", i)
	}
	wg.Done()
}
