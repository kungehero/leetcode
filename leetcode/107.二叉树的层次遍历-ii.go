/*
 * @lc app=leetcode.cn id=107 lang=golang
 *
 * [107] 二叉树的层次遍历 II
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
func levelOrderBottom(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return [][]int{}
	}
	stack := []*TreeNode{}
	stack = append(stack, root)

	for len(stack) != 0 {
		path := []int{}
		length := len(stack)

		for length > 0 {
			if stack[0].Left != nil {
				stack = append(stack, stack[0].Left)

			}
			if stack[0].Right != nil {
				stack = append(stack, stack[0].Right)

			}
			path = append(path, stack[0].Val)
			length--
			stack = stack[1:]
		}
		res = append(res, path)
	}
	return reverse(res)
}

func reverse(res [][]int) [][]int {
	i, j := 0, len(res)-1
	for i < j {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	return res
}

// @lc code=end