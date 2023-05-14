package main

import "fmt"

// 贪心算法
func canJump(nums []int) bool {
	n := len(nums)
	maxDis := 0
	for i := 0; i < n-1; i++ {
		if maxDis >= i {
			maxDis = max(maxDis, i+nums[i])
		}
	}
	return maxDis >= n-1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(canJump(nums))
	nums = []int{3, 2, 1, 0, 4}
	fmt.Println(canJump(nums))
	nums = []int{2, 0, 0}
	fmt.Println(canJump(nums))
}
