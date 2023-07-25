package binarytree

import (
	"math"
)

func maxPathSumUtil(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	lm := maxPathSumUtil(root.Left, res)
	rm := maxPathSumUtil(root.Right, res)
	if lm < 0 {
		lm = 0
	}
	if rm < 0 {
		rm = 0
	}
	if *res < (lm + rm + root.Val) {
		// fmt.Println(lm, rm, *res, root.Val)
		*res = lm + rm + root.Val
	}
	if lm < rm {
		return root.Val + rm
	}
	return root.Val + lm
}

func maxPathSum(root *TreeNode) int {
	res := math.MinInt
	maxPathSumUtil(root, &res)
	return res
}
