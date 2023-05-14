package main

import "fmt"

// dp[i][j] 表示0-i, 0-j的子串变成相同长度需要操作的步数
// if s1[i] == s2[j], dp[i][j] = dp[i-1][j-1]
// else dp[i][j] = min(dp[i-1][j], dp[i][j-1])+1
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}
	for i, v1 := range word1 {
		for j, v2 := range word2 {
			if v1 == v2 {
				dp[i+1][j+1] = dp[i][j]
			} else {
				dp[i+1][j+1] = min(dp[i][j+1], dp[i+1][j]) + 1
			}
		}
	}
	//fmt.Println("dp:", dp)
	return dp[m][n]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	word1 := "sea"
	word2 := "eat"
	fmt.Println(minDistance(word1, word2))

	word1 = "leetcode"
	word2 = "etco"
	fmt.Println(minDistance(word1, word2))
}
