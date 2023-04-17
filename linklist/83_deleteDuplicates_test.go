package linklist

import (
	"fmt"
	"testing"
)

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := head.Next
	prev := head
	for p != nil {
		if p.Val != prev.Val {
			prev.Next = p
			prev = p
		} else {
			prev.Next = nil
		}
		p = p.Next
	}
	return head
}

func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	cur := head
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

func TestDeleteDuplicates(t *testing.T) {
	data := []int{1, 1, 2}
	head := MakeList(data)
	newHead := deleteDuplicates(head)
	PrintList(newHead)

	data = []int{1, 1, 2, 3, 3}
	head = MakeList(data)
	newHead = deleteDuplicates(head)
	PrintList(newHead)

	fmt.Println("=====================")
	data = []int{1, 1, 2}
	head = MakeList(data)
	newHead = deleteDuplicates2(head)
	PrintList(newHead)

	data = []int{1, 1, 2, 3, 3}
	head = MakeList(data)
	newHead = deleteDuplicates2(head)
	PrintList(newHead)
}
