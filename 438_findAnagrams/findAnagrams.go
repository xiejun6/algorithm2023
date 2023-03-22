package main

import "fmt"

func findAnagrams(s string, p string) []int {
	m, n := len(s), len(p)
	if m < n {
		return nil
	}
	var sCount, pCount [26]int
	for i := 0; i < n; i++ {
		pCount[p[i]-'a']++
		sCount[s[i]-'a']++
	}
	var ans []int
	if sCount == pCount {
		ans = append(ans, 0)
	}
	for i := n; i < m; i++ {
		sCount[s[i]-'a']++
		sCount[s[i-n]-'a']--
		if sCount == pCount {
			ans = append(ans, i-n+1)
		}
	}
	return ans
}

func main() {
	s := "cbaebabacd"
	p := "abc"
	fmt.Println(findAnagrams(s, p))
	s = "abab"
	p = "ab"
	fmt.Println(findAnagrams(s, p))
}
