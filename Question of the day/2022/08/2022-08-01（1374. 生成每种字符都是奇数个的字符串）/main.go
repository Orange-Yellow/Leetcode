package main

import "strings"

/*
	name: 1374. 生成每种字符都是奇数个的字符串
	link: https://leetcode.cn/problems/generate-a-string-with-characters-that-have-odd-counts/
*/

func main() {

}

func generateTheString(n int) string {
	if n%2 == 1 {
		return strings.Repeat("x", n)
	}
	return strings.Repeat("x", n-1) + "y"
}
