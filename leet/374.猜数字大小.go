func guessNumber(n int) int {
	for i := 0; i < n; i++ {
		if guess(i) == 0 {
			return i
		}
	}
	return n
}