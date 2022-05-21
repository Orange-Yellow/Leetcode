package main

import "math/rand"

/*
	name: 961. 在长度 2N 的数组中找出重复 N 次的元素
	link: https://leetcode.cn/problems/n-repeated-element-in-size-2n-array/
*/

func main() {

}

func RepeatedNTimes(nums []int) int {
	ans := 0
	for {
		x, y := rand.Intn(len(nums)), rand.Intn(len(nums))
		if x != y && nums[x] == nums[y] {
			ans = nums[x]
			break
		}
	}
	return ans
}

// 自己的题解
func repeatedNTimes(nums []int) int {
	hash := make(map[int]int, len(nums)>>1)
	ans := 0
	for _, v := range nums {
		if _, ok := hash[v]; ok {
			ans = v
			break
		} else {
			hash[v] += 1
		}
	}
	return ans
}
