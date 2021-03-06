package main

/*
	name: 890. 查找和替换模式
	link: https://leetcode.cn/problems/find-and-replace-pattern/
*/

func main() {

}

func match(word, pattern string) bool {
	mp := map[rune]byte{}
	for i, x := range word {
		y := pattern[i]
		if mp[x] == 0 {
			mp[x] = y
		} else if mp[x] != y { // word 中的同一字母必须映射到 pattern 中的同一字母上
			return false
		}
	}
	return true
}

func findAndReplacePattern(words []string, pattern string) (ans []string) {
	for _, word := range words {
		if match(word, pattern) && match(pattern, word) {
			ans = append(ans, word)
		}
	}
	return
}
