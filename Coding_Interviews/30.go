package codinginterviews

import "math"

/*
面试题 30. 包含 min 函数的栈


题目描述
定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，调用 min、push 及 pop 的时间复杂度都是 O(1)。



示例:

MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.min();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.min();   --> 返回 -2.


提示：

各函数的调用总次数不超过 20000 次
*/
// stack1 is Used to store data, stack2 is used to store the minimum value in the current stack
type MinStack struct {
	stack1 []int
	stack2 []int
}

func Constructor() MinStack {
	return MinStack{[]int{}, []int{math.MaxInt32}}
}

func (this *MinStack) Push(val int) {
	this.stack1 = append(this.stack1, val)
	this.stack2 = append(this.stack2, min(val, this.stack2[len(this.stack2)-1]))
}

func (this *MinStack) Pop() {
	this.stack1 = this.stack1[:len(this.stack1)-1]
	this.stack2 = this.stack2[:len(this.stack2)-1]
}

func (this *MinStack) Top() int {
	return this.stack1[len(this.stack1)-1]
}

func (this *MinStack) GetMin() int {
	return this.stack2[len(this.stack2)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
