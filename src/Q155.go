package main

import "fmt"

func main() {
	stack := Constructor1()
	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)
	fmt.Println(stack.GetMin())
	stack.Pop()
	fmt.Println(stack.Top())
	fmt.Println(stack.GetMin())
}



type MinStack struct {
	S []int
}

/** initialize your data structure here. */
func Constructor1() MinStack {
	return MinStack{
		S: []int{},
	}
}

func (this *MinStack) Push(x int) {
	this.S = append(this.S, x)
}

func (this *MinStack) Pop() {
	if len(this.S) <= 1 {
		this.S = []int{}
	}else{
		this.S = this.S[:len(this.S)-1]
	}

}

func (this *MinStack) Top() int {
	return this.S[0]
}

func (this *MinStack) GetMin() int {
	if len(this.S) <= 1 {
		return this.S[0]
	}else{
		min := this.S[0]
		for _,v :=range this.S{
			if v < min {
				min = v
			}
		}
		return min
	}
}
