package main

import (
	"fmt"
)

/*
	name: 467. 环绕字符串中唯一的子字符串
	link: https://leetcode.cn/problems/unique-substrings-in-wraparound-string/
*/

func main() {
	fmt.Println(findSubstringInWraproundString("zab"))
}

func findSubstringInWraproundString(p string) int {
	ans := 0
	// dp[a] 表示以a结尾的字串个数
	dp := [26]int{}
	k := 0
	for i, ch := range p {
		if i > 0 && (byte(ch)-p[i-1]+26)%26 == 1 { // 字符之差为 1 或 -25
			k++
		} else {
			k = 1
		}
		if k > dp[ch-'a'] {
			dp[ch-'a'] = k
		}
	}
	for _, v := range dp {
		ans += v
	}
	return ans

}
