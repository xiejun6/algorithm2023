package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	n := len(s)
	table := make(map[byte]byte)
	exists := make(map[byte]struct{})
	for i := 0; i < n; i++ {
		if b, ok := table[s[i]]; ok {
			if b != t[i] {
				return false
			}
		} else {
			if _, ok = exists[t[i]]; ok {
				return false
			}
			table[s[i]] = t[i]
			exists[t[i]] = struct{}{}
		}
	}
	return true
}

func main() {
	s := "egg"
	t := "add"
	fmt.Println(isIsomorphic(s, t))
	s = "foo"
	t = "bar"
	fmt.Println(isIsomorphic(s, t))
	s = "paper"
	t = "title"
	fmt.Println(isIsomorphic(s, t))
}
