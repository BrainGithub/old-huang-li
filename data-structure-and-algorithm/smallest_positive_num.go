package main

import "fmt"

const INT_MAX = int(^uint(0) >> 1)

func smallestPositiveNums(nums []int) int {
    cap := map[int]bool{}
    for _, i := range nums {
        cap[i] = true
    }

    for i := 1;i<INT_MAX ;i++ {
        if !cap[i] {
            return i
        }
    }
    return 0
}

func main() {
    nums := []int{1, 3, 6, 4, 1, 2}
    fmt.Printf("result:%d", smallestPositiveNums(nums))
}
