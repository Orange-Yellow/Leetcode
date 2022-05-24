package main

import (
	"fmt"
)

/*
	name: 1477. 找两个和为目标值且不重叠的子数组
	link: https://leetcode.cn/problems/find-two-non-overlapping-sub-arrays-each-with-target-sum/
*/

func main() {
	arr := []int{1, 1, 1, 2, 2, 2, 4, 4}
	target := 6

	fmt.Println(minSumOfLengths(arr, target))

}

/*
	动态规划

	-定义
		d[i]表示到编号i（不包含i）的情况下满足target目标的最短子数组的长度

	-初始化
	 	较大值表示无效

	-计算
		·因为是连续子数组，使用前缀和来解决，并且基于双指针很去计算范围即可
		·d[i] 取决于两种情况：
			能构成target，则min(d[i-1], r-l+1) ，取更短的子数组
			不能，则等于d[i-1], 即没找到直接取上一个情况的值

	-结果
		·计算d过程中不断计算满足条件的两个长度之和
		·优化： 一次遍历就可以解决，具体参见代码


*/

func minSumOfLengths(arr []int, target int) int {
	n := len(arr)
	// 长度+1 是避免只有一个子数组符合target时的情况
	dp := make([]int, n+1)
	dp[0] = n + 1
	// 结果
	ans := n + 1
	// 双指针
	l, r := 0, 0
	// 总和
	sum := 0
	for ; r < n; r++ {
		sum += arr[r]
		// 使子数组总和小于target
		for sum > target {
			// 总和大于target时，左指针不断右移使总和小于target
			sum -= arr[l]
			l++
		}
		if sum == target {
			// 当前子数组长度
			curr := r - l + 1
			// 结果为 当前值子数组长度+dp[l]长度，原来的ans 两者中取最小值
			// +dp[l] 是保证子数组不重叠
			ans = min(ans, dp[l]+curr)
			// dp
			dp[r+1] = min(dp[r], curr)
		} else {
			dp[r+1] = dp[r]
		}
	}
	if ans > n {
		return -1
	}
	return ans
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
