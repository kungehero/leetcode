/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	cur := result
	for l1 != nil || l2 != nil {
		if l1 == nil {
			result.Next = l2
			l2 = nil
			continue
		}
		if l2 == nil {
			result.Next = l1
			l1 = nil
			continue
		}
		if l1.Val > l2.Val {
			result.Next = l2
			result, l2 = result.Next, l2.Next
		} else {
			result.Next = l1
			result, l1 = result.Next, l1.Next
		}
	}
	return cur.Next
}
