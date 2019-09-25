package leet

import "container/list"

/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}

	list := list.New()
	list.PushBack(root)
	maxDeep := 0
	for {
		length := list.Len()
		if length == 0 {
			break
		}
		for i := 0; i < length; i++ {
			v := list.Front()
			node := (v.Value).(*TreeNode)
			list.Remove(v)
			if node.Left != nil {
				list.PushBack(node.Left)
			}
			if node.Right != nil {
				list.PushBack(node.Right)
			}
		}
		maxDeep++
	}
	return maxDeep
}
