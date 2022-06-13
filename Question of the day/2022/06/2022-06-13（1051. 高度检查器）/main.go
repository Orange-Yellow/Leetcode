package main

import "sort"

/*
	name: 1051. 高度检查器
	link: https://leetcode.cn/problems/height-checker/
*/

func main() {

}

func heightChecker(heights []int) int {
	expected := make([]int, len(heights))
	copy(expected, heights)
	sort.Ints(expected)
	ans := 0
	for i := 0; i < len(heights); i++ {
		if expected[i] != heights[i] {
			ans++
		}
	}
	return ans
}

func HeightChecker(heights []int) (ans int) {
	cnt := [101]int{}
	for _, v := range heights {
		cnt[v]++
	}

	idx := 0
	for i, c := range cnt {
		for ; c > 0; c-- {
			if heights[idx] != i {
				ans++
			}
			idx++
		}
	}
	return
}
