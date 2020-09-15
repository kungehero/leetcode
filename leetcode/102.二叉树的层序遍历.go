/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
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
func levelOrder(root *TreeNode) [][]int {

	var res [][]int
	if root == nil {
		return res
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	for len(queue) != 0 {
		length := len(queue)
		path := []int{}
		for length > 0 {
			root = queue[0]
			queue = queue[1:]
			path = append(path, root.Val)
			if root.Left != nil {
				queue = append(queue, root.Left)
			}
			if root.Right != nil {
				queue = append(queue, root.Right)
			}
			length--
		}
		res = append(res, path)
	}
	return res
}

// @lc code=end

