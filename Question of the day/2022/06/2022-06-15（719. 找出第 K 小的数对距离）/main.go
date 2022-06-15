package main

import "sort"

/*
	name: 719. 找出第 K 小的数对距离
	link: https://leetcode.cn/problems/find-k-th-smallest-pair-distance/
*/

func main() {

}

func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)
	return sort.Search(nums[len(nums)-1]-nums[0], func(mid int) bool {
		cnt := 0
		for j, num := range nums {
			i := sort.SearchInts(nums[:j], num-mid)
			cnt += j - i
		}
		return cnt >= k
	})
}
