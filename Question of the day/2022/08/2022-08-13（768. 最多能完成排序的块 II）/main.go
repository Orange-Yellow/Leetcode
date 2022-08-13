package main

/*
	name: 768. 最多能完成排序的块 II
	link: https://leetcode.cn/problems/max-chunks-to-make-sorted-ii/
*/

func main() {

}

func maxChunksToSorted(arr []int) int {
	st := []int{}
	for _, x := range arr {
		if len(st) == 0 || x >= st[len(st)-1] {
			st = append(st, x)
		} else {
			mx := st[len(st)-1]
			st = st[:len(st)-1]
			for len(st) > 0 && st[len(st)-1] > x {
				st = st[:len(st)-1]
			}
			st = append(st, mx)
		}
	}
	return len(st)
}
