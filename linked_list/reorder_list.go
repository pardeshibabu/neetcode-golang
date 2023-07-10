package linkedlist

func ReorderList(head *ListNode) {
	temp := head
	slow := temp
	fast := temp
	// finding middle through two pointers
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	newHead := slow.Next
	newHead = reverse(slow.Next)
	slow.Next = nil
	for temp != nil && newHead != nil {
		next := temp.Next
		newHeadNext := newHead.Next
		temp.Next = newHead
		newHead.Next = next
		newHead = newHeadNext
		temp = next
	}
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
