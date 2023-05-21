package main

import "fmt"

func maxProfit(prices []int) int {
	profit := 0
	for i := range prices {
		profit = max(profit, trade(prices[:i])+trade(prices[i:]))
	}
	return profit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func trade(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	ans := 0
	min := prices[0]
	for i := 1; i < n; i++ {
		if prices[i]-min > ans {
			ans = prices[i] - min
		} else if prices[i] < min {
			min = prices[i]
		}
	}
	return ans
}

// dp[i][0] 到第i天只买入一次
// d[i][1] 到第i天只卖出一次
// dp[i][2] 到第i天完成第二次买入操作
// dp[i][3] 到第i天完成第二次卖出操作
func maxProfit2(prices []int) int {
	n := len(prices)
	dp := make([][4]int, n)
	dp[0][0] = -prices[0]
	dp[0][2] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], -prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
		dp[i][2] = max(dp[i-1][2], dp[i-1][1]-prices[i])
		dp[i][3] = max(dp[i-1][3], dp[i-1][2]+prices[i])
	}
	return dp[n-1][3]
}

func main() {
	prices := []int{3, 3, 5, 0, 0, 3, 1, 4}
	fmt.Println(maxProfit(prices))
	fmt.Println(maxProfit2(prices))
	prices = []int{1, 2, 3, 4, 5}
	fmt.Println(maxProfit(prices))
	fmt.Println(maxProfit2(prices))
	prices = []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(prices))
	fmt.Println(maxProfit2(prices))
	prices = []int{1}
	fmt.Println(maxProfit(prices))
	fmt.Println(maxProfit2(prices))
}
