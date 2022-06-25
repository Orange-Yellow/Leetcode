package main

/*
	name: 剑指 Offer II 091. 粉刷房子
	linK: https://leetcode.cn/problems/JEj789/
*/
func main() {

}

// dp 动态规划
func minCost(costs [][]int) int {
	dp := costs[0]
	for _, cost := range costs[1:] {
		dpNew := make([]int, 3)
		for j := range cost {
			//3列互不相同： j=j%3,（j+1)%3, (j+2)%3
			dpNew[j] = min(dp[(j+1)%3], dp[(j+2)%3]) + cost[j]
		}
		dp = dpNew
	}

	return min(min(dp[0], dp[1]), dp[2])
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
