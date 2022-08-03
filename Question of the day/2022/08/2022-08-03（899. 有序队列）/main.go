package main

import (
	"fmt"
	"sort"
)

/*
	name: 899. 有序队列
	link: https://leetcode.cn/problems/orderly-queue/
*/

func main() {
	s := "baaca"
	k := 1
	fmt.Println(s[0] > s[1])
	orderlyQueue(s, k)

}

func orderlyQueue(s string, k int) string {
	if k == 1 {
		ans := s
		for i := 1; i < len(s); i++ {
			s = s[1:] + string(s[0])
			if s < ans {
				ans = s
			}
		}
		return ans
	}
	t := []byte(s)
	sort.Slice(t, func(i, j int) bool { return t[i] < t[j] })
	return string(t)
}
