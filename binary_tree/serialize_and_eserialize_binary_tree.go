package binarytree

import (
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	// if root == nil {
	// 	return ""
	// }
	// var res string
	// var q []*TreeNode
	// q = append(q, root)
	// for len(q) > 0 {
	// 	r1 := q[0]
	// 	q = q[1:]
	// 	if r1 == nil {
	// 		res += "#,"
	// 	} else {
	// 		res += string(r1.Val) + ","
	// 	}
	// 	if r1 != nil {
	// 		q = append(q, r1.Left)
	// 		q = append(q, r1.Right)
	// 	}
	// }
	var res string
	var dfs func(root *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			res += "#" + ","
			return
		}
		res = res + strconv.Itoa(node.Val) + ","
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return res
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	data = strings.Trim(data, ",")
	d := strings.Split(data, ",")
	i := 0

	var dfs func() *TreeNode

	dfs = func() *TreeNode {
		if d[i] == "#" {
			i += 1
			return nil
		}
		val, _ := strconv.Atoi(d[i])
		node := &TreeNode{Val: val}
		node.Left = dfs()
		node.Right = dfs()
		return node
	}
	return dfs()
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
