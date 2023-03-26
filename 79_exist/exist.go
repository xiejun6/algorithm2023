package main

import "fmt"

var dirList = [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

func dfs(board [][]byte, word string, visited [][]bool, m, n, i, j, index int) bool {
	if i < 0 || i >= m || j < 0 || j >= n || visited[i][j] {
		return false
	}
	if board[i][j] != word[index] {
		return false
	}
	visited[i][j] = true
	index++
	defer func() {
		visited[i][j] = false
	}()
	if index == len(word) {
		return true
	}
	for _, dir := range dirList {
		x, y := i+dir[0], j+dir[1]
		if dfs(board, word, visited, m, n, x, y, index) {
			return true
		}
	}
	return false
}

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])

	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(board, word, visited, m, n, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func main() {
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	word := "ABCCED"
	fmt.Println(exist(board, word))
	word = "SEE"
	fmt.Println(exist(board, word))
	word = "ABCB"
	fmt.Println(exist(board, word))
}
