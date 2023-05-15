package main

import "fmt"

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

func main() {
	nums := []int{1, 1, 1, 1, 1}
	target := 3
	fmt.Println(findTargetSumWays(nums, target))

	nums = []int{1}
	target = 1
	fmt.Println(findTargetSumWays(nums, target))
}
