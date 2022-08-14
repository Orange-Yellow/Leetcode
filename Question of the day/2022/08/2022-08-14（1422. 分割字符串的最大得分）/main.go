package main

import "strings"

/*
	name: 1422. 分割字符串的最大得分
	link: https://leetcode.cn/problems/maximum-score-after-splitting-a-string/
*/

func main() {

}

func maxScore(s string) int {
	score := int('1'-s[0]) + strings.Count(s[1:], "1")
	ans := score
	for _, c := range s[1 : len(s)-1] {
		if c == '0' {
			score++
		} else {
			score--
		}
		ans = max(ans, score)
	}
	return ans
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
