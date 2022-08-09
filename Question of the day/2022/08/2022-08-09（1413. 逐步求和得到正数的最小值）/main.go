package main

import "sort"

/*
	name: 1413. 逐步求和得到正数的最小值
	link: https://leetcode.cn/problems/minimum-value-to-get-positive-step-by-step-sum/
*/

func main() {

}

func minStartValue(nums []int) int {
	min := 0
	count := 0
	for _, num := range nums {
		count += num
		if count < min {
			min = count
		}
	}

	return 0 - min + 1
}

// MinStartValue 二分查找
func MinStartValue(nums []int) int {
	m := nums[0]
	for _, num := range nums[1:] {
		m = min(m, num)
	}
	return 1 + sort.Search(-m*len(nums), func(startValue int) bool {
		startValue++
		for _, num := range nums {
			startValue += num
			if startValue <= 0 {
				return false
			}
		}
		return true
	})
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
