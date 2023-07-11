package linkedlist

// Definition for a Node.
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// this solution is optimized but can't modify original LL
// func copyRandomList(head *Node) *Node {
// 	temp := head
// 	// m := make(map[*Node]*Node)
// 	for temp != nil {
// 		t := &Node{
// 			Val:    temp.Val,
// 			Next:   nil,
// 			Random: &Node{},
// 		}
// 		next := temp.Next
// 		temp.Next = t
// 		t.Next = next
// 		temp = next
// 	}
// 	temp = head
// 	for temp != nil {
//         if temp.Random != nil{
// 		    temp.Next.Random = temp.Random.Next
//         }
// 		temp = temp.Next.Next
// 	}
// 	cHead := head.Next
// 	cTemp := cHead
// 	for cTemp != nil && cTemp.Next != nil{
//         n := cTemp.Next
// 		cTemp.Next = n.Next
//         cTemp = n
// 	}
// 	return cHead
// }

func copyRandomList(head *Node) *Node {
	// prev := head
	temp := head
	m := make(map[*Node]*Node)
	for temp != nil {
		t := &Node{
			Val:    temp.Val,
			Next:   nil,
			Random: nil,
		}
		m[temp] = t
		temp = temp.Next
	}
	temp = head
	for temp != nil {
		random := temp.Random
		m[temp].Random = m[random]
		m[temp].Next = m[temp.Next]
		temp = temp.Next
	}
	return m[head]
}
