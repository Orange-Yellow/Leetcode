package main

/*
	name: 832. 翻转图像
	link: https://leetcode-cn.com/problems/flipping-an-image/
*/

func main() {
	image := [][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}}

	FlipAndInvertImage(image)

	flipAndInvertImage(image)
}

//FlipAndInvertImage 官方de题解
func FlipAndInvertImage(image [][]int) (ans [][]int) {
	ans = image
	for i := range image {
		l := len(image[i])
		for j := 0; j < (l+1)/2; j++ {
			re := l - 1 - j
			if ans[i][j] == ans[i][re] {
				ans[i][j] ^= 1
				ans[i][re] ^= 1
			}
			if j == re {
				ans[i][re] ^= 1
			}
		}
	}
	return
}

// 自己de题解
func flipAndInvertImage(image [][]int) (ans [][]int) {
	ans = image
	for i := range image {
		l := len(image[i])
		for j := 0; j < (l+1)/2; j++ {
			re := l - 1 - j
			ans[i][j], ans[i][re] = ans[i][re], ans[i][j]
		}
	}

	for i := range image {
		for j := range image[i] {
			if image[i][j] == 0 {
				image[i][j] = 1
			} else {
				image[i][j] = 0
			}
		}
	}

	return
}
