package main

/*
	name: 745. 前缀和后缀搜索
	link: https://leetcode.cn/problems/prefix-and-suffix-search/
*/

func main() {

}

type WordFilter map[string]int

func Constructor(words []string) WordFilter {
	wf := WordFilter{}
	for i, word := range words {
		for j, n := 1, len(word); j <= n; j++ {
			for k := 0; k < n; k++ {
				wf[word[:j]+"#"+word[k:]] = i
			}
		}
	}
	return wf
}

func (wf WordFilter) F(pref, suff string) int {
	if i, ok := wf[pref+"#"+suff]; ok {
		return i
	}
	return -1
}
