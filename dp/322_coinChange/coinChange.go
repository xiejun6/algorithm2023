package main

import "fmt"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = -1
		for _, coin := range coins {
			if i >= coin && dp[i-coin] >= 0 {
				if dp[i] == -1 || dp[i] > dp[i-coin]+1 {
					dp[i] = dp[i-coin] + 1
				}
			}
		}
	}
	return dp[amount]
}

func main() {
	coins := []int{1, 2, 5}
	amount := 11
	fmt.Println(coinChange(coins, amount))
	coins = []int{2}
	amount = 3
	fmt.Println(coinChange(coins, amount))
	coins = []int{1}
	amount = 0
	fmt.Println(coinChange(coins, amount))
}
