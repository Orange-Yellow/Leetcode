package main

import "math"

func main() {
	minSubArrayLen(6, []int{2, 3, 2, 2, 4, 3})
}

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)

	if n == 0 {
		return 0
	}

	ans := math.MaxInt32
	left, right := 0, 0

	sum := 0

	for right < n {
		sum += nums[right]
		for sum >= target {
			ans = min(ans, right-left+1)
			sum -= nums[left]
			left++
		}
		right++
	}

	if ans == math.MaxInt32 {
		return 0
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
