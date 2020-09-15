/*
 * @lc app=leetcode.cn id=30 lang=golang
 *
 * [30] 串联所有单词的子串
 */

// @lc code=start
func findSubstring(s string, words []string) []int {

	res := []int{}
	if len(s) == 0 || len(words) == 0 {
		return res
	}

	wordssigle := len(words[0])
	wordslen := len(words) * wordssigle
	//	ms := make(map[string]int)
	wordmap := make(map[string]int)
	l, r := 0, wordslen
	for i := 0; i < len(words); i++ {
		wordmap[words[i]]++
	}

	for r <= len(s) {
		str := s[l:r]
		ms := make(map[string]int)
		for i := 0; i < len(words); i++ {
			ms[string(str[i*wordssigle:wordssigle*(i+1)])]++
		}
		if check(ms, wordmap) && r <= len(s) {
			res = append(res, l)
		}

		l++
		r = l + wordslen
	}
	return res
}

func check(ms map[string]int, wordmap map[string]int) bool {
	for k, v := range wordmap {
		if ms[k] < v {
			return false
		}
	}
	return true
}

// @lc code=end

