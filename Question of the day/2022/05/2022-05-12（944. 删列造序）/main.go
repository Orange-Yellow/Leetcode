package main

import "fmt"

/*
	name: 944. 删列造序
	link: https://leetcode.cn/problems/delete-columns-to-make-sorted/
*/

func main() {
	tem := []string{"cba", "daf", "ghi"}
	fmt.Println(minDeletionSize(tem))
}

// 自己的题解跟官方题解相同
func minDeletionSize(strs []string) int {
	ans := 0
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if strs[j-1][i] > strs[j][i] {
				ans++
				break
			}
		}
	}
	return ans
}
