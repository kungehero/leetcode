/*
 * @lc app=leetcode.cn id=111 lang=golang
 *
 * [111] 二叉树的最小深度
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
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}
	min_Depth := math.MaxInt32
	if root.Left != nil {
		min_Depth = min(minDepth(root.Left), min_Depth)
	}
	if root.Right != nil {
		min_Depth = min(minDepth(root.Right), min_Depth)
	}
	return min_Depth + 1
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end

