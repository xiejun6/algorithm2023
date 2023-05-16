package main

import "fmt"

// 转换为0-1背包问题
// 选取一部分数字 使得和为数组加起来总和的一半
func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	dp := make([]bool, target+1)
	dp[0] = true
	for i := 1; i <= len(nums); i++ {
		for j := target; j >= nums[i-1]; j-- {
			dp[j] = dp[j] || dp[j-nums[i-1]]
		}
		//fmt.Println("i =", i, "dp:", dp)
	}
	return dp[target]
}

func main() {
	nums := []int{1, 5, 11, 5}
	fmt.Println(canPartition(nums))
}
