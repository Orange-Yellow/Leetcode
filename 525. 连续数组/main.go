package main

import "fmt"

func main() {
	fmt.Println(findMaxLength([]int{0, 1, 0}))
}

func findMaxLength(nums []int) int {
	ans := 0
	mp := map[int]int{0: -1}

	count := 0
	for i, v := range nums {
		if v == 1 {
			count++
		} else {
			count--
		}
		/*
			由于哈希表存储的是 \textit{counter}counter 的每个取值第一次出现的下标，
			因此当遇到重复的前缀和时，根据当前下标和哈希表中存储的下标计算得到的子数组长度是以当前下标结尾的子数组中满足有相同数量的 0 和 1 的最长子数组的长度。
			遍历结束时，即可得到 \textit{nums}nums 中的有相同数量的 00 和 11 的最长子数组的长度。
		*/
		if preIndex, ok := mp[count]; ok {
			ans = max(ans, i-preIndex)
		} else {
			mp[count] = i
		}
	}

	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
