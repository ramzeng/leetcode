/*
 * @lc app=leetcode.cn id=232 lang=golang
 *
 * [232] 用栈实现队列
 */

// @lc code=start
type MyQueue struct {
	in, out []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) transfer() {
	for i := len(this.in) - 1; i >= 0; i-- {
		this.out = append(this.out, this.in[i])
	}

	this.in = nil
}

func (this *MyQueue) Push(x int) {
	this.in = append(this.in, x)
}

func (this *MyQueue) Pop() int {
	if this.Empty() {
		return -1
	}

	if len(this.out) == 0 {
		this.transfer()
	}

	value := this.out[len(this.out)-1]
	this.out = this.out[:len(this.out)-1]
	return value
}

func (this *MyQueue) Peek() int {
	if this.Empty() {
		return -1
	}

	if len(this.out) == 0 {
		this.transfer()
	}

	return this.out[len(this.out)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.in) == 0 && len(this.out) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
// @lc code=end

