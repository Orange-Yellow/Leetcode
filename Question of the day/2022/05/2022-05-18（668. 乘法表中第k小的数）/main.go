package main

import (
	"fmt"
	"sort"
)

/*
	name: 668. 乘法表中第k小的数
	link: https://leetcode.cn/problems/kth-smallest-number-in-multiplication-table/
*/

func main() {
	fmt.Println(findKthNumber(3, 3, 5))
}

func findKthNumber(m int, n int, k int) int {
	return sort.Search(m*n, func(x int) bool {
		count := x / n * n
		for i := x/n + 1; i <= m; i++ {
			count += x / i
		}
		return count >= k
	})
}
