package main

import "math"

/*
	name: 812. 最大三角形面积
	link: https://leetcode.cn/problems/largest-triangle-area/
*/
func main() {

}

// 枚举题解
func largestTriangleArea(points [][]int) float64 {
	var ans float64
	for i, p := range points {
		for j, q := range points[:i] {
			for _, r := range points[:j] {
				ans = math.Max(ans, triangleArea(p[0], p[1], q[0], q[1], r[0], r[1]))
			}
		}
	}
	return ans
}

// 三角形面积公式可以用行列式的绝对值表示
func triangleArea(x1, y1, x2, y2, x3, y3 int) float64 {
	return math.Abs(float64(x1*y2+x2*y3+x3*y1-x1*y3-x2*y1-x3*y2)) / 2
}
