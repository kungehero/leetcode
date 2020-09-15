/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 */

// @lc code=start
func exist(board [][]byte, word string) bool {
	path := []byte{}
	return backTracking(board, word, 0, 0, path) == true || false
}

func backTracking(board [][]byte, word string, l, r int, path []byte) bool {
	if len(path) == len(word) && string(path) == word {
		return true
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			path = append(path, board[i][j])
			backTracking(board, word, i, j, path)
			path = path[:len(path)-1]

		}
	}
	return false
}

// @lc code=end

