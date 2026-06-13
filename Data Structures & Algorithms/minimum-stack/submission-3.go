type MinStack struct {
	arr []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{arr: []int{}, minStack: []int{math.MaxInt}}
}

func (this *MinStack) Push(val int) {
	this.arr = append(this.arr, val)
	if val <= this.minStack[len(this.minStack)-1] {
		this.minStack = append(this.minStack, val)
	}
}

func (this *MinStack) Pop() {
	if len(this.arr) == 0 {
		return 
	}
	if this.minStack[len(this.minStack) - 1] == this.arr[len(this.arr)-1] {
		this.minStack = this.minStack[:len(this.minStack)-1]
	}
	this.arr = this.arr[:len(this.arr)-1]
}

func (this *MinStack) Top() int {
	return this.arr[len(this.arr)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}
