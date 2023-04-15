package linklist

import "testing"

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummyHead := &ListNode{Next: head}
	prev := dummyHead
	p := head
	for i := 1; i < left; i++ {
		prev = p
		p = p.Next
	}
	leftHead := prev.Next
	q := leftHead

	for i := left; i < right; i++ {
		q = q.Next
	}
	rightNode := q.Next
	q.Next = nil
	newHead := reverse(leftHead)
	prev.Next = newHead
	leftHead.Next = rightNode
	return dummyHead.Next
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	p := head
	for p != nil {
		q := p.Next
		p.Next = prev
		prev = p
		p = q
	}
	return prev
}

func TestReverseBetween(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	head := MakeList(data)
	newHead := reverseBetween(head, 2, 4)
	PrintList(newHead)
}
