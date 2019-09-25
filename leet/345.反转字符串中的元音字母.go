import "strings"

/*
 * @lc app=leetcode.cn id=345 lang=golang
 *
 * [345] 反转字符串中的元音字母
 */
func reverseVowels(s string) string {
	vowel := "aeiouAEIOU"
	arr := []byte(s)
	i, j := 0, len(s)-1
	for j > i {
		if strings.Contains(vowel, string(arr[i])) && !strings.Contains(vowel, string(arr[j])) {
			j--
		} else if !strings.Contains(vowel, string(arr[i])) && strings.Contains(vowel, string(arr[j])) {
			i++
		} else if strings.Contains(vowel, string(arr[i])) && strings.Contains(vowel, string(arr[j])) {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		} else {
			i++
			j--
		}

	}
	return string(arr)
}
