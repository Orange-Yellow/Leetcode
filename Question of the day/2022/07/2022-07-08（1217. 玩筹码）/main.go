package main

/*
	name: 1217. 玩筹码
	link: https://leetcode.cn/problems/minimum-cost-to-move-chips-to-the-same-position/
*/

func main() {

}

func minCostToMoveChips(position []int) int {
	cnt := [2]int{}
	for _, p := range position {
		cnt[p%2]++
	}
	return min(cnt[0], cnt[1])
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
