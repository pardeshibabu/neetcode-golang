package binarytree

func IsSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
}

//	func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
//		if root == nil && subRoot == nil {
//			return true
//		}
//		if root == nil || subRoot == nil {
//			return false
//		}
//		if IsSameTree(root, subRoot) {
//			return true
//		}
//		return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
//	}
func FindMathedNode(root, subRoot *TreeNode, matchedNode []*TreeNode) []*TreeNode {
	if root == nil || subRoot == nil {
		return matchedNode
	}
	if root.Val == subRoot.Val {
		matchedNode = append(matchedNode, root)
	}
	FindMathedNode(root.Left, subRoot, matchedNode)
	FindMathedNode(root.Right, subRoot, matchedNode)
	return matchedNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	var matchedNode []*TreeNode
	temp := root
	matchedNode = FindMathedNode(temp, subRoot, matchedNode)
	for i := 0; i < len(matchedNode); i++ {
		if IsSameTree(matchedNode[i], subRoot) {
			return true
		}
	}
	return false
}
