package main

import "fmt"

// 贪心算法，有钱赚就卖
func maxProfit(prices []int) int {
	ans := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			ans += prices[i] - prices[i-1]
		}
	}
	return ans
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
	prices = []int{1, 2, 3, 4, 5}
	fmt.Println(maxProfit(prices))
	prices = []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(prices))
}
