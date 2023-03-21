package main

import "fmt"

// 右指针逐步迭代，左指针尽可能的向右靠（求最小距离嘛）
func minSubArrayLen(target int, nums []int) int {
	ans := 0
	start := 0
	sum := 0
	for end := 0; end < len(nums); end++ {
		sum += nums[end]
		for start <= end && sum >= target {
			if ans == 0 || end-start+1 < ans {
				ans = end - start + 1
			}
			sum -= nums[start]
			start++
		}
	}
	return ans
}

// 枚举左边界
func minSubArrayLen2(target int, nums []int) int {
	ans := 0
	j := 0
	sum := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		if i > 0 {
			sum -= nums[i-1]
		}
		for j < n && sum < target {
			sum += nums[j]
			j++
		}
		if sum < target {
			break
		}
		if ans == 0 || j-i < ans {
			ans = j - i
		}
	}
	return ans
}
func main() {
	target := 7
	nums := []int{2, 3, 1, 2, 4, 3}
	fmt.Println(minSubArrayLen(target, nums))
	target = 4
	nums = []int{1, 4, 4}
	fmt.Println(minSubArrayLen(target, nums))
	target = 11
	nums = []int{1, 1, 1, 1, 1, 1, 1, 1}
	fmt.Println(minSubArrayLen(target, nums))
	fmt.Println("---------------------------------")
	target = 7
	nums = []int{2, 3, 1, 2, 4, 3}
	fmt.Println(minSubArrayLen2(target, nums))
	target = 4
	nums = []int{1, 4, 4}
	fmt.Println(minSubArrayLen2(target, nums))
	target = 11
	nums = []int{1, 1, 1, 1, 1, 1, 1, 1}
	fmt.Println(minSubArrayLen2(target, nums))
}
