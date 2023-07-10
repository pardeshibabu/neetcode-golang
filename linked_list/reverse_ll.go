package linkedlist

// * Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	next := &ListNode{}
	prev := &ListNode{}
	prev = nil
	for head != nil {
		next = head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}
