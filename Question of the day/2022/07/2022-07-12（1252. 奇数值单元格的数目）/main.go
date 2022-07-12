package main

/*
	name: 1252. 奇数值单元格的数目
	link: https://leetcode.cn/problems/cells-with-odd-values-in-a-matrix/
*/

func main() {

}

func oddCells(m int, n int, indices [][]int) int {
	ans := 0
	rows := make([]int, m)
	cols := make([]int, n)
	for _, v := range indices {
		rows[v[0]]++
		cols[v[1]]++
	}
	for _, row := range rows {
		for _, col := range cols {
			ans += (row + col) % 2
		}
	}
	return ans
}

func OddCells(m, n int, indices [][]int) int {
	rows := make([]int, m)
	cols := make([]int, n)
	for _, p := range indices {
		rows[p[0]]++
		cols[p[1]]++
	}
	oddx := 0
	for _, row := range rows {
		oddx += row % 2
	}
	oddy := 0
	for _, col := range cols {
		oddy += col % 2
	}
	return oddx*(n-oddy) + (m-oddx)*oddy
}
