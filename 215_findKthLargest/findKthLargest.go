package main

import "fmt"

func findKthLargest(nums []int, k int) int {
	heapSize := len(nums)
	for i := heapSize / 2; i >= 0; i-- {
		heapify(nums, i, heapSize)
	}

	for i := heapSize - 1; i > heapSize-k; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapify(nums, 0, i)
	}
	return nums[0]
}

func heapify(nums []int, i, heapSize int) {
	left, right := 2*i+1, 2*i+2
	max := i
	if left < heapSize && nums[left] > nums[max] {
		max = left
	}
	if right < heapSize && nums[right] > nums[max] {
		max = right
	}
	if max != i {
		nums[i], nums[max] = nums[max], nums[i]
		heapify(nums, max, heapSize)
	}
}
func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2
	fmt.Println(findKthLargest(nums, k))
	nums = []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k = 4
	fmt.Println(findKthLargest(nums, k))
}
