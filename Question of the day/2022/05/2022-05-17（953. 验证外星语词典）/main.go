package main

import "fmt"

/*
	name: 953. 验证外星语词典
	link: https://leetcode.cn/problems/verifying-an-alien-dictionary/
*/

func main() {
	words := []string{"apple", "app"}
	order := "abcdefghijklmnopqrstuvwxyz"
	fmt.Println(isAlienSorted(words, order))
}

// 自己de题解
func isAlienSorted(words []string, order string) bool {
	hash := map[int32]int{}
	for i, v := range order {
		hash[v] = i
	}
	for i := 0; i < len(words)-1; i++ {
		short := len(words[i])
		if len(words[i]) > len(words[i+1]) {
			if words[i+1] == words[i][:len(words[i+1])] {
				return false
			}
			short = len(words[i+1])
		}
		for j := 0; j < short; j++ {
			x := hash[int32(words[i][j])]
			y := hash[int32(words[i+1][j])]
			if x == y {
				continue
			} else if x < y {
				break
			} else {
				return false
			}
		}
	}
	return true
}

// 官方的题解
func IsAlienSorted(words []string, order string) bool {
	index := [26]int{}
	for i, c := range order {
		index[c-'a'] = i
	}
next:
	for i := 1; i < len(words); i++ {
		for j := 0; j < len(words[i-1]) && j < len(words[i]); j++ {
			pre, cur := index[words[i-1][j]-'a'], index[words[i][j]-'a']
			if pre > cur {
				return false
			}
			if pre < cur {
				continue next
			}
		}
		if len(words[i-1]) > len(words[i]) {
			return false
		}
	}
	return true
}
