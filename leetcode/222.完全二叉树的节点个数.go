/*
 * @lc app=leetcode.cn id=222 lang=golang
 *
 * [222] 完全二叉树的节点个数
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
func countNodes(root *TreeNode) int {
	count := 0
	queue := []*TreeNode{}
	if root != nil {
		queue = append(queue, root)
	}

	for len(queue) != 0 {
		length := len(queue)

		for length > 0 {
			node := queue[0]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)

			}
			count++
			queue = queue[1:]
			length--
		}
	}
	return count
}

// @lc code=end

