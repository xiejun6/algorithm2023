package main

import "fmt"

func maxProfit(prices []int) int {
	ans := 0
	min := prices[0]
	for i := 1; i < len(prices); i++ {
		profit := prices[i] - min
		if profit > ans {
			ans = profit
		}
		if min > prices[i] {
			min = prices[i]
		}
	}
	return ans
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))

	prices = []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(prices))
}
