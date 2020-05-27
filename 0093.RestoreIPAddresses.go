package main

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	res := []string{}
	if len(s) < 4 || len(s) > 16 {
		return res
	}

	start := 0
	cur := []string{}
	backtrack93(s, start, cur, &res)

	return res
}

func backtrack93(s string, start int, cur []string, res *[]string) {
	if len(cur) == 4 && start == len(s) {
		*res = append(*res, strings.Join(cur, "."))
		return
	}
	for i := 1; i < 4; i++ {
		if start+i > len(s) {
			break
		}
		temp := s[start : start+i]
		if isValidIP(temp, i) {
			cur = append(cur, temp)

			backtrack93(s, start+i, cur, res)

			cur = cur[:len(cur)-1]
		}
	}
}

func isValidIP(ip string, index int) bool {
	if len(ip) == 0 {
		return false
	}
	if ip[0] == '0' && len(ip) > 1 {
		return false
	}
	value, err := strconv.Atoi(ip)
	if err != nil {
		return false
	}
	if value > 255 && index == 3 {
		return false
	}
	return true
}

// https://leetcode-cn.com/problems/restore-ip-addresses/solution/jian-dan-yi-yu-li-jie-de-hui-su-fa-java-by-caipeng/
