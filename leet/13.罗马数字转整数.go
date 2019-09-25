/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] 罗马数字转整数
 */
package leet

func romanToInt(s string) int {
	roman := map[string]int{
		"I":  1,
		"IV": 4,
		"V":  5,
		"IX": 9,
		"X":  10,
		"XL": 40,
		"L":  50,
		"XC": 90,
		"C":  100,
		"CD": 400,
		"D":  500,
		"CM": 900,
		"M":  1000,
	}
	if _, ok := roman[s]; ok {
		return roman[s]
	}

	result := 0
	i := 0
	for {
		if i >= len(s) {
			break
		}
		if i == len(s)-1 {
			result += roman[s[i:i+1]]
			break
		}
		switch s[i] {
		case 'I':
			fallthrough
		case 'X':
			fallthrough
		case 'C':
			if v, ok := roman[s[i:i+2]]; ok {
				result += v
				i++
			} else {
				result += roman[s[i:i+1]]
			}
		default:
			result += roman[s[i:i+1]]
		}
		i++
	}
	return result
}
