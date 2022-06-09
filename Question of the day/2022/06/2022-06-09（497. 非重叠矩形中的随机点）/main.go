package main

import (
	"math/rand"
	"sort"
)

/*
	name: 497. 非重叠矩形中的随机点
	link: https://leetcode.cn/problems/random-point-in-non-overlapping-rectangles/
*/

func main() {

}

type Solution struct {
	rects [][]int
	sum   []int
}

func Constructor(rects [][]int) Solution {
	sum := make([]int, len(rects)+1)
	for i, r := range rects {
		sum[i+1] = sum[i] + (r[2]-r[0]+1)*(r[3]-r[1]+1)
	}
	return Solution{rects, sum}
}

func (s *Solution) Pick() []int {
	// 随机一个编号
	k := rand.Intn(s.sum[len(s.sum)-1])
	// 根据编号找出对应的矩形编号
	rectIndex := sort.SearchInts(s.sum, k+1) - 1
	r := s.rects[rectIndex]
	// 从左下角开始计算
	// 位于该矩形的第几个横坐标
	da := (k - s.sum[rectIndex]) / (r[3] - r[1] + 1)
	// 位于该矩形的第几个纵坐标
	db := (k - s.sum[rectIndex]) % (r[3] - r[1] + 1)
	return []int{r[0] + da, r[1] + db}
}
