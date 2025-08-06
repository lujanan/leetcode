// 155.最小栈

package algorithm_100

type MinStack struct {
	Array    []int
	MinArray []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.Array = append(this.Array, val)
	if len(this.Array) > 1 {
		this.MinArray = append(this.MinArray, min(this.MinArray[len(this.MinArray)-1], val))
	} else {
		this.MinArray = append(this.MinArray, val)
	}
}

func (this *MinStack) Pop() {
	this.Array = this.Array[:len(this.Array)-1]
	this.MinArray = this.MinArray[:len(this.MinArray)-1]
}

func (this *MinStack) Top() int {
	return this.Array[len(this.Array)-1]
}

func (this *MinStack) GetMin() int {
	return this.MinArray[len(this.MinArray)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
