package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	index := m + n - 1
	i, j := m-1, n-1
	//从后往前放
	for j >= 0 {
		//i >= 0 这个判断不能少
		if i >= 0 && nums1[i] >= nums2[j] {
			nums1[index] = nums1[i]
			i--
		} else {
			nums1[index] = nums2[j]
			j--
		}
		index--
	}
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}
