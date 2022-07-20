package main

/*
	name: 1260. 二维网格迁移
	link: https://leetcode.cn/problems/shift-2d-grid/
*/

func main() {

}

func shiftGrid(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
	}
	for i, row := range grid {
		for j, v := range row {
			index1 := (i*n + j + k) % (m * n)
			ans[index1/n][index1%n] = v
		}
	}
	return ans
}
