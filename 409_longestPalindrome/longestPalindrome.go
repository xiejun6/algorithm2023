package main

import "fmt"

// 用数组表示map
func longestPalindrome(s string) int {
	var count [128]int
	single := 0
	ans := 0
	for i := 0; i < len(s); i++ {
		v := int(s[i])
		count[v]++
		if count[v] == 2 {
			single--
			ans += 2
			count[v] = 0
		} else {
			single++
		}
	}
	if single > 0 {
		ans++
	}
	return ans
}

func main() {
	s := "abccccdd"
	fmt.Println(longestPalindrome(s))
	s = "a"
	fmt.Println(longestPalindrome(s))
	s = "aaaaaccc"
	fmt.Println(longestPalindrome(s))
}
