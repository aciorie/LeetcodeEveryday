package codetop2nd

type MinStack struct {
	minstack, everystack []int
}

func Constructor_155() MinStack {
	return MinStack{minstack: make([]int, 0), everystack: make([]int, 0)}
}

func (this *MinStack) Push(val int) {
	this.everystack = append(this.everystack, val)
	if len(this.minstack) == 0 || val <= this.minstack[len(this.minstack)-1] {
		this.minstack = append(this.minstack, val)
	}
}

func (this *MinStack) Pop() {
	top := this.everystack[len(this.everystack)-1]
	if top == this.minstack[len(this.minstack)-1] {
		this.minstack = this.minstack[:len(this.minstack)-1]
	}
	this.everystack = this.everystack[:len(this.everystack)-1]
}

func (this *MinStack) Top() int {
	return this.everystack[len(this.everystack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minstack[len(this.minstack)-1]
}
