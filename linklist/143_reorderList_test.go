package linklist

import "testing"

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}
	//1、分割链表成两半 (长度奇数的时候前半段多一个)
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	newHead := slow.Next
	slow.Next = nil
	//PrintList(newHead)
	//2、反转链表
	newHead = doReverse(newHead)
	//3、合并链表
	mergeTwoList(head, newHead)
}

func doReverse(head *ListNode) *ListNode {
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

func mergeTwoList(head1, head2 *ListNode) {
	p, q := head1, head2
	for p != nil && q != nil {
		n1 := p.Next
		n2 := q.Next
		p.Next = q
		q.Next = n1
		p = n1
		q = n2
	}
}

func Test_reorderList(t *testing.T) {
	h := MakeList([]int{1, 2, 3, 4, 5})
	reorderList(h)
	PrintList(h)

	h = MakeList([]int{1, 2, 3, 4})
	reorderList(h)
	PrintList(h)
}
