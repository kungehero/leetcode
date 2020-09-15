/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层次遍历
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
	res := [][]int{}
	stack := []*TreeNode{}
	stack = append(stack, root)

	if root == nil {
		return res
	}
	for len(stack) != 0 {
		length := len(stack)
		path := []int{}
		for length > 0 {
			path = append(path, stack[0].Val)
			if stack[0].Left != nil {
				stack = append(stack, stack[0].Left)
			}
			if stack[0].Right != nil {
				stack = append(stack, stack[0].Right)
			}
			length--
			stack = stack[1:]
		}
		res = append(res, path)
	}
	return res
}

// @lc code=end
