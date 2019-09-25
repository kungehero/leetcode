/*
 * @lc app=leetcode.cn id=225 lang=golang
 *
 * [225] 用队列实现栈
 */
type MyStack struct {
	head *Node
}
type Node struct {
	Value int
	Next  *Node
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{nil}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	oldHead := this.head
	this.head = &Node{x, oldHead}
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if this.Empty() {
		return 0
	}
	cur := this.head
	this.head = this.head.Next
	return cur.Value
}

/** Get the top element. */
func (this *MyStack) Top() int {
	if this.Empty() {
		return 0
	}
	return this.head.Value
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.head == nil
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
