/*
 * @lc app=leetcode.cn id=103 lang=golang
 *
 * [103] 二叉树的锯齿形层次遍历
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
 func zigzagLevelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	isLeftStart := true
	for len(queue) > 0 {
		l := len(queue)
		ans := make([]int, l)
		for i := 0; i < l; i++ {
			node := queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if isLeftStart {
				ans[i] = node.Val
			} else {
				ans[l-i-1] = node.Val
			}
		}
		res = append(res, ans)
		isLeftStart = !isLeftStart
		queue = queue[l:]
	}
	return res
}
// @lc code=end

