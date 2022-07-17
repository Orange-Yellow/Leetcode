package main

/*
	name: 565. 数组嵌套
	link: https://leetcode.cn/problems/array-nesting/
*/

func main() {
	ArrayNesting([]int{5, 4, 0, 3, 1, 6, 2})
}

func ArrayNesting(nums []int) int {
	ans := 0
	vis := make([]bool, len(nums))
	for i := range vis {
		cnt := 0
		for !vis[i] {
			vis[i] = true
			i = nums[i]
			cnt++
		}
		if cnt > ans {
			ans = cnt
		}
	}
	return ans
}

func arrayNesting(nums []int) int {
	ans := 0
	n := len(nums)
	for i := range nums {
		cnt := 0
		for nums[i] < n {
			// n 标记访问过的数组
			i, nums[i] = nums[i], n
			cnt++
		}
		if cnt > ans {
			ans = cnt
		}
	}
	return ans
}
