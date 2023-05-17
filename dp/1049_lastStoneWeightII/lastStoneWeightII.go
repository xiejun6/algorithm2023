package main

import "fmt"

// 本题可转换题意为：从所以石头中选若干个石头，使之和尽可能的解决sum的一半
// 可以转换为0-1背包
// dp[i][j] = max(dp[i-1][j], dp[i-1][j-stone[i]] + stone[i]
// 背包容量为sum/2,每个石头的价值为stone[i], 重量也为stone[i]
func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, v := range stones {
		sum += v
	}
	target := sum / 2
	n := len(stones)
	dp := make([]int, target+1)
	for i := 1; i <= n; i++ {
		for j := target; j >= stones[i-1]; j-- {
			dp[j] = max(dp[j], dp[j-stones[i-1]]+stones[i-1])
		}
	}
	return sum - dp[target]*2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	stones := []int{2, 7, 4, 1, 8, 1}
	fmt.Println(lastStoneWeightII(stones))
	stones = []int{31, 26, 33, 21, 40}
	fmt.Println(lastStoneWeightII(stones))
}
