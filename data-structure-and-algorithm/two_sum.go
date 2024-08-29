package main

import (
    "log"
    "sort"
)

func twoSum(nums []int, target int) []int {
    var index = make(map[int]int, len(nums))
    for i:=0; i<len(nums); i++ {
        index[nums[i]] = i
    }
    sort.Ints(nums)
    for i:=0; i<len(nums); i++ {
        for j:=len(nums)-1; j>i && nums[i] + nums[j] >= target; j-- {
            if i<j && nums[i] + nums[j] == target {
                return []int{index[nums[i]], index[nums[j]]}
            }
        }
    }
    return []int{}
}

func main() {
    input := []int {3,3}
    output := twoSum(input, 6)
    log.Printf("outputs:+%v", output)
}
