package linklist

import (
	"testing"
)

func reverseList(head *ListNode) *ListNode {
	var prev, p *ListNode
	p = head
	for p != nil {
		q := p.Next
		p.Next = prev
		prev = p
		p = q
	}
	return prev
}

func reverseList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := reverseList1(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}

func TestReverseList(t *testing.T) {
	head := []int{1, 2, 3, 4, 5}
	h := MakeList(head)
	PrintList(h)
	h1 := reverseList(h)
	PrintList(h1)

	head = []int{1, 2}
	h = MakeList(head)
	PrintList(h)
	h1 = reverseList(h)
	PrintList(h1)

}

func TestReverseList1(t *testing.T) {
	head := []int{1, 2, 3, 4, 5}
	h := MakeList(head)
	PrintList(h)
	h1 := reverseList1(h)
	PrintList(h1)

	head = []int{1, 2}
	h = MakeList(head)
	PrintList(h)
	h1 = reverseList1(h)
	PrintList(h1)

}
