package main

import "fmt"

func generateParenthesis(n int) []string {
	if n == 0 {
		return nil
	}
	var ans []string
	var dfs func(tmp string, left int, right int)
	dfs = func(tmp string, left int, right int) {
		if left == n && right == n {
			ans = append(ans, tmp)
			return
		}
		if left < n {
			dfs(tmp+"(", left+1, right)
		}
		//如果在本层左括号比右括号多,那么也可以选右括号
		if left > right {
			dfs(tmp+")", left, right+1)
		}
	}
	dfs("", 0, 0)
	return ans
}

func main() {
	fmt.Println(generateParenthesis(3))
	fmt.Println(generateParenthesis(2))
	fmt.Println(generateParenthesis(1))
}
