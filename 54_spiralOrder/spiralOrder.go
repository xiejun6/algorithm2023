package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	minRow, maxRow := 0, m-1
	minCol, maxCol := 0, n-1

	count, maxCount := 0, m*n
	ans := make([]int, 0, maxCount)
	for {
		//上
		for i := minCol; i <= maxCol; i++ {
			ans = append(ans, matrix[minRow][i])
			count++
		}
		if count >= maxCount {
			break
		}
		minRow++
		//右边
		for i := minRow; i <= maxRow; i++ {
			ans = append(ans, matrix[i][maxCol])
			count++
		}
		if count >= maxCount {
			break
		}
		maxCol--
		//下边
		for i := maxCol; i >= minCol; i-- {
			ans = append(ans, matrix[maxRow][i])
			count++
		}
		if count >= maxCount {
			break
		}
		maxRow--
		//左边
		for i := maxRow; i >= minRow; i-- {
			ans = append(ans, matrix[i][minCol])
			count++
		}
		if count >= maxCount {
			break
		}
		minCol++
	}
	return ans
}

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(spiralOrder(matrix))
	matrix = [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println(spiralOrder(matrix))
}
