package main

/*
	name: 676. 实现一个魔法字典
	link: https://leetcode.cn/problems/implement-magic-dictionary/
*/

func main() {

}

type MagicDictionary []string

func Constructor() MagicDictionary {
	return MagicDictionary{}
}

func (d *MagicDictionary) BuildDict(dictionary []string) {
	*d = dictionary
}

func (d *MagicDictionary) Search(searchWord string) bool {
next:
	for _, word := range *d {
		if len(word) != len(searchWord) {
			continue
		}
		diff := false
		for i := range word {
			if word[i] != searchWord[i] {
				if diff {
					continue next
				}
				diff = true
			}
		}
		if diff {
			return true
		}
	}
	return false
}
