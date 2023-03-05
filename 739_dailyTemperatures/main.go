package main

import "fmt"

// 单调栈，栈底最大
func dailyTemperatures(temperatures []int) []int {
	length := len(temperatures)
	res := make([]int, length)
	var stack []int //栈存下标
	for i, v := range temperatures {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < v {
			prevIndex := stack[len(stack)-1]
			res[prevIndex] = i - prevIndex
			stack = stack[:len(stack)-1] //出栈
		}
		stack = append(stack, i) //入栈
	}
	return res
}

func main() {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(temperatures))
	temperatures = []int{30, 40, 50, 60}
	fmt.Println(dailyTemperatures(temperatures))
	temperatures = []int{30, 60, 90}
	fmt.Println(dailyTemperatures(temperatures))
}
