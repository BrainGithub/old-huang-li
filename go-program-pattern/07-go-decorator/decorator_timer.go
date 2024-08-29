package main

import "fmt"

type handler func(s string)

func(h handler) timing(s string) {
	fmt.Println("Started")
	h(s)
	fmt.Println("Done")
}

func timing2(f func (string), s string) {
	fmt.Println("Timing Started")
	f(s)
	fmt.Println("Timing Done")
}

func decorator(f func(s string)) func(s string) {
	return func(s string) {
		fmt.Println("Started")
		f(s)
		fmt.Println("Done")
	}
}

func Hello(s string) {
	fmt.Println(s)
}


func main() {
	decorator(Hello)("Hello, World!")
	handler(Hello).timing("hhhhhhhhh")
	timing2(Hello, "timing2222222")
}
