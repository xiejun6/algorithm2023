package main

import "fmt"

func maxSubArray(nums []int) int {
	sum := nums[0]
	ans := nums[0]
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
	num := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(num))
	num = []int{1}
	fmt.Println(maxSubArray(num))
	num = []int{5, 4, -1, 7, 8}
	fmt.Println(maxSubArray(num))
}
