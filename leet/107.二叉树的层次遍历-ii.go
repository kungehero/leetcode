import "container/list"

/*
 * @lc app=leetcode.cn id=107 lang=golang
 *
 * [107] 二叉树的层次遍历 II
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	list := list.New()
	list.PushBack(root)
	for {
		var tt []int
		len := list.Len()
		if len == 0 {
			break
		}
		for i := 0; i < len; i++ {
			v := list.Front()
			node := (v.Value).(*TreeNode)

			tt = append(tt, node.Val)
			list.Remove(v)
			if node.Left != nil {
				list.PushBack(node.Left)
			}
			if node.Right != nil {
				list.PushBack(node.Right)
			}
		}
		res = append(res, tt)
	}
	return reverse(res)
}

func reverse(param [][]int) [][]int {
	for i, j := 0, len(param)-1; i < j; i, j = i+1, j-1 {
		param[i], param[j] = param[j], param[i]
	}
	return param
}
