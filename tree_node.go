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

/*
二叉树的前序、中序和后序遍历分别是按以下顺序读取的：
	1.	前序遍历（Pre-order Traversal）：
	•	顺序：根节点 -> 左子树 -> 右子树
	•	先访问根节点，然后递归地访问左子树，最后递归地访问右子树。
	•	例如：根 -> 左 -> 右
	2.	中序遍历（In-order Traversal）：
	•	顺序：左子树 -> 根节点 -> 右子树
	•	先递归访问左子树，然后访问根节点，最后递归访问右子树。
	•	例如：左 -> 根 -> 右
	•	对于二叉搜索树，中序遍历的结果是按升序排列的。
	3.	后序遍历（Post-order Traversal）：
	•	顺序：左子树 -> 右子树 -> 根节点
	•	先递归访问左子树，然后递归访问右子树，最后访问根节点。
	•	例如：左 -> 右 -> 根
*/

func preorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.Val)
	preorderTraversal(root.Left)
	preorderTraversal(root.Right)
}
