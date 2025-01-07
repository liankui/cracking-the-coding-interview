package cracking_the_coding_interview

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(val int) *ListNode {
	return &ListNode{Val: val}
}

func (n *ListNode) Add(val int) *ListNode {
	head := n

	if n.Next == nil {
		n.Next = NewListNode(val)
	} else {
		next := n.Next
		for next.Next != nil {
			next = next.Next
		}
		next.Next = NewListNode(val)
	}

	return head
}

func (n *ListNode) Print() {
	fmt.Printf("print: %d->", n.Val)
	next := n.Next
	for next != nil {
		fmt.Printf("%d->", next.Val)
		next = next.Next
	}
	fmt.Printf("\n")
}

func (n *ListNode) Reverse() *ListNode {
	var prev *ListNode
	cur := n
	for cur != nil {
		prev, cur, cur.Next = cur, cur.Next, prev
	}

	return prev
}
