/*
 * @lc app=leetcode.cn id=114 lang=golang
 *
 * [114] 二叉树展开为链表
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
func flatten(root *TreeNode) {
	list := []*TreeNode{}
	stack := []*TreeNode{}
	stack = append(stack, root)
	for root != nil || len(stack) != 0 {
		length := len(stack)
		for length > 0 {
			list = append(list, root)
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		root = root.Right
		stack = stack[:len(stack)-1]
	}

	for i := 1; i < len(list); i++ {
		pre, cur := list[i-1], list[i]
		pre.Left, cur.Rigth = nil, cur
	}
}

// @lc code=end

