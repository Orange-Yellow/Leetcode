package main

import "strings"

/*
	name: 1408. 数组中的字符串匹配
	link: https://leetcode.cn/problems/string-matching-in-an-array/
*/

func main() {

}

func stringMatching(words []string) (ans []string) {
	for i, x := range words {
		for j, y := range words {
			if j != i && strings.Contains(y, x) {
				ans = append(ans, x)
				break
			}
		}
	}
	return
}
