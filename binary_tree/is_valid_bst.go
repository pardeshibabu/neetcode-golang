package binarytree

import "math"

func isValidBSTUtil(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Val <= min || root.Val >= max {
		return false
	}
	return isValidBSTUtil(root.Left, min, root.Val) && isValidBSTUtil(root.Right, root.Val, max)
}

func isValidBST(root *TreeNode) bool {
	return isValidBSTUtil(root, math.MinInt, math.MaxInt)
}
