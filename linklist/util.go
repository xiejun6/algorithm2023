package linklist

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func MakeList(data []int) *ListNode {

	head := &ListNode{}
	p := head
	for _, v := range data {
		node := &ListNode{Val: v}
		p.Next = node
		p = node
	}
	return head.Next
}

func PrintList(head *ListNode) {
	p := head
	for p != nil {
		fmt.Print(p.Val, " ")
		p = p.Next
	}
	fmt.Println()
}
