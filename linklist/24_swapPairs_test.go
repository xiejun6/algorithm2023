package linklist

import "testing"

func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{Next: head}
	p := head
	prev := dummyHead
	for p != nil && p.Next != nil {
		q := p.Next
		p.Next = q.Next
		q.Next = p
		prev.Next = q
		prev = p
		p = p.Next
	}
	return dummyHead.Next
}

func Test_swapPairs(t *testing.T) {
	data := []int{1, 2, 3, 4}
	head := MakeList(data)
	newHead := swapPairs(head)
	PrintList(newHead)

	data = []int{1}
	head = MakeList(data)
	newHead = swapPairs(head)
	PrintList(newHead)
}
