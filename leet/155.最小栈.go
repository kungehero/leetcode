/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 */
type intStack struct {
	s []int
}

func (s *intStack) empty() bool {
	return len(s.s) == 0
}

func (s *intStack) push(n int) {
	s.s = append(s.s, n)
}

func (s *intStack) peek() int {
	if s.empty() {
		panic("stack is empty")
	}

	return s.s[len(s.s)-1]
}

func (s *intStack) pop() int {
	if s.empty() {
		panic("stack is empty")
	}

	val := s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]

	return val
}

type MinStack struct {
	mins intStack
	nums intStack
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.nums.push(x)

	// we need to check if the mins stack is empty
	// the element has to be the smallest so far
	if this.mins.empty() || x <= this.mins.peek() {
		this.mins.push(x)
	}
}

func (this *MinStack) Pop() {
	val := this.nums.pop()

	if val == this.mins.peek() {
		this.mins.pop()
	}
}

func (this *MinStack) Top() int {
	return this.nums.peek()
}

func (this *MinStack) GetMin() int {
	return this.mins.peek()
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
