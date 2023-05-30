package main

import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	m, n := len(g), len(s)
	i, j := 0, 0
	count := 0
	for i < m && j < n {
		if s[j] >= g[i] {
			count++
			i++
			j++
		} else {
			j++
		}
	}
	return count
}

func main() {
	g := []int{1, 2, 3}
	s := []int{1, 1}
	fmt.Println(findContentChildren(g, s))
	g = []int{1, 2}
	s = []int{1, 2, 3}
	fmt.Println(findContentChildren(g, s))
}
