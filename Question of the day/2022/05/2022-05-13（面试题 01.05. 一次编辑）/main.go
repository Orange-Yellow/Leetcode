package main

/*
	name: 面试题 01.05. 一次编辑
	link: https://leetcode.cn/problems/one-away-lcci/
*/

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(oneEditAway("teacher", "beacher"))
}

// 官方题解
func OneEditAway(first string, second string) bool {
	f, s := len(first), len(second)
	if f < s {
		return OneEditAway(second, first)
	}
	// 长度相差大于1
	if f-s > 1 {
		return false
	}
	// 遍历短的字符串
	for i := range second {
		if first[i] != second[i] {
			if f == s {
				return first[i+1:] == second[i+1:]
			}
			return first[i+1:] == second[i:]
		}
	}

	return true
}

// 自己题解
func oneEditAway(first string, second string) bool {
	if math.Abs(float64(len(first)-len(second))) > 1 {
		return false
	}
	count := 0
	i, j := 0, 0
	for i < len(first) && j < len(second) {
		// 如果长度相等 双指针任何情况都同时移动
		if len(first) == len(second) {
			if first[i] != second[j] {
				if count > 1 {
					return false
				}
				count++
			}
			i++
			j++
		} else {
			// 长度不相等
			if first[i] != second[j] {
				// first[i] 与 second[j] 不相等时
				if count > 1 {
					return false
				}
				count++
				// 如果first长度大于second
				if len(first) > len(second) {
					//first移动
					i++
				} else {
					//second移动
					j++
				}
			} else {
				// 相等时双指针同时移动
				i++
				j++
			}
		}
	}

	return count <= 1
}
