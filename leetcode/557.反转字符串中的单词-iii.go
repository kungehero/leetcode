/*
 * @lc app=leetcode.cn id=557 lang=golang
 *
 * [557] 反转字符串中的单词 III
 */

// @lc code=start

func reverseWords(s string) string {
	strings.TrimSpace(s)
	str := ""
	i := 0
	path := []byte{}

	for i < len(s) {
		//if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') || (s[i] =='')
		if s[i] != ' ' {
			path = append(path, s[i])
		} else {
			str += reverse(path)
			path = path[0:0]
			str += string(s[i])
		}

		if i == len(s)-1 {
			str += reverse(path)
		}
		i++
	}
	return str
}

func reverse(b []byte) string {
	i, j := 0, len(b)-1
	for i < j {
		b[i], b[j] = b[j], b[i]
		i++
		j--
	}
	return string(b)
}

// @lc code=end

