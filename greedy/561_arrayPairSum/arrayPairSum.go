package main

import (
	"fmt"
	"sort"
)

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	ans := 0
	for i := 0; i < n; i = i + 2 {
		ans += nums[i]
	}
	return ans
}

func main() {
	nums := []int{1, 4, 3, 2}
	fmt.Println(arrayPairSum(nums))
	nums = []int{6, 2, 6, 5, 1, 2}
	fmt.Println(arrayPairSum(nums))
}
