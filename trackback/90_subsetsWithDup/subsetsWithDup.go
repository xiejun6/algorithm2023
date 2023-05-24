package main

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	var ans [][]int
	var set []int
	n := len(nums)
	sort.Ints(nums)
	var dfs func(startIndex int)
	dfs = func(startIndex int) {
		ans = append(ans, append([]int{}, set...))
		for i := startIndex; i < n; i++ {
			//同一层，不要取和前一个重复的值
			if i > startIndex && nums[i] == nums[i-1] {
				continue
			}
			set = append(set, nums[i])
			dfs(i + 1)
			set = set[:len(set)-1]
		}
	}
	dfs(0)
	return ans
}

func main() {
	nums := []int{1, 2, 2}
	fmt.Println(subsetsWithDup(nums))
	nums = []int{0}
	fmt.Println(subsetsWithDup(nums))
}
