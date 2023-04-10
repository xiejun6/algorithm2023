package linklist

import (
	"testing"
)

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	prev := dummyHead
	carry := 0
	for l1 != nil || l2 != nil {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		node := &ListNode{Val: sum % 10}
		prev.Next = node
		prev = node
		carry = sum / 10
	}
	if carry > 0 {
		node := &ListNode{Val: carry}
		prev.Next = node
	}
	return dummyHead.Next
}

func TestAddTwoNum(t *testing.T) {
	a := []int{2, 4, 3}
	b := []int{5, 6, 4}
	l1 := MakeList(a)
	l2 := MakeList(b)
	l3 := addTwoNumbers(l1, l2)
	PrintList(l3)

	a = []int{9, 9, 9, 9, 9, 9, 9}
	b = []int{9, 9, 9, 9}
	l1 = MakeList(a)
	l2 = MakeList(b)
	l3 = addTwoNumbers(l1, l2)
	PrintList(l3)
}
