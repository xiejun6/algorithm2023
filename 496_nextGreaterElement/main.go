package main

import "fmt"

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var stack []int
	m := make(map[int]int)
	for _, v := range nums2 {
		for len(stack) > 0 && stack[len(stack)-1] < v {
			num := stack[len(stack)-1]
			m[num] = v
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, v)
	}

	ans := make([]int, len(nums1))
	for i, v := range nums1 {
		if num, ok := m[v]; ok {
			ans[i] = num
		} else {
			ans[i] = -1
		}
	}
	return ans
}

func main() {
	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}
	fmt.Println(nextGreaterElement(nums1, nums2))
	nums1 = []int{2, 4}
	nums2 = []int{1, 2, 3, 4}
	fmt.Println(nextGreaterElement(nums1, nums2))
}
