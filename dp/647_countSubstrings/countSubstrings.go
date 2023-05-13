package main

import "fmt"

// 方法1 中心扩散法
func countSubstrings(s string) int {
	count := 0
	for i := range s {
		count += extends(s, i, i)
		count += extends(s, i, i+1)
	}
	return count
}

func extends(s string, left int, right int) int {
	n := len(s)
	count := 0
	for left >= 0 && right < n && s[left] == s[right] {
		count++
		left--
		right++
	}
	return count
}

// 方法2 dp[i][j] = dp[i-1][j+1] && s[i] == s[j]
func countSubstrings2(s string) int {
	count := 0
	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if s[i] == s[j] {
				if j-i < 2 { //长度为1或者为2
					dp[i][j] = true
					count++
				} else if dp[i+1][j-1] {
					dp[i][j] = true
					count++
				}
			}
		}
	}
	return count
}

func main() {
	s := "abc"
	fmt.Println(countSubstrings(s))
	fmt.Println(countSubstrings2(s))
	s = "aaa"
	fmt.Println(countSubstrings(s))
	fmt.Println(countSubstrings2(s))
}
