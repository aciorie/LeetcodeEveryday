package stack

type MinStack struct {
	minStack   []int
	everyStack []int
}

func Constructor() MinStack {
	return MinStack{
		minStack:   make([]int, 0),
		everyStack: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.everyStack = append(this.everyStack, val)
	if len(this.minStack) == 0 {
		this.minStack = append(this.minStack, val)
	}

	if val <= this.minStack[len(this.minStack)-1] {
		this.minStack = append(this.minStack, val)
	}
}

func (this *MinStack) Pop() {
	topElement := this.everyStack[len(this.everyStack)-1]
	if topElement == this.minStack[len(this.minStack)-1] {
		this.minStack = this.minStack[:len(this.minStack)-1]
	}
	this.everyStack = this.everyStack[:len(this.everyStack)-1]
}

func (this *MinStack) Top() int {
	return this.everyStack[len(this.everyStack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}
