package main

import "fmt"

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}
	i := 2
	for j := 2; j < n; j++ {
		if nums[j] != nums[i-2] {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	fmt.Println(removeDuplicates(nums))
	nums = []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	fmt.Println(removeDuplicates(nums))
}
