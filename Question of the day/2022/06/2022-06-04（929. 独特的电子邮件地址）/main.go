package main

import (
	"fmt"
	"strings"
)

/*
	name: 929. 独特的电子邮件地址
	link: https://leetcode.cn/problems/unique-email-addresses/
*/

func main() {
	x := []string{"test.email+alex@leetcode.com", "test.email.leet+alex@code.com"}
	NumUniqueEmails(x)
	fmt.Println(strings.Split("x+y+z", "+"))

}
func NumUniqueEmails(emails []string) int {
	emailSet := map[string]struct{}{}
	for _, email := range emails {
		i := strings.Index(email, "@")
		local := strings.SplitN(email[:i], "+", 2)[0] // 去掉本地名第一个加号之后的部分
		local = strings.ReplaceAll(local, ".", "")    // 去掉本地名中所有的句点
		emailSet[local+email[i:]] = struct{}{}
	}
	return len(emailSet)
}

func numUniqueEmails(emails []string) int {
	hash := map[string]struct{}{}
	for _, v := range emails {
		list := strings.Split(v, "@")
		local := list[0]
		index := strings.Index(local, ".")
		for index != -1 {
			local = local[:index] + local[index+1:]
			index = strings.Index(local, ".")
		}
		j := strings.Index(local, "+")
		for j != -1 {
			local = local[:j]
			j = strings.Index(local, "+")
		}
		if _, ok := hash[local+"@"+list[1]]; !ok {
			hash[local+"@"+list[1]] = struct{}{}
		}
	}
	return len(hash)
}
