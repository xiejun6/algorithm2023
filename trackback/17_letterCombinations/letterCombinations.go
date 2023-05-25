package main

import "fmt"

var arr = [10]string{
	"",
	"",
	"abc",
	"def",
	"ghi",
	"jkl",
	"mno",
	"pqrs",
	"tuv",
	"wxyz",
}

func letterCombinations(digits string) []string {
	//这个判断没写提交会报错
	if len(digits) == 0 {
		return nil
	}
	var ans []string
	var b []byte
	n := len(digits)
	var dfs func(startIndex int)
	dfs = func(startIndex int) {
		if startIndex == n {
			tmp := string(append([]byte(nil), b...))
			ans = append(ans, tmp)
			return
		}
		str := arr[digits[startIndex]-'0']
		for i := 0; i < len(str); i++ {
			b = append(b, str[i])
			dfs(startIndex + 1)
			b = b[:len(b)-1]
		}
	}
	dfs(0)
	return ans
}

func main() {
	digits := "23"
	fmt.Println(letterCombinations(digits))
	digits = "2"
	fmt.Println(letterCombinations(digits))
	digits = ""
	fmt.Println(letterCombinations(digits))
}
