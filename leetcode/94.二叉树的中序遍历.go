/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
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
func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	stack := list.New()
	for root != nil || stack.Len() != 0 {
		for root != nil {
			stack.PushBack(root)
			root = root.Left
		}
		if stack.Len() != 0 {
			v := stack.Back()
			root = v.Value.(*TreeNode)
			res = append(res, root.Val)
			root = root.Right
			stack.Remove(v)
		}
	}
	return res
}

// @lc code=end
