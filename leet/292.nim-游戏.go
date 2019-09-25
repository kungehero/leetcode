/*
 * @lc app=leetcode.cn id=292 lang=golang
 *
 * [292] Nim 游戏
 */
func canWinNim(n int) bool {
	if n%4 == 0 {
		return false
	}
	return true
}
