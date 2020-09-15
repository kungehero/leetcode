/*
 * @lc app=leetcode.cn id=129 lang=golang
 *
 * [129] 求根到叶子节点数字之和
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {

	return helper(root, 0)
}

func helper(root *TreeNode, res int) int {
	tem := res*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return tem
	}
	return helper(root.Left, tem) + helper(root.Right, tem)
}

// @lc code=end

