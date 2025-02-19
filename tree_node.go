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

// CreateTreeNode 前序创建
func CreateTreeNode(vals []int) *TreeNode {
	if len(vals) == 0 {
		return nil
	}

	root := NewTreeNode(vals[0])
	queue := []*TreeNode{root}

	for i := 1; i < len(vals); i++ {
		node := queue[0]
		queue = queue[1:]

		if vals[i] != -1 {
			node.Left = &TreeNode{Val: vals[i]}
			queue = append(queue, node.Left)
		}

		i++

		if i < len(vals) && vals[i] != -1 {
			node.Right = &TreeNode{Val: vals[i]}
			queue = append(queue, node.Right)
		}
	}

	return root
}

// Add 从左到右依次从小到大排列
func (n *TreeNode) Add(val int) *TreeNode {
	if n == nil {
		return n
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

	return n
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
