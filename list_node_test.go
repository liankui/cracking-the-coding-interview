package cracking_the_coding_interview

import "testing"

func TestListNode_Print(t *testing.T) {
	node := NewListNode(1)
	node.Print()

	node.Add(2)
	node.Print()

	node.Add(3)
	node.Print()
}

func TestListNode_Reverse(t *testing.T) {
	node := NewListNode(1).Add(2).Add(3)
	node.Print()

	node.Reverse().Print()
}
