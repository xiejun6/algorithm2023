package main

import (
	"fmt"
	"math"
)

func myAtoi(s string) int {
	size := len(s)
	start := 0
	//1 丢弃前导空格
	for start < size && s[start] == ' ' {
		start++
	}
	if start >= size {
		return 0
	}
	//2 检查下一个字符是正号还是负号
	flag := 0 //正
	if s[start] == '-' {
		flag = -1
		start++
	} else if s[start] == '+' {
		start++
	}
	if start >= size {
		return 0
	}
	//3 读入下一个字符，直到到达下一个非数字字符或到达输入的结尾
	var res int
	maxVal := math.MaxInt32
	for i := start; i < size; i++ {
		if !(s[i] >= '0' && s[i] <= '9') {
			break
		}
		v := int(s[i] - '0')
		res = res*10 + v
		if res > maxVal {
			if flag == 0 {
				res = maxVal
			} else {
				res = maxVal + 1
			}
		}
	}
	if flag == 0 {
		return res
	}
	return -res
}

func main() {
	s := "42"
	fmt.Println(myAtoi(s))
	s = "   -42"
	fmt.Println(myAtoi(s))
	s = "4193 with words"
	fmt.Println(myAtoi(s))
}
