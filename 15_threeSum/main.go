package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}

	sort.Ints(nums)
	var ans [][]int
	n := len(nums)
	for first := 0; first < n; first++ {
		if nums[first] > 0 {
			break
		}
		//第一个数去重
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		left, right := first+1, n-1
		for left < right {
			sum := nums[first] + nums[left] + nums[right]
			if sum == 0 {
				ans = append(ans, []int{nums[first], nums[left], nums[right]})
				//第二个数去重
				left++
				for nums[left] == nums[left-1] && left < right {
					left++
				}
				//第三个数去重
				right--
				for nums[right] == nums[right+1] && left < right {
					right--
				}
				if left >= right {
					break
				}
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}

	}
	return ans
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
	nums = []int{0, 1, 1}
	fmt.Println(threeSum(nums))
}
