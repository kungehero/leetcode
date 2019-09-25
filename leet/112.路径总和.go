/*
 * @lc app=leetcode.cn id=112 lang=golang
 *
 * [112] 路径总和
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Val == sum && root.Left == nil && root.Right == nil {
		return true
	}
	sum -= root.Val
	return hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum)
}
