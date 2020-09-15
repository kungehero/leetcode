/*
 * @lc app=leetcode.cn id=230 lang=golang
 *
 * [230] 二叉搜索树中第K小的元素
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
func kthSmallest(root *TreeNode, k int) int {

	//中序遍历
	stack := []*TreeNode{}
	count := 0
	for len(stack) != 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		if len(stack) != 0 {
			count++
			if count == k {
				return stack[len(stack)-1].Val
			}
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			root = root.Right
		}
	}
	return 0
}

// @lc code=end

