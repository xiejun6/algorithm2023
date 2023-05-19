package main

import "fmt"

// 求排列，必须外层遍历目标，内层遍历物品
// 比如 nums = [1,2], target = 3
// 可以排列{1,2}, {2,1}，也就是第二个数可以选1,也可以选2,
// 所有内层遍历物品
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, v := range nums {
			if i >= v {
				dp[i] += dp[i-v]
			}
		}
	}
	return dp[target]
}

func main() {
	nums := []int{1, 2, 3}
	target := 4
	fmt.Println(combinationSum4(nums, target))
	nums = []int{9}
	target = 3
	fmt.Println(combinationSum4(nums, target))
}
