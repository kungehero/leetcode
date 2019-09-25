/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package leet

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSnameTree(root.Left, root.Right)
}

// TODO:再理解一遍
func isSnameTree(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	// FIXME:节点命名修改
	if isSnameTree(left.Left, right.Right) && isSnameTree(left.Right, right.Left) {
		return true
	}
	return false
}
