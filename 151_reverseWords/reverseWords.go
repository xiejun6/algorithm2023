package main

import "fmt"

/*
输入：s = "the sky is blue"
输出："blue is sky the"
*/
func reverseWords(s string) string {
	ans := ""
	start := -1
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' && start < 0 {
			start = i
		}
		if s[i] == ' ' && start >= 0 {
			if ans == "" {
				ans = s[start:i] + ans
			} else {
				ans = s[start:i] + " " + ans
			}
			start = -1
		}
	}
	if start >= 0 {
		if ans == "" {
			ans = s[start:] + ans
		} else {
			ans = s[start:] + " " + ans
		}
	}
	return ans
}

func main() {
	fmt.Println(reverseWords("the sky is blue"))
	s := "  hello world  "
	fmt.Println(reverseWords(s))
	s = "a good   example"
	fmt.Println(reverseWords(s))
	s = "EPY2giL"
	fmt.Println(reverseWords(s))
}
