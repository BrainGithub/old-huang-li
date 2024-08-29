package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSubArray(nums []int) int {
	var n = len(nums)
	var dp = make([]int, n)
	dp[0] = nums[0]
	for i:=1; i<n; i++ {
		dp[i] = max(nums[i], dp[i-1] + nums[i])
	}

	m := dp[0]
	for _, i := range dp {
		if i > m {
			m = i
		}
	}
	return m
}

func main() {
	input := []int {-2,1,-3,4,-1,2,1,-5,4}
	m := maxSubArray(input)
	fmt.Println(m)
}
