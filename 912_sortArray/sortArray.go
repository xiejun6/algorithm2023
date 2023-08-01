package main

import "fmt"

func sortArray(nums []int) []int {
	n := len(nums)
	if n <= 1 {
		return nums
	}
	left, right := 0, n-1
	pivot := nums[left]
	for left < right {
		for left < right && nums[right] >= pivot {
			right--
		}
		nums[left] = nums[right]
		for left < right && nums[left] < pivot {
			left++
		}
		nums[right] = nums[left]
	}
	nums[left] = pivot
	sortArray(nums[:left])
	sortArray(nums[left+1:])
	return nums
}

func sortArray1(nums []int) []int {
	qSort(nums, 0, len(nums)-1)
	return nums
}

func qSort(nums []int, left, right int) {
	if left < right {
		pos := partition(nums, left, right)
		qSort(nums, left, pos-1)
		qSort(nums, pos+1, right)
	}
}
func partition(nums []int, left, right int) int {
	pivot := nums[left]
	for left < right {
		for left < right && nums[right] >= pivot {
			right--
		}
		nums[left] = nums[right]
		for left < right && nums[left] < pivot {
			left++
		}
		nums[right] = nums[left]
	}
	nums[left] = pivot
	return left
}

func main() {
	nums := []int{5, 2, 3, 1}
	n1 := sortArray1(nums)
	fmt.Println(n1)
	nums = []int{5, 1, 1, 2, 0, 0}
	n1 = sortArray1(nums)
	fmt.Println(n1)
}
