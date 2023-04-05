package main

import "fmt"

func isAnagram(s string, t string) bool {
	m, n := len(s), len(t)
	if m != n {
		return false
	}
	var a, b [26]int
	for i := 0; i < m; i++ {
		a[s[i]-'a']++
		b[t[i]-'a']++
	}
	return a == b
}

func main() {
	s := "anagram"
	t := "nagaram"
	fmt.Println(isAnagram(s, t))
	s = "rat"
	t = "car"
	fmt.Println(isAnagram(s, t))
}
