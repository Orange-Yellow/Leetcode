package main

func main() {

	stickers := []string{"with", "example", "science"}
	target := "thehat"
	minStickers(stickers, target)
}

/*
	状态压缩：设target: thehat 长为6
 	则 记忆【f】长度为 2^6=64 也即有64个子字符串
	二进制的每一位对应的就是字符串中的每一个字符： 即：6位二进制代表字符串的6位
*/
func minStickers(stickers []string, target string) int {
	// 目标字符长度
	tl := len(target)
	// 记忆化：  记录处理某一个子序列mask所需的最小贴纸数量
	f := make([]int, 1<<tl)
	for i := range f {
		f[i] = -1
	}
	f[0] = 0
	// 动态规划
	var dp func(int) int
	dp = func(mask int) int {
		if f[mask] != -1 {
			return f[mask]
		}
		f[mask] = tl + 1
		// 状态初始化，该状态至少需要tl+1个贴纸（大于总贴纸数m）
		//不在一开始初始化是为了避免递归过程中重复计算状态，该做法属于剪枝
		//下面开始计算使用贴纸拼出该状态所需的最小贴纸数
		for _, sticker := range stickers {
			left := mask
			// 统计当前贴纸包含的字母及其数量
			cnt := [26]int{}
			for _, c := range sticker {
				cnt[c-'a']++
			}
			//用该贴纸处理一遍mask状态，如果某字母可以进行提供，则按顺序将字母对应位置0
			for i, c := range target {
				if mask>>i&1 == 1 && cnt[c-'a'] > 0 {
					//倒数第i位与1异或，其他位与0异或   //异或，相同为0，不同为1
					//实质操作：对倒数第i位取反，其他位不变
					cnt[c-'a']--
					// 计算当前子序列剩余还未处理的字符
					left ^= 1 << i
				}
			}
			// 本次sticker与mask存在交集才进行下一次dp
			if left < mask {
				// //mask状态可以由left状态使用的全部贴纸+当前sticker贴纸(也就是+1)转移而来
				if f[mask] > dp(left)+1 {
					f[mask] = dp(left) + 1
				}
			}
		}
		return f[mask]
	}
	ans := dp(1<<tl - 1)
	if ans <= tl {
		return ans
	}
	return -1
}
