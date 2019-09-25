import "strconv"

/*
 * @lc app=leetcode.cn id=412 lang=golang
 *
 * [412] Fizz Buzz
 */
func fizzBuzz(n int) []string {
	var res []string
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 != 0 {
			res = append(res, "Fizz")
		} else if i%5 == 0 && i%3 != 0 {
			res = append(res, "Buzz")
		} else if i%5 == 0 && i%3 == 0 {
			res = append(res, "FizzBuzz")
		} else {
			res = append(res, strconv.Itoa(i))
		}
	}
	return res
}
