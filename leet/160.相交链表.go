/*
 * @lc app=leetcode.cn id=160 lang=golang
 *
 * [160] 相交链表
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	la, lb := findLen(headA), findLen(headB)
	if la > lb {
		for ; la > lb; la-- {
			headA = headA.Next
		}
	} else {
		for ; la < lb; lb-- {
			headB = headB.Next
		}
	}
	for headA != headB {
		headA = headA.Next
		headB = headB.Next
	}
	return headA
}

func findLen(res *ListNode) int {
	length := 0
	for ; res != nil; res = res.Next {
		length++
	}
	return length
}
