package main

import "fmt"

// dp[i][j]表示下标i到下标j的子串是否为回文串
// dp[i][j] = dp[i+1][j-1] && s[i]==s[j]
// 长度小于等于2时, dp[i][j] = s[i]== s[j]
func longestPalindrome(s string) string {
	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	start, end := 0, 0
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if s[i] == s[j] {
				if j-i < 2 { //子串长度为1或者2
					dp[i][j] = true
					if j-i > end-start {
						start, end = i, j
					}
				} else if dp[i+1][j-1] {
					dp[i][j] = true
					if j-i > end-start {
						start, end = i, j
					}
				}
			}
		}
	}
	return s[start : end+1]
}

// 中心扩散法
func longestPalindrome2(s string) string {
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		left, right := expand(s, i, i) //长度为奇数的回文串
		if right-left > end-start {
			start, end = left, right
		}
		left, right = expand(s, i, i+1) //长度为偶数的回文串
		if right-left > end-start {
			start, end = left, right
		}
	}
	return s[start : end+1]
}

func expand(s string, left, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return left + 1, right - 1
}

func main() {
	s := "babad"
	fmt.Println(longestPalindrome(s))
	fmt.Println(longestPalindrome2("babad"))
	s = "cbbd"
	fmt.Println(longestPalindrome(s))
	fmt.Println(longestPalindrome2(s))
}
