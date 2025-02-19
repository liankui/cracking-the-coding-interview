package cracking_the_coding_interview

import "testing"

func TestTreeNode_Add(t *testing.T) {
	n := NewTreeNode(1)
	n.Print()

	n.Add(2)
	n.Print()

	n.Add(0)
	n.Print()

	preorderTraversal(n)
}

func TestTreeNode_CreateTreeNode(t *testing.T) {
	vals := []int{5, 4, 8, 11, -1, 13, 4, 7, 2, -1, -1, 5, 1}
	n := CreateTreeNode(vals)
	n.Print()
}
