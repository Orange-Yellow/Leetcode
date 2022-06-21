package main

import "strings"

/*
	name: 1108. IP 地址无效化
	link: https://leetcode.cn/problems/defanging-an-ip-address/
*/

func main() {

}

func defangIPaddr(address string) string {
	return strings.ReplaceAll(address, ".", "[.]")
}
