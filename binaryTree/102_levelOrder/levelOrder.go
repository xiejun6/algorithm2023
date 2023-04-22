package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var ans [][]int
	q := []*TreeNode{root}
	for q != nil {
		n := len(q)
		leveNums := make([]int, 0, n)
		var q1 []*TreeNode
		for i := 0; i < n; i++ {
			leveNums = append(leveNums, q[i].Val)
			if q[i].Left != nil {
				q1 = append(q1, q[i].Left)
			}
			if q[i].Right != nil {
				q1 = append(q1, q[i].Right)
			}
		}
		ans = append(ans, leveNums)
		q = q1
	}
	return ans
}

func main() {

}
