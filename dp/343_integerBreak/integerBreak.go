package main

import "fmt"

func integerBreak(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 1
	} else if n == 3 {
		return 2
	} else if n == 4 {
		return 4
	}
	val := 1
	for n > 4 {
		val *= 3
		n -= 3
	}
	val *= n
	return val
}

func integerBreak2(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(dp[i], max(dp[i-j]*j, (i-j)*j))
		}
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	fmt.Println(integerBreak(10))
	fmt.Println(integerBreak2(10))
}
