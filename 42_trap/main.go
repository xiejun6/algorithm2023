package main

import "fmt"

func trap(height []int) int {
	sum := 0
	var stack []int //单调递减栈，存下标
	for i, v := range height {
		for len(stack) > 0 && v > height[stack[len(stack)-1]] {
			topIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			leftIndex := stack[len(stack)-1]
			h := min(height[leftIndex], v) - height[topIndex]
			sum += (i - leftIndex - 1) * h
		}

		stack = append(stack, i)
	}
	return sum
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(height))
}
