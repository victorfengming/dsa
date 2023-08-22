package test

import "testing"

func TestName(t *testing.T) {

}

type MyQueue struct {
	inStack  []int
	outStack []int
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

func Constructor() MyQueue {
	// 构造方法, 返回myqueue对象
	return MyQueue{
		inStack:  make([]int, 0),
		outStack: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	// 往队列中存放元素
	this.inStack = append(this.inStack, x)
}

func (this *MyQueue) Pop() int {
	// 判断长度, 如果outStack为空,, 就把inStack里面的放到outStack里面, 反向
	lenIn := len(this.inStack)
	lenOut := len(this.outStack)
	if lenOut == 0 {
		if lenIn == 0 {
			//队列为空, 返回直接返回,
			return 0
		}
		// 将 inStack 的内容直接放入 outStack 并逆序
		for i := lenIn - 1; i >= 0; i-- {
			this.outStack = append(this.outStack, this.inStack[i])
		}
		this.inStack = make([]int, 0)

	}
	res := this.outStack[len(this.outStack)-1]
	this.outStack = this.outStack[:len(this.outStack)-1]
	return res

}

func (this *MyQueue) Peek() int {
	// 判断长度, 如果outStack为空,, 就把inStack里面的放到outStack里面, 反向
	lenIn := len(this.inStack)
	lenOut := len(this.outStack)
	if lenOut == 0 {
		if lenIn == 0 {
			//队列为空, 返回直接返回,
			return 0
		}
		// 将 inStack 的内容直接放入 outStack 并逆序
		for i := lenIn - 1; i >= 0; i-- {
			this.outStack = append(this.outStack, this.inStack[i])
		}
		this.inStack = make([]int, 0)

	}
	res := this.outStack[len(this.outStack)-1]
	return res
}

func (this *MyQueue) Empty() bool {
	return len(this.inStack) == 0 && len(this.outStack) == 0
}
