/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		head.Next = prev
		prev = head
		head = head.Next
	}
	return prev
}
