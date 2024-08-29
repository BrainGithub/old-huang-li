package main

import (
    "fmt"
    "strings"
    "testing"
)

func TestMap(t *testing.T) {
    var list = []string{"Hao", "Chen", "MegaEase"}
    x := MapStrToStr(list, func(s string) string {
        return strings.ToUpper(s)
    })
    fmt.Printf("%v\n", x)//["HAO", "CHEN", "MEGAEASE"]


    y := MapStrToInt(list, func(s string) int {
        return len(s)
    })
    fmt.Printf("%v\n", y)//[3, 4, 8]
}

func TestReduce(t *testing.T) {
    var list = []string{"Hao", "Chen", "MegaEase"}
    x := Reduce(list, func (s string) int {
        return len(s)
    })
    fmt.Printf("%v\n", x)
}

func TestFilter(t *testing.T) {
    var intset = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    out := Filter(intset, func(n int) bool {
        return n%2 == 1
    })
    fmt.Printf("%v\n", out)
    out = Filter(intset, func(n int) bool {
        return n > 5
    })
    fmt.Printf("%v\n", out)
}
