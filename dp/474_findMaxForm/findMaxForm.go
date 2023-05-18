package main

import "fmt"

// 本质是0-1背包，但是容量有2个维度，0个数量，1的数量限制
// dp[i][j][k] //前i物品放入0容量为j、1容量为k的背包
// dp[i][j][k] = max(d[i-1][j][k], dp[i-1][j-zeros][k-ones]+1)
// 利用滚动数组降维,去掉第1维
// d[j][k] = max(dp[j][k], d[j-zeros][k-ones]+1)
// 降维必须倒序遍历，不然会发生覆盖，因为依赖上层的数据
func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for _, str := range strs {
		zeros, ones := 0, 0
		for _, v := range str {
			if v == '0' {
				zeros++
			} else {
				ones++
			}
		}
		for i := m; i >= zeros; i-- {
			for j := n; j >= ones; j-- {
				dp[i][j] = max(dp[i][j], dp[i-zeros][j-ones]+1)
			}
		}
	}
	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	strs := []string{"10", "0001", "111001", "1", "0"}
	m := 5
	n := 3
	fmt.Println(findMaxForm(strs, m, n))
	strs = []string{"10", "0", "1"}
	m = 1
	n = 1
	fmt.Println(findMaxForm(strs, m, n))
}
