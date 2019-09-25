package leet

import "strconv"

/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	flag := true
	str := strconv.Itoa(x)
	for i := 0; i <= len(str)/2; i++ {
		if str[i] == str[len(str)-i-1] {
			flag = true
		} else {
			flag = false
			return flag
		}
	}
	return flag
}
