package main

import "fmt"

func integerBreak(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 1
	} else if n == 3 {
		return 2
	} else if n == 4 {
		return 4
	}
	val := 1
	for n > 4 {
		val *= 3
		n -= 3
	}
	val *= n
	return val
}

func main() {
	fmt.Println(integerBreak(10))
}
