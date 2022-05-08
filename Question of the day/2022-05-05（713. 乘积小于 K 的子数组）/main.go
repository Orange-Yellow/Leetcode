package main

/*
	name: 713. 乘积小于 K 的子数组
	link: https://leetcode-cn.com/problems/subarray-product-less-than-k/
*/

func main() {
	nums := []int{10, 5, 2, 6}

	NumSubarrayProductLessThanK(nums, 100)

	numSubarrayProductLessThanK(nums, 100)
}

// NumSubarrayProductLessThanK 官方de题解
func NumSubarrayProductLessThanK(nums []int, k int) int {
	ans := 0
	prod, i := 1, 0
	for j, v := range nums {
		prod *= v
		for ; i <= j && prod >= k; i++ {
			prod /= nums[i]
		}
		ans += j - i + 1
	}
	return ans
}

// 自己de题解
func numSubarrayProductLessThanK(nums []int, k int) int {
	num := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < k {
			num++
		}
		product := nums[i]
		for j := i + 1; j < len(nums); j++ {
			product *= nums[j]
			if product < k {
				num++
			} else {
				break
			}
		}
	}
	return num
}
