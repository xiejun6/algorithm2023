package main

import "fmt"

// 单调栈，栈内元素从栈底到栈顶递增（非严格，可能出现等于的情况）
// 当当前要进栈的元素小于栈顶元素的时候，说明栈内元素有些作为高度的面积可以确定了
// 左右边界加入哨兵
// 和42题接雨水有点相似
func largestRectangleArea(heights []int) int {
	h := make([]int, 0, len(heights)+2)
	h = append([]int{0}, heights...)
	h = append(h, 0)
	var stack []int
	stack = append(stack, 0)
	ans := 0
	for i := 1; i < len(h); i++ {
		//fmt.Println("stack:", stack, "i=", i)
		for len(stack) > 0 && h[i] < h[stack[len(stack)-1]] {
			topIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			leftIndex := stack[len(stack)-1]
			if h[leftIndex] < h[topIndex] {
				curWidth := i - leftIndex - 1
				tmp := curWidth * h[topIndex]
				//fmt.Println("...", "leftIndex:", leftIndex, " rightIndex:", i, " topval:", h[topIndex])
				if tmp > ans {
					ans = tmp
				}
			}
		}
		stack = append(stack, i)
	}
	return ans
}

func main() {
	heights := []int{2, 1, 5, 6, 2, 3}
	fmt.Println(largestRectangleArea(heights))
	heights = []int{2, 4}
	fmt.Println(largestRectangleArea(heights))
}
