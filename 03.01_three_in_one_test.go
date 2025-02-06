package cracking_the_coding_interview

import (
	"fmt"
	"testing"
)

type TripleInOne struct {
	array     []int
	stackSize int
	tops      []int
}

func Constructor(stackSize int) TripleInOne {
	return TripleInOne{
		array:     make([]int, 3*stackSize),
		stackSize: stackSize,
		tops:      []int{-1, -1, -1},
	}
}

func (t *TripleInOne) Push(stackNum int, value int) {
	if t.tops[stackNum] < t.stackSize-1 {
		t.tops[stackNum]++
		t.array[stackNum*t.stackSize+t.tops[stackNum]] = value
	}
}

func (t *TripleInOne) Pop(stackNum int) int {
	if t.tops[stackNum] == -1 {
		return -1
	}

	val := t.array[stackNum*t.stackSize+t.tops[stackNum]]
	t.tops[stackNum]--
	return val
}

func (t *TripleInOne) Peek(stackNum int) int {
	if t.tops[stackNum] == -1 {
		return -1
	}

	return t.array[stackNum*t.stackSize+t.tops[stackNum]]
}

func (t *TripleInOne) IsEmpty(stackNum int) bool {
	return t.tops[stackNum] == -1
}

/**
 * Your TripleInOne object will be instantiated and called as such:
 * obj := Constructor(stackSize);
 * obj.Push(stackNum,value);
 * param_2 := obj.Pop(stackNum);
 * param_3 := obj.Peek(stackNum);
 * param_4 := obj.IsEmpty(stackNum);
 */

func TestTripleInOne(t *testing.T) {
	obj := Constructor(3)
	obj.Push(0, 2)
	obj.Push(0, 1)
	fmt.Printf("%v\n", obj)
	pop1 := obj.Pop(0)
	fmt.Printf("Pop(%d) = %d\n", 0, pop1)
	pop2 := obj.Pop(0)
	fmt.Printf("Pop(%d) = %d\n", 0, pop2)
	pop3 := obj.Pop(0)
	fmt.Printf("Pop(%d) = %d\n", 0, pop3)
	isEmpty := obj.IsEmpty(0)
	fmt.Printf("IsEmpty(%d) = %t\n", 0, isEmpty)
}
