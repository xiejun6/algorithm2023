package main

import "fmt"

// text1[i] == text[j] dp[i][j] = dp[i-1][j-1]+1
// 不相等 dp[i][j] = max(dp[i-1][j],dp[i][j-1])
func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i, c1 := range text1 {
		for j, c2 := range text2 {
			if c1 == c2 {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	return dp[m][n]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	text1 := "abcde"
	text2 := "ace"
	fmt.Println(longestCommonSubsequence(text1, text2))
	text1 = "abc"
	text2 = "abc"
	fmt.Println(longestCommonSubsequence(text1, text2))
	text1 = "abc"
	text2 = "def"
	fmt.Println(longestCommonSubsequence(text1, text2))
}
