/*
 * @lc app=leetcode.cn id=437 lang=golang
 *
 * [437] 路径总和 III
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	return paths(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

func paths(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	res := 0
	if root.Val == sum {
		res += 1
	}
	res += paths(root.Left, sum-root.Val)
	res += paths(root.Right, sum-root.Val)
	return res
}
