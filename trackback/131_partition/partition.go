package main

import "fmt"

func partition(s string) [][]string {
	var ans [][]string
	var path []string
	n := len(s)
	var dfs func(startIndex int)
	dfs = func(startIndex int) {
		if startIndex == n {
			tmp := append([]string{}, path...)
			ans = append(ans, tmp)
			return
		}

		for i := startIndex; i < n; i++ {
			if isPalindrome(s[startIndex : i+1]) {
				path = append(path, s[startIndex:i+1])
				dfs(i + 1)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0)
	return ans
}

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	s := "aab"
	fmt.Println(partition(s))
}
