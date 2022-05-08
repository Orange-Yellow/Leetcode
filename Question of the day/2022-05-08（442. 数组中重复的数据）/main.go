package main

/*
	name: 442. 数组中重复的数据
	link: https://leetcode-cn.com/problems/find-all-duplicates-in-an-array/
*/

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}

	//FindDuplicates(nums)

	FindDuplicates2(nums)

	//findDuplicates(nums)
}

// 官方de题解
func FindDuplicates(nums []int) []int {
	ans := []int{}

	for i := range nums {
		// 因为1<=nums[i]<=n,所以让nums[i] - 1 = i
		for nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	for i := range nums {
		if i != nums[i]-1 {
			ans = append(ans, nums[i])
		}
	}

	return ans
}

// 官方题解2
func FindDuplicates2(nums []int) (ans []int) {
	for _, x := range nums {
		if x < 0 {
			x = -x
		}
		if nums[x-1] > 0 {
			nums[x-1] = -nums[x-1]
		} else {
			ans = append(ans, x)
		}
	}
	return
}

// 自己de题解
func findDuplicates(nums []int) []int {
	ans := []int{}
	numsMap := map[int]struct{}{}
	for _, v := range nums {
		if _, ok := numsMap[v]; ok {
			ans = append(ans, v)
		} else {
			numsMap[v] = struct{}{}
		}
	}
	return ans
}
