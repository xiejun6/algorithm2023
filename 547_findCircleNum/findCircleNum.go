package main

import "fmt"

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	visited := make([]bool, n)
	ans := 0
	var dfs func(from int)
	dfs = func(from int) {
		visited[from] = true
		for to, v := range isConnected[from] {
			if v == 1 && !visited[to] {
				dfs(to)
			}
		}
	}
	for i := 0; i < n; i++ {
		if !visited[i] {
			ans++
			dfs(i)
		}
	}
	return ans
}

func main() {
	isConnected := [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}
	fmt.Println(findCircleNum(isConnected))
	isConnected = [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	fmt.Println(findCircleNum(isConnected))
}
