package main

/*
	name: 926. 将字符串翻转到单调递增
	link: https://leetcode.cn/problems/flip-string-to-monotone-increasing/
*/

func main() {
	minFlipsMonoIncr("010110")
}

func minFlipsMonoIncr(s string) int {
	// dp0, dp1 分别表示 当前值为0，值为1的最小反转值
	dp0, dp1 := 0, 0
	for _, v := range s {
		// 0 的前面只能是0才能单调递增，所以等于dp0
		dpNew0 := dp0
		// 1 的前面是0和1都符合单调递增， 所以取前一个dp0或dp1中的最小值
		dpNew1 := dp1
		if dp0 < dp1 {
			dpNew1 = dp0
		}
		if v == '1' {
			dpNew0++
		} else {
			dpNew1++
		}
		dp0, dp1 = dpNew0, dpNew1
	}
	if dp0 < dp1 {
		return dp0
	}
	return dp1
}
