package main

import "fmt"

func findSubsequences(nums []int) [][]int {
	var ans [][]int
	var seq []int
	n := len(nums)
	var dfs func(startIndex int)
	dfs = func(startIndex int) {
		if len(seq) >= 2 {
			ans = append(ans, append([]int{}, seq...))
		}
		//本次每个数字只能用一次 用 map对本层去重，不用回溯
		used := make(map[int]struct{})
		for i := startIndex; i < n; i++ {
			if _, ok := used[nums[i]]; ok {
				continue
			}
			if len(seq) > 0 && seq[len(seq)-1] > nums[i] {
				continue
			}
			seq = append(seq, nums[i])
			used[nums[i]] = struct{}{} //只在本层用到这个遍历，不需要回溯
			dfs(i + 1)
			seq = seq[:len(seq)-1]
		}
	}
	dfs(0)
	return ans
}

func main() {
	nums := []int{4, 6, 7, 7}
	fmt.Println(findSubsequences(nums))
	nums = []int{4, 4, 3, 2, 1}
	fmt.Println(findSubsequences(nums))
}
