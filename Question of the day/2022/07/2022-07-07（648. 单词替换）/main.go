package main

import "strings"

/*
	name: 648. 单词替换
	link: https://leetcode.cn/problems/replace-words/
*/
func main() {

}

func replaceWords(dictionary []string, sentence string) string {
	dictionarySet := map[string]bool{}
	for _, s := range dictionary {
		dictionarySet[s] = true
	}
	words := strings.Split(sentence, " ")
	for i, word := range words {
		for j := 1; j <= len(word); j++ {
			if dictionarySet[word[:j]] {
				words[i] = word[:j]
				break
			}
		}
	}
	return strings.Join(words, " ")
}
