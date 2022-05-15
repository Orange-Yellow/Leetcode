package main

/*
	name: 942. 增减字符串匹配
	link: https://leetcode.cn/problems/di-string-match/
*/

func main() {
	diStringMatch("DDI")
}

// 自己de题解 ans 官方的题解
func diStringMatch(s string) []int {
	n := len(s)
	ans := make([]int, n+1)

	l, h := 0, n

	for i := range s {
		if s[i] == 'D' {
			ans[i] = h
			h--
		} else {
			ans[i] = l
			l++
		}
	}
	ans[n] = l
	return ans
}
