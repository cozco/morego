package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Goroutines beginning\t", runtime.NumGoroutine())

	fmt.Println("Hello, playground")

	wg.Add(2)

	go routineOne()
	go routineTwo()

	fmt.Println("Goroutines middle\t", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("Goroutines end\t", runtime.NumGoroutine())
}

func routineOne() {
	fmt.Println("This is routine one")
	wg.Done()
}

func routineTwo() {
	fmt.Println("This is routine two")
	wg.Done()
}
