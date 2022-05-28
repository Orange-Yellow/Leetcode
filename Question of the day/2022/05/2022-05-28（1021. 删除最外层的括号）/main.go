package main

/*
	name: 1021. 删除最外层的括号
	link: https://leetcode.cn/problems/remove-outermost-parentheses/
*/

func main() {

}

func RemoveOuterParentheses(s string) string {
	var ans, st []rune
	for _, v := range s {
		if v == ')' {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			ans = append(ans, v)
		}
		if v == '(' {
			st = append(st, v)
		}
	}
	return string(ans)
}

func removeOuterParentheses(s string) string {
	ans, flag := []rune{}, 0
	for _, v := range s {
		if v == ')' {
			flag--
		}
		if flag > 0 {
			ans = append(ans, v)
		}
		if v == '(' {
			flag++
		}
	}
	return string(ans)
}
