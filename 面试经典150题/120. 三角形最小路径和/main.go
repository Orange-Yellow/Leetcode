package main

import "math"

func main() {
	minimumTotal([][]int{{-1}, {2, 3}, {1, -1, 3}})
}

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 1 {
		return triangle[0][0]
	}

	// 直接在三角形数组中进行操作
	for i := 1; i < len(triangle); i++ {
		triangle[i][0] = triangle[i-1][0] + triangle[i][0]
		for j := 1; j < i; j++ {
			triangle[i][j] = min(triangle[i-1][j], triangle[i-1][j-1]) + triangle[i][j]
		}
		// 每一层的最后一个元素都是第i个，上一层是i-1
		triangle[i][i] = triangle[i-1][i-1] + triangle[i][i]
	}

	ans := math.MaxInt32
	for i := range triangle[len(triangle)-1] {
		ans = min(ans, triangle[len(triangle)-1][i])
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
