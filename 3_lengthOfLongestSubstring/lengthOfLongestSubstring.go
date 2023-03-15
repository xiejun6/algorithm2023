package main

import "fmt"

// for 循环枚举end位置,再判断左边界
func lengthOfLongestSubstring(s string) int {
	posMap := make(map[byte]int)
	ans := 0
	start := 0
	for end := 0; end < len(s); end++ {
		index, ok := posMap[s[end]]
		posMap[s[end]] = end
		if ok {
			//这一句很重要 //例子："tmmzuxt"
			start = max(start, index+1)
		}
		//最后一个t的时候,第一个t还在map中，所以不管存不存在，都要算一下
		//start往后移了，就不能往回退了
		//fmt.Println("start:", start, "end:", end)
		ans = max(ans, end-start+1)
	}
	return ans
}
func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// 循环左边界
func lengthOfLongestSubstring2(s string) int {
	n := len(s)
	m := make(map[byte]int)
	j := 0
	ans := 0
	for i := 0; i < n; i++ {
		if i > 0 {
			delete(m, s[i-1])
		}
		for j < n && m[s[j]] == 0 {
			m[s[j]]++
			j++
		}
		ans = max(ans, j-i)
		if j == n {
			break
		}
	}
	return ans
}
func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s)) //3
	s = "bbbbb"
	fmt.Println(lengthOfLongestSubstring(s)) //1
	s = "pwwkew"
	fmt.Println(lengthOfLongestSubstring(s)) //3
	s = "tmmzuxt"
	fmt.Println(lengthOfLongestSubstring(s)) //5
	fmt.Println("==========================")
	s = "abcabcbb"
	fmt.Println(lengthOfLongestSubstring2(s)) //3
	s = "bbbbb"
	fmt.Println(lengthOfLongestSubstring2(s)) //1
	s = "pwwkew"
	fmt.Println(lengthOfLongestSubstring2(s)) //3
	s = "tmmzuxt"
	fmt.Println(lengthOfLongestSubstring2(s)) //5
}
