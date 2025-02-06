package cracking_the_coding_interview

import "testing"

func TestTreeNode_Add(t *testing.T) {
	n := NewTreeNode(1)
	n.Print()

	n.Add(2)
	n.Print()

	n.Add(0)
	n.Print()
}
