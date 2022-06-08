package main

/*
	name: 1037. 有效的回旋镖
	link: https://leetcode.cn/problems/valid-boomerang/
*/
func main() {

}

func isBoomerang(points [][]int) bool {
	return (points[0][0]*points[1][1] + points[1][0]*points[2][1] + points[2][0]*points[0][1] - points[1][0]*points[0][1] - points[2][0]*points[1][1] - points[0][0]*points[2][1]) != 0
}
