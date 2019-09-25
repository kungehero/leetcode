/*
 * @lc app=leetcode.cn id=237 lang=golang
 *
 * [237] 删除链表中的节点
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
	*node = *node.Next
}
