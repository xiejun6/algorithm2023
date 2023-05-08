package main

import "fmt"

func maxSubArray(nums []int) int {
	ans := nums[0]
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		if sum > 0 {
			sum += nums[i]
		} else {
			sum = nums[i]
		}
		if sum > ans {
			ans = sum
		}
	}
	return ans
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}
