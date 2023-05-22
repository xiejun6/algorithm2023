package main

import "fmt"

func permute(nums []int) [][]int {
	var ans [][]int
	var seq []int
	var dfs func(nums []int)
	dfs = func(nums []int) {
		if len(nums) == 0 {
			ans = append(ans, append([]int{}, seq...))
			return
		}
		for i := 0; i < len(nums); i++ {
			nums[0], nums[i] = nums[i], nums[0]
			seq = append(seq, nums[0])
			dfs(nums[1:])
			nums[0], nums[i] = nums[i], nums[0]
			seq = seq[:len(seq)-1]
		}
	}
	dfs(nums)
	return ans
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
	nums = []int{0, 1}
	fmt.Println(permute(nums))
	nums = []int{1}
	fmt.Println(permute(nums))
}
