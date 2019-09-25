/*
 * @lc app=leetcode.cn id=168 lang=golang
 *
 * [168] Excel表列名称
 */
func convertToTitle(n int) string {
	Slice := []byte{'c', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O',
		'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	var temp []byte
	if n < 27 {
		temp = append(temp, Slice[n])
		return string(temp)
	}
	if n >= 27 {
		for {
			k := n % 26 //从个位开始拆分，如果求余为0，说明末尾为26，也就是Z，如果是转化为26进制数，则末尾是可以为0的，这里必须为A-Z中的一个
			if k == 0 {
				temp = append(temp, Slice[26])
				k = 26
			} else {
				temp = append(temp, Slice[k])
			}
			n = (n - k) / 26 //减去Num最后一位数的值，因为已经记录在temp中
			if n <= 26 {     //小于等于26直接进行匹配，不需要进行数据拆分
				temp = append(temp, Slice[n])
				break
			}
		}
	}
	Reverse(temp)
	return string(temp)
}

func Reverse(arr []byte) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
