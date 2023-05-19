package main

import "fmt"

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = i //最极端的情况，全是由1组成
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	n := 12
	fmt.Println(numSquares(n))
	n = 13
	fmt.Println(numSquares(n))
}
