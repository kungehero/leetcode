/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第N个节点
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	res := head
	mark := head
	tmp := head
	position, alreadyScan := 0, 0
	for tmp.Next != nil {
		if alreadyScan >= n {
			mark = mark.Next
			position++
		}
		alreadyScan++
		tmp = tmp.Next
	}
	if position == 0 && (n-alreadyScan != position) {
		res = res.Next
	} else {
		mark.Next = mark.Next.Next
	}
	return res
}
