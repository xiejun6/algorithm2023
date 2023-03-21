package main

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {
	//队列存下标,队列内元素维持单调递减
	var queue []int
	update := func(i int) {
		for len(queue) > 0 && nums[i] >= nums[queue[len(queue)-1]] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
	}
	for i := 0; i < k; i++ {
		update(i)
	}
	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[queue[0]]
	for i := k; i < n; i++ {
		update(i)
		for queue[0] <= i-k {
			queue = queue[1:]
		}
		ans = append(ans, nums[queue[0]])
	}
	return ans
}

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	fmt.Println(maxSlidingWindow(nums, k))
	nums = []int{1}
	k = 1
	fmt.Println(maxSlidingWindow(nums, k))
}
