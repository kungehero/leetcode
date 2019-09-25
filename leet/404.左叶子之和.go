/*
 * @lc app=leetcode.cn id=404 lang=golang
 *
 * [404] 左叶子之和
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := 0
	if isLeftLeave(root.Left) {
		sum += root.Left.Val
	} else {
		sum += sumOfLeftLeaves(root.Left)
	}
	sum += sumOfLeftLeaves(root.Right)
	return sum
}

func isLeftLeave(root *TreeNode) bool {
	if root != nil && (root.Left == nil && root.Right == nil) {
		return true
	}
	return false
}
