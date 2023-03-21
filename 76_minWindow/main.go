package main

import "fmt"

// 遍历枚举右边界，左边界尝试不断收缩
func minWindow(s string, t string) string {
	sm := make(map[byte]int)
	tm := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		tm[t[i]]++
	}
	check := func() bool {
		for b, count := range tm {
			if sm[b] < count {
				return false
			}
		}
		return true
	}
	left := 0
	ans := ""
	for right := 0; right < len(s); right++ {
		sm[s[right]]++
		for left <= right && check() {
			//fmt.Println("left, right:", left, right)
			length := right - left + 1
			if ans == "" || length < len(ans) {
				ans = s[left : right+1]
			}
			sm[s[left]]--
			left++

		}
	}
	return ans
}
func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	fmt.Println(minWindow(s, t))

	s = "aaaaaaaaaaaabbbbbcdd"
	t = "abcdd"
	fmt.Println(minWindow(s, t))
}
