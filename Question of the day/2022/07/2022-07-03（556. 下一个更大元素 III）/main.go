package main

import "math"

/*
	name: 556. 下一个更大元素 III
	link: https://leetcode.cn/problems/next-greater-element-iii/
*/

func main() {

}

func nextGreaterElement(n int) int {
	x, cnt := n, 1
	for ; x >= 10 && x/10%10 >= x%10; x /= 10 {
		cnt++
	}
	x /= 10
	if x == 0 {
		return -1
	}

	targetDigit := x % 10
	x2, cnt2 := n, 0
	for ; x2%10 <= targetDigit; x2 /= 10 {
		cnt2++
	}
	x += x2%10 - targetDigit // 把 x2%10 换到 targetDigit 上

	for i := 0; i < cnt; i++ { // 反转 n 末尾的 cnt 个数字拼到 x 后
		d := targetDigit
		if i != cnt2 {
			d = n % 10
		}
		if x > math.MaxInt32/10 || x == math.MaxInt32/10 && d > 7 {
			return -1
		}
		x = x*10 + d
		n /= 10
	}
	return x
}
