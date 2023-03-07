package main

import "fmt"

func trap(height []int) int {
	sum := 0
	var stack []int
	for i, v := range height {
		fmt.Println("-----------stack:", stack)
		for len(stack) > 0 && v > stack[len(stack)-1] {
			fmt.Println("stack:", stack)
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
