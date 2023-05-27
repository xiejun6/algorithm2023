package main

import (
	"fmt"
	"strconv"
)

func getPermutation(n int, k int) string {
	factor := make([]int, n+1)
	factor[0] = 1
	for i := 1; i < n; i++ {
		factor[i] = factor[i-1] * i
	}
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}
	ans := ""
	k--
	for i := n - 1; i >= 0; i-- {
		index := k / factor[i]
		ans += strconv.Itoa(nums[index])
		nums = append(nums[:index], nums[index+1:]...)
		k = k % factor[i]
	}
	return ans
}

func main() {
	fmt.Println(getPermutation(3, 3))
	fmt.Println(getPermutation(4, 9))
	fmt.Println(getPermutation(3, 1))
}
