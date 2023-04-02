package main

import "fmt"

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	minRow, maxRow := 0, n-1
	minCol, maxCol := 0, n-1
	count := n * n
	num := 1
	for num <= count {
		//上边
		for i := minCol; i <= maxCol; i++ {
			matrix[minRow][i] = num
			num++
		}
		minRow++
		//右边
		for i := minRow; i <= maxRow; i++ {
			matrix[i][maxCol] = num
			num++
		}
		maxCol--
		//下边
		for i := maxCol; i >= minCol; i-- {
			matrix[maxRow][i] = num
			num++
		}
		maxRow--
		//左边
		for i := maxRow; i >= minRow; i-- {
			matrix[i][minCol] = num
			num++
		}
		minCol++
	}
	return matrix
}

func main() {
	fmt.Println(generateMatrix(3))
	fmt.Println(generateMatrix(1))
}
