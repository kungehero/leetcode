/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号 ()[]{}
 */
func isValid(s string) bool {
	m := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	if len(s) == 0 {
		return true
	}
	stack := make([]rune, 0, 1)
	for _, v := range s {
		switch v {
		case 40, 123, 91:
			stack = append(stack, v)
		case 41, 125, 93:
			if len(stack) == 0 || stack[len(stack)-1] != m[v] {
				return false
			} else {
				stack = stack[:len(stack)-1]
				
			}
		}
	}
	return len(stack) == 0
}
