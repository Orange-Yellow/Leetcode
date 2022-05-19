package main

import (
	"fmt"
	"sort"
)

/*
	name: 462. 最少移动次数使数组元素相等 II
	link: https://leetcode.cn/problems/minimum-moves-to-equal-array-elements-ii/
*/

func main() {
	x := []int{1, 10, 2, 9}
	fmt.Println(minMoves2(x))
}

func minMoves2(nums []int) int {
	sort.Ints(nums)
	mid := nums[len(nums)/2]
	ans := 0
	for i := range nums {
		if nums[i]-mid < 0 {
			ans += -(nums[i] - mid)
		} else {
			ans += nums[i] - mid
		}
	}
	return ans
}
