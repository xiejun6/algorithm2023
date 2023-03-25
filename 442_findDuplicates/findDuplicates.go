package main

import "fmt"

// 数字为1-N,
// 如果是1就是把第一个数置为负数，
// 如果置换前已经是负数了，就说明重复了
func findDuplicates(nums []int) []int {
	var ans []int
	for _, v := range nums {
		if v < 0 {
			v = -v
		}
		if nums[v-1] > 0 {
			nums[v-1] = -nums[v-1]
		} else {
			ans = append(ans, v)
		}
	}
	return ans
}

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDuplicates(nums))
	nums = []int{1, 1, 2}
	fmt.Println(findDuplicates(nums))
}
