package main

import "fmt"

func asubsets(nums []int) [][]int {
    var ans [][]int
    var findSubset func(int)
    var res []int

    findSubset = func(i int) {
        fmt.Printf("i=%d", i)
        if i >= len(nums) {
            ans = append(ans, res[:])
            return
        }
        findSubset(i+1)
        res = append(res, nums[i])
        findSubset(i+1)
        res = res[:(len(res)-1)]
    }

    findSubset(0)


    return ans
}

func main() {
    nums := []int{1,2,3}
    ans := asubsets(nums)
    fmt.Printf("ans:%+v", ans)
}
