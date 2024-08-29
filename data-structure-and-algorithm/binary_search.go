package main

import "log"

func search(nums []int, target int) int {
    low, high := 0, len(nums) -1
    for low <= high {
        mid := (low + high)/2
        if target > nums[mid] {
            low = mid + 1
        } else if target < nums[mid] {
            high = mid - 1
        } else {
            return target
        }
    }
    return -1
}

func main() {
    nums := []int{1,4,6,7}
    log.Printf("result:%d", search(nums, 5))
}
