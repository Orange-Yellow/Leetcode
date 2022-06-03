package main

/*
	name: 829. 连续整数求和
	link: https://leetcode.cn/problems/consecutive-numbers-sum/
*/

func main() {

}

func isKConsecutive(n, k int) bool {
	if k%2 == 1 {
		return n%k == 0
	}
	return n%k != 0 && 2*n%k == 0
}

func consecutiveNumbersSum(n int) int {
	ans := 0
	for k := 1; k*(k+1) <= n*2; k++ {
		if isKConsecutive(n, k) {
			ans++
		}
	}
	return ans
}
