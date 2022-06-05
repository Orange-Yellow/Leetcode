package main

import "math/rand"

/*
	name: 478. 在圆内随机生成点
	link: https://leetcode.cn/problems/generate-random-point-in-a-circle/
*/

func main() {

}

type Solution struct {
	radius, xCenter, yCenter float64
}

func Constructor(radius, xCenter, yCenter float64) Solution {
	return Solution{radius, xCenter, yCenter}
}

func (s *Solution) RandPoint() []float64 {
	for {
		x := rand.Float64()*2 - 1
		y := rand.Float64()*2 - 1 // [-1,1) 的随机数
		if x*x+y*y < 1 {
			return []float64{s.xCenter + x*s.radius, s.yCenter + y*s.radius}
		}
	}
}
