package binarytree

func kthSmallestUtil(root *TreeNode, k int, c *int) int {
	if root == nil {
		return -1
	}
	res := kthSmallestUtil(root.Left, k, c)
	if res != -1 {
		return res
	}
	*c++
	if *c == k {
		return root.Val
	}
	return kthSmallestUtil(root.Right, k, c)
}

func kthSmallest(root *TreeNode, k int) int {
	c := 0
	return kthSmallestUtil(root, k, &c)
}
