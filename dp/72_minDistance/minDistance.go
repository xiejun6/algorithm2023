package main

import "fmt"

// 插入操作和删除操作效果是一样的, 例如：ab和a, 可以第二个插入b 或者 第一个删除b
// 相等 dp[i][j] = dp[i-1][j-1]
// 不相等
// (1)插入或者删除 dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + 1
// (2)替换 dp[i][j] = dp[i-1][j-1] + 1
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
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
			}
		}
	}
	return dp[m][n]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	word1 := "horse"
	word2 := "ros"
	fmt.Println(minDistance(word1, word2))
	word1 = "intention"
	word2 = "execution"
	fmt.Println(minDistance(word1, word2))
}
