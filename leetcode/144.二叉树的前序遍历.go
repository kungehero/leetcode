/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] 二叉树的前序遍历
 */

// @lc code=start
/**
 * Definition for a binary tree node.

 */


func preorderTraversal(root *TreeNode) []int {
	res:= []int{}
	stack := list.New()
	if root == nil {
		return res
	}
	for root != nil||stack.Len()!=0 {
		for root != nil {
			res = append(res, root.Val)
			stack.PushBack(root)
			root = root.Left
		}
		if stack.Len() != 0 {
			v := stack.Back()
			root = v.Value.(*TreeNode)
			root = root.Right
			stack.Remove(v)
		}
	}
	return res
}

// @lc code=end
