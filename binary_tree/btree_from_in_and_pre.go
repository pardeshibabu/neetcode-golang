package binarytree

func FindInd(inOrder []int, k int) int {
	for i := 0; i < len(inOrder); i++ {
		if inOrder[i] == k {
			return i
		}
	}
	return 0
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{}
	mid := FindInd(inorder, preorder[0])
	// fmt.Println(mid)
	root.Val = preorder[0]
	root.Left = buildTree(preorder[1:mid+1], inorder[:mid])
	root.Right = buildTree(preorder[mid+1:], inorder[mid+1:])
	return root
}
