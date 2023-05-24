package main

import "fmt"

func subsets(nums []int) [][]int {
	var ans [][]int
	var set []int
	n := len(nums)
	var dfs func(startIndex int)
	dfs = func(startIndex int) {
		ans = append(ans, append([]int{}, set...))
		for i := startIndex; i < n; i++ {
			set = append(set, nums[i])
			dfs(i + 1)
			set = set[:len(set)-1]
		}
	}
	dfs(0)
	return ans
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(subsets(nums))
	nums = []int{0}
	fmt.Println(subsets(nums))
}
