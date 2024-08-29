package main

import (
	"fmt"
	"time"
)

func test1() {
	defer println("in main")
	go func() {
		defer println("in goroutine")
		panic("")
	}()

	time.Sleep(1 * time.Second)
}

func main() {
	defer fmt.Println("in main")


	defer func() {

		defer func() {
			defer func() {
				recover()
				fmt.Println("in recover 3")
			}()
			fmt.Println("in 3")
			panic("panic again and again")
		}()
		fmt.Println("in 2")
		panic("panic again")
	}()
	fmt.Println("in 1")
	panic("panic once")
}

