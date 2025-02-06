package cracking_the_coding_interview

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{
		Val: val,
	}
}

func (n *TreeNode) Add(val int) {
	if n == nil {
		return
	}

	if val < n.Val {
		if n.Left == nil {
			n.Left = NewTreeNode(val)
		} else {
			n.Left.Add(val)
		}
	} else {
		if n.Right == nil {
			n.Right = NewTreeNode(val)
		} else {
			n.Right.Add(val)
		}
	}
}

func (n *TreeNode) Print() {
	if n == nil {
		return
	}

	stringify(n, 0)
	fmt.Println("--------------------")
}

func stringify(n *TreeNode, level int) {
	if n == nil {
		return
	}

	format := ""
	for i := 0; i < level; i++ {
		format += "\t"
	}
	format += "~ "
	level++
	stringify(n.Left, level)
	fmt.Printf(format+"%d\n", n.Val)
	stringify(n.Right, level)
}
