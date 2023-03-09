package main

import "fmt"

// 496下一个更大的元素 变体
func nextGreaterElements(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}
	var stack []int
	ans := make([]int, len(nums))
	for i := range ans {
		ans[i] = -1
	}

	for i, v := range nums {
		for len(stack) > 0 && v > nums[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[top] = v
		}
		stack = append(stack, i)
	}
	if len(stack) == 1 {
		return ans
	}
	for _, v := range nums {
		for len(stack) > 0 && v > nums[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[top] = v
		}
		if len(stack) == 0 {
			break
		}
	}
	return ans
}

func main() {
	nums := []int{1, 2, 1}
	fmt.Println(nextGreaterElements(nums))
	nums = []int{1, 2, 3, 4, 3}
	fmt.Println(nextGreaterElements(nums))
	nums = []int{4, 3, 2}
	fmt.Println(nextGreaterElements(nums))
	nums = []int{-2, 0, -3, -3}
	fmt.Println(nextGreaterElements(nums))
}
