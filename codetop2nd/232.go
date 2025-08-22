package codetop2nd

type MyQueue struct {
	inStack, outStack []int
}

func Constructor_232() MyQueue {
	return MyQueue{inStack: make([]int, 0), outStack: make([]int, 0)}
}

func (this *MyQueue) Push(x int) {
	this.inStack = append(this.inStack, x)
}

func (this *MyQueue) Pop() int {
	if len(this.outStack) == 0 {
		// 反转栈内元素顺序
		this.transfer()
	}
	val := this.outStack[len(this.outStack)-1]
	this.outStack = this.outStack[:len(this.outStack)-1]
	return val
}

func (this *MyQueue) transfer() {
	for len(this.inStack) > 0 {
		val := this.inStack[len(this.inStack)-1]
		this.inStack = this.inStack[:len(this.inStack)-1]
		this.outStack = append(this.outStack, val)
	}
}

func (this *MyQueue) Peek() int {
	if len(this.outStack) == 0 {
		this.transfer()
	}
	return this.outStack[len(this.outStack)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.inStack) == 0 && len(this.outStack) == 0
}
