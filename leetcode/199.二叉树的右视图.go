/*
 * @lc app=leetcode.cn id=199 lang=golang
 *
 * [199] 二叉树的右视图
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
func rightSideView(root *TreeNode) []int {
	var (
		res   []int
		queue []*TreeNode
	)
	if root == nil {
		return res
	}
	queue = append(queue, root)
	for len(queue) != 0 {
		length := len(queue)
		res = append(res, queue[len(queue)-1].Val)
		for length != 0 {
			root = queue[0]
			queue = queue[1:]
			if root.Left != nil {
				queue = append(queue, root.Left)
			}
			if root.Right != nil {
				queue = append(queue, root.Right)
			}
			length--
		}
	}
	return res
}

// @lc code=end

