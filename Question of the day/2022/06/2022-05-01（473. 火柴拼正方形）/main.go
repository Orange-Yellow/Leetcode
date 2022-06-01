package main

/*
	name: 473. 火柴拼正方形
	link: https://leetcode.cn/problems/matchsticks-to-square/
*/

func main() {

}

func makesquare(matchsticks []int) bool {
	totalLen := 0
	for _, l := range matchsticks {
		totalLen += l
	}
	if totalLen%4 != 0 {
		return false
	}

	tLen := totalLen / 4
	dp := make([]int, 1<<len(matchsticks))
	for i := 1; i < len(dp); i++ {
		dp[i] = -1
	}
	for s := 1; s < len(dp); s++ {
		for k, v := range matchsticks {
			if s>>k&1 == 0 {
				continue
			}
			s1 := s &^ (1 << k)
			if dp[s1] >= 0 && dp[s1]+v <= tLen {
				dp[s] = (dp[s1] + v) % tLen
				break
			}
		}
	}
	return dp[len(dp)-1] == 0
}
