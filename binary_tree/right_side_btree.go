package binarytree

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var q []*TreeNode
	var res []int
	q = append(q, root)
	for len(q) > 0 {
		res = append(res, q[0].Val)
		cl := len(q)
		for i := 0; i < cl; i++ {
			r1 := q[0]
			if r1.Right != nil {
				q = append(q, r1.Right)
			}
			if r1.Left != nil {
				q = append(q, r1.Left)
			}
			q = q[1:]
		}
	}
	return res
}
