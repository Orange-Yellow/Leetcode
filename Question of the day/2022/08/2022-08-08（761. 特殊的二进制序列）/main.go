package main

import (
	"sort"
	"strings"
)

/*
	name: 761. 特殊的二进制序列
	link: https://leetcode.cn/problems/special-binary-string/
*/

func main() {

}

func makeLargestSpecial(s string) string {
	if len(s) <= 2 {
		return s
	}
	subs := sort.StringSlice{}
	cnt, left := 0, 0
	for i, ch := range s {
		if ch == '1' {
			cnt++
		} else if cnt--; cnt == 0 {
			subs = append(subs, "1"+makeLargestSpecial(s[left+1:i])+"0")
			left = i + 1
		}
	}
	sort.Sort(sort.Reverse(subs))
	return strings.Join(subs, "")
}
