package linklist

import (
	"testing"
)

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

// cur节点一点一直保持不变,prev一直保持不变
// 把cur.next的节点插入到prev的next
// 更新cur.next
func reverseBetween2(head *ListNode, left int, right int) *ListNode {
	dummyHead := &ListNode{Next: head}
	prev := dummyHead
	for i := 1; i < left; i++ {
		prev = prev.Next
	}
	//需要翻转的第一个节点(cur节点一直没变过)
	cur := prev.Next
	for i := left; i < right; i++ {
		nxt := cur.Next
		cur.Next = nxt.Next
		nxt.Next = prev.Next
		prev.Next = nxt
	}
	return dummyHead.Next
}

func TestReverseBetween(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	head := MakeList(data)
	newHead := reverseBetween(head, 2, 4)
	PrintList(newHead)

	data = []int{1, 2, 3, 4, 5}
	head = MakeList(data)
	newHead = reverseBetween2(head, 2, 4)
	PrintList(newHead)
}
