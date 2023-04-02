package main

import "fmt"

// 判断：倒序拼出的数字是否等于原数字
func isPalindrome(x int) bool {
	if x < 0 || (x > 0 && x%10 == 0) {
		return false
	}
	ans := 0
	y := x
	for y != 0 {
		rem := y % 10
		ans = ans*10 + rem
		y /= 10
	}
	return ans == x
}

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(10))
	fmt.Println(isPalindrome(-121))
}
