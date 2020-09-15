package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{Val: 2}
	l1.Next = &ListNode{Val: 4}
	l1.Next.Next = &ListNode{Val: 3}
	//l1.Next.Next.Next = &ListNode{Val: 7}

	l2 := &ListNode{Val: 5}
	l2.Next = &ListNode{Val: 6}
	l2.Next.Next = &ListNode{Val: 4}

	/* str := "i am a student"
	strarr := []byte(str) */

	/* 	l, r := 0, 0
	   	num := len(strarr)
	   	for num > 0 {
	   		if strarr[r] == ' ' || r == len(strarr)-1 {

	   			if r == len(strarr)-1 {
	   				arrReverse(strarr, l, r+1)
	   			} else {
	   				arrReverse(strarr, l, r)
	   			}

	   			l = r + 1
	   			r = l

	   		} else {
	   			r++
	   		}
	   		num--
	   	}
	   	arrReverse(strarr, 0, len(strarr))
		   fmt.Println(string(strarr)) */

	test := []string{""}
	for i := 0; i < 5; i++ {
		for _, v := range test {
			fmt.Println(i, v)
		}
	}
}

func arrReverse(nums []byte, l, r int) {
	r = r - 1
	for r > l {
		nums[r], nums[l] = nums[l], nums[r]
		l++
		r--
	}
}

func reverse(head *ListNode) (res *ListNode, tail *ListNode) {
	var pre *ListNode
	cur := head
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre, head
}

func isCheck(mps, mpdinx map[byte]int) bool {
	for k, v := range mpdinx {
		if mps[k] < v {
			return false
		}
	}
	return true
}

func mergeSort(nums []int) []int {

	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2
	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])
	return merge(left, right)
}

func merge(l1, l2 []int) []int {
	var res []int
	for len(l1) > 0 && len(l2) > 0 {
		if l1[0] < l2[0] {
			res = append(res, l1[0])
			l1 = l1[1:]
		} else {
			res = append(res, l2[0])
			l2 = l2[1:]
		}
	}
	if len(l1) == 0 {
		res = append(res, l2...)
	}
	if len(l2) == 0 {
		res = append(res, l1...)
	}
	return res
}

func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	explodindex := left

	for i := left + 1; i <= right; i++ {
		if nums[left] > nums[i] {
			explodindex++
			nums[i], nums[explodindex] = nums[explodindex], nums[i]
		}
	}
	nums[left], nums[explodindex] = nums[explodindex], nums[left]
	quickSort(nums, left, explodindex-1)
	quickSort(nums, explodindex+1, right)
}
