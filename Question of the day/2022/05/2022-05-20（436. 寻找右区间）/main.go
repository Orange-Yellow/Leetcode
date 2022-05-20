package main

import (
	"fmt"
	"sort"
)

/*
	name: 436. 寻找右区间
	link: https://leetcode.cn/problems/find-right-interval/
*/

func main() {
	in := [][]int{{3, 4}, {2, 3}, {1, 2}}
	fmt.Println(findRightInterval(in))
}

//
func findRightInterval(intervals [][]int) []int {
	for i := range intervals {
		intervals[i] = append(intervals[i], i)
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	fmt.Println(intervals)
	ans := make([]int, len(intervals))
	for i := range intervals {
		index := sort.Search(len(intervals), func(j int) bool {
			return intervals[j][0] >= intervals[i][1]
		})
		if index < len(intervals) {
			ans[intervals[i][2]] = intervals[index][2]
		} else {
			ans[intervals[i][2]] = -1
		}
	}
	return ans
}
