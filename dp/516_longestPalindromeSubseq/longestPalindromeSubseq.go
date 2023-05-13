package main

import "fmt"

// dp[i][j]表示i到j的子串的回文串长度,
// dp[i][j] =1
// 如果i < j
// 1、s[i] == s[j] ,dp[i][j] = dp[i+1][j-1] + 2
// 2、s[i] != s[j], dp[i][j] = max(dp[i+1][j], dp[i][j-1]
func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2

			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	s := "bbbab"
	fmt.Println(longestPalindromeSubseq(s))
	s = "cbbd"
	fmt.Println(longestPalindromeSubseq(s))
}
