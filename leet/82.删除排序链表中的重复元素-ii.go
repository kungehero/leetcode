/*
 * @lc app=leetcode.cn id=82 lang=golang
 *
 * [82] 删除排序链表中的重复元素 II
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	res := &ListNode{0, head}
	for cur := res; cur != nil && cur.Next != nil; {
		val, duplicate, follow := cur.Next.Val, false, cur.Next.Next
		for follow != nil && follow.Val == val {
			duplicate = true
			follow = follow.Next
		}
		if duplicate {
			cur.Next = follow
		} else {
			cur = cur.Next
		}
	}
	return res.Next
}
