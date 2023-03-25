package main

import "fmt"

func findDisappearedNumbers(nums []int) []int {
	for _, v := range nums {
		if v < 0 {
			v = -v
		}
		if nums[v-1] > 0 {
			nums[v-1] = -nums[v-1]
		}
	}
	var ans []int
	for i, v := range nums {
		if v > 0 {
			ans = append(ans, i+1)
		}
	}
	return ans
}

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDisappearedNumbers(nums))
	nums = []int{1, 1}
	fmt.Println(findDisappearedNumbers(nums))
}
