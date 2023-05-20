package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	size := len(s)
	dp := make([]bool, size+1)
	dp[0] = true
	for i := 1; i <= size; i++ {
		for _, v := range wordDict {
			n1, n2 := len(s[:i]), len(v)
			if n1 >= n2 && s[i-n2:i] == v {
				dp[i] = dp[i] || dp[i-n2]
			}
		}
	}
	return dp[size]
}

func main() {
	s := "leetcode"
	wordDict := []string{"leet", "code"}
	fmt.Println(wordBreak(s, wordDict))
	s = "applepenapple"
	wordDict = []string{"apple", "pen"}
	fmt.Println(wordBreak(s, wordDict))
	s = "catsandog"
	wordDict = []string{"cats", "dog", "sand", "and", "cat"}
	fmt.Println(wordBreak(s, wordDict))
}
