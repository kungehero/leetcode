/*
 * @lc app=leetcode.cn id=226 lang=golang
 *
 * [226] 翻转二叉树
 */

// @lc code=start
/**
 * Definition for a binary tree node.

 */

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	queue := list.New()
	queue.PushFront(root)

	for queue.Len() > 0 {
		count := queue.Len()
		for count > 0 {
			node := queue.Front()
			nextnode := node.Value.(*TreeNode)

			nextnode.Left, nextnode.Right = nextnode.Right, nextnode.Left

			if nextnode.Left != nil {
				queue.PushBack(nextnode.Left)
			}
			if nextnode.Right != nil {
				queue.PushBack(nextnode.Right)
			}		  
			queue.Remove(node)
			count--
		}
	}
	return root
}

// @lc code=end
