package main

import "fmt"

// 暴力搜索
func findTargetSumWays(nums []int, target int) int {
	count := 0
	size := len(nums)
	var dfs func(index int, sum int)
	dfs = func(index int, sum int) {
		if index >= size {
			if sum == target {
				count++
			}
			return
		}

		dfs(index+1, sum+nums[index])
		dfs(index+1, sum-nums[index])
	}
	dfs(0, 0)
	return count
}

// 正数之和：a
// 负数之后: b
// 总和为:sum
// a+b = sum, a-b= target 可得 a = (sum+target)/2, b = (sum-target)/2
// 转换为 dp[i][j] 前i个数 和为j的有几种方案
// dp[i][j] = d[i-1][j] + dp[i-1][j-num[i]] (不选第i个数+选第i个数)
// 用滚动数组压缩成一维, dp[j] = dp[j]+dp[j-num[i]]
func findTargetSumWays2(nums []int, target int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	diff := sum - target
	//判断小于0很重要，不然初始化slice会报错
	if diff < 0 || diff%2 != 0 {
		return 0
	}
	if diff%2 != 0 {
		return 0
	}
	total := diff / 2
	dp := make([]int, total+1)
	dp[0] = 1
	for i := 1; i <= len(nums); i++ {
		for j := total; j >= nums[i-1]; j-- {
			dp[j] += dp[j-nums[i-1]]
		}
	}
	return dp[total]
}

func main() {
	nums := []int{1, 1, 1, 1, 1}
	target := 3
	fmt.Println(findTargetSumWays(nums, target))
	fmt.Println(findTargetSumWays2(nums, target))

	nums = []int{1}
	target = 1
	fmt.Println(findTargetSumWays(nums, target))
	fmt.Println(findTargetSumWays2(nums, target))
}
