package binarytree

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// if root == nil {
	// 	return nil
	// }
	// if p.Val < root.Val && q.Val < root.Val {
	// 	return lowestCommonAncestor(root.Left, p, q)
	// }
	// if p.Val > root.Val && q.Val > root.Val {
	// 	return lowestCommonAncestor(root.Right, p, q)
	// }
	// return root
	for root != nil {
		if p.Val < root.Val && q.Val < root.Val {
			root = root.Left
			// return lowestCommonAncestor(root.Left, p, q)
		} else if p.Val > root.Val && q.Val > root.Val {
			root = root.Right
			// return lowestCommonAncestor(root.Right, p, q)
		} else {
			break
		}
	}
	return root
}
