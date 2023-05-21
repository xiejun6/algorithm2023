package main

import "fmt"

func numberOfArithmeticSlices(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	ans := 0
	for i := 2; i < n; i++ {
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			dp[i] = dp[i-1] + 1
			ans += dp[i]
		}
	}
	//fmt.Println(dp)
	return ans
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(numberOfArithmeticSlices(nums))
}
