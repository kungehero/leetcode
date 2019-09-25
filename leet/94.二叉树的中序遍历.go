import "container/list"

/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	list := list.New()
	for root != nil || list.Len() != 0 {
		for root != nil {
			list.PushBack(root)
			root = root.Left
		}
		if list.Len() != 0 {
			v := list.Back()
			root = v.Value.(*TreeNode)
			res = append(res, root.Val)
			root = root.Right
			list.Remove(v)
		}
	}
	return res
}
