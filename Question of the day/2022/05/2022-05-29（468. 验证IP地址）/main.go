package main

import (
	"strconv"
	"strings"
)

/*
	name: 468. 验证IP地址
	link: https://leetcode.cn/problems/validate-ip-address/
*/

func main() {

}

func validIPAddress(queryIP string) string {
	if sl := strings.Split(queryIP, "."); len(sl) == 4 {
		for _, s := range sl {
			if len(s) > 1 && s[0] == '0' {
				return "Neither"
			}
			if v, err := strconv.Atoi(s); err != nil || v > 255 {
				return "Neither"
			}
			return "IPv4"
		}

	}
	if sl := strings.Split(queryIP, ":"); len(sl) == 8 {
		for _, s := range sl {
			if len(s) > 4 {
				return "Neither"
			}
			if _, err := strconv.ParseInt(s, 16, 64); err != nil {
				return "Neither"
			}
			return "IPv6"
		}
	}
	return "Neither"
}
