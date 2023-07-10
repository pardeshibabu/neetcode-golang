package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	count := 0
	temp := head
	for temp != nil {
		count++
		temp = temp.Next
	}
	temp = head
	if count > 1 && n == count {
		temp = temp.Next
		head = temp
		return head
	}

	if count == 1 && n == 1 {
		return nil
	}
	k := 0
	for temp != nil {
		if k+1 == count-n {
			next := temp.Next.Next
			temp.Next = next
			temp = nil
			break
		}
		k++
		temp = temp.Next
	}
	return head
}
