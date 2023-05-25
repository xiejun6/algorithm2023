package main

import (
	"fmt"
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	var ans []string
	var segment []string
	n := len(s)
	var dfs func(startIndex int)
	dfs = func(startIndex int) {
		if len(segment) > 4 {
			return
		}
		if startIndex == n && len(segment) == 4 {
			ip := strings.Join(segment, ".")
			ans = append(ans, ip)
			return
		}

		for i := startIndex; i < n; i++ {
			if s[startIndex] == '0' && i > startIndex {
				break
			}
			num, _ := strconv.Atoi(s[startIndex : i+1])
			if num > 255 {
				break
			}
			segment = append(segment, s[startIndex:i+1])
			dfs(i + 1)
			segment = segment[:len(segment)-1]
		}
	}
	dfs(0)
	return ans
}

func main() {
	s := "25525511135"
	fmt.Println(restoreIpAddresses(s))
	s = "101023"
	fmt.Println(restoreIpAddresses(s))
}
