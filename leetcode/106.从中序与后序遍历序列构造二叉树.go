/*
 * @lc app=leetcode.cn id=106 lang=golang
 *
 * [106] 从中序与后序遍历序列构造二叉树
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
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{postorder[len(postorder)-1], nil, nil}
	r := len(postorder) - 1
	i := 0
	for i = 0; i < len(inorder); i++ {
		if inorder[i] == postorder[r] {
			break
		}
	}
	root.Left = buildTree(inorder[:i], postorder[:i])
	root.Right = buildTree(inorder[i+1:], postorder[i:len(postorder)-1])
	return root
}

// @lc code=end
