package binarytree

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var queue []*TreeNode
	var res [][]int
	queue = append(queue, root)
	level := 0
	for len(queue) > 0 {
		cl := len(queue)
		for i := 0; i < cl; i++ {
			r1 := queue[0]
			queue = queue[1:]

			if len(res) <= level {
				res = append(res, []int{r1.Val})
			} else {
				res[level] = append(res[level], r1.Val)
			}
			if r1.Left != nil {
				queue = append(queue, r1.Left)
			}
			if r1.Right != nil {
				queue = append(queue, r1.Right)
			}
		}
		level++
	}
	return res
}
