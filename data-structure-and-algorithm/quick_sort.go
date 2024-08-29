package main

import "log"

func partition(nums []int, l, r int) int {
    n := nums
    m := l
    pivot := n[r]
    for i:=l; i<r; i++ {
        if n[i] < pivot {
            n[i], n[m] = n[m], n[i]
            m++
        }
    }
    n[m], n[r] = n[r], n[m]
    return m
}

func quickSort(nums []int, l, r int) {
    if l >= r {
        return
    }
    m := partition(nums, l, r)
    quickSort(nums, l, m-1)
    quickSort(nums, m+1, r)
}


func sortArray(nums []int) []int {
    if len(nums) == 1 {
        return nums
    }

    m := partion(nums)
    if m-1 > 0 {
        sortArray(nums[:m])
    }

    if m < (len(nums)-1) {
        sortArray(nums[m:])
    }
    return nums
}

func partion(nums []int) int {
    if len(nums) == 1 {
        return 0
    }

    var pivot = nums[len(nums)-1]
    j := 0
    for i:=0; i<len(nums)-1; i++ {
        if nums[i] <= pivot {
            nums[j], nums[i] = nums[i], nums[j]
            j++
        }
    }
    nums[j], nums[len(nums)-1] = pivot, nums[j]
    return j
}

func main() {
    nums := []int {-4,0,7,4,9,-5,-1,0,-7,-1}
    //quickSort(nums, 0, len(nums) - 1)
    sortArray(nums)
    log.Printf("result:%+v", nums)
    return
}