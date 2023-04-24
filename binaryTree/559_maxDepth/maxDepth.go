package _59_maxDepth

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	ans := 0
	for _, child := range root.Children {
		deep := maxDepth(child)
		if ans < deep {
			ans = deep
		}
	}
	return ans + 1
}
