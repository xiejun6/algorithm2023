package main

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	ans := math.MaxInt32
	update := func(sum int) {
		if abs(sum-target) < abs(ans-target) {
			ans = sum
		}
	}
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		//left第二个数 right第三个数
		left, right := i+1, n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == target {
				return sum
			}
			update(sum)
			if sum < target {
				left++
				for left < right && nums[left] == nums[left-1] {
					left++
				}
			} else {
				right--
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			}
		}
	}
	return ans
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func main() {
	nums := []int{-1, 2, 1, -4}
	target := 1
	fmt.Println(threeSumClosest(nums, target))
	nums = []int{0, 0, 0}
	target = 1
	fmt.Println(threeSumClosest(nums, target))
}
