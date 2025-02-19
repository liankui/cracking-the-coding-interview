package cracking_the_coding_interview

import "testing"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 递归遍历二叉树并在每个节点上计算路径和。
// 深度优先搜索（DFS）
// 时间复杂度：O（n^2）
// 空间复杂度：O（log n）
func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	return countPaths(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

// 辅助函数：计算从当前节点开始的路径和为 sum 的路径数
func countPaths(node *TreeNode, sum int) int {
	if node == nil {
		return 0
	}

	count := 0
	if node.Val == sum {
		count++
	}

	count += countPaths(node.Left, sum-node.Val)
	count += countPaths(node.Right, sum-node.Val)

	return count
}

// 哈希+前缀和
// 深度优先搜索（DFS）
// 时间复杂度：O（n）
// 空间复杂度：O（n）
func pathSum2(root *TreeNode, sum int) int {
	return dfs2(root, sum, make(map[int]int), 0)
}

func dfs2(node *TreeNode, sum int, prefixSum map[int]int, curSum int) int {
	if node == nil {
		return 0
	}

	count := 0
	// 更新当前路径的和
	curSum += node.Val

	if curSum == sum {
		count++
	}

	// 如果当前路径的减去目标值存在，则找到一个符合的路径
	if val, ok := prefixSum[curSum-sum]; ok {
		count += val
	}

	// 记录当前路径和出现的次数
	prefixSum[curSum]++

	count += dfs2(node.Left, sum, prefixSum, curSum)
	count += dfs2(node.Right, sum, prefixSum, curSum)

	// 回溯，撤销当前节点的路径和
	prefixSum[curSum]--

	return count
}

func Test_pathSum(t *testing.T) {
	tests := []struct {
		name string
		tree *TreeNode
		sum  int
		want int
	}{
		{
			/*
			         5
			        / \
			       4   8
			      /   / \
			     11  13  4
			    /  \    / \
			   7    2  5   1
			*/
			name: "case1",
			tree: CreateTreeNode([]int{5, 4, 8, 11, -1, 13, 4, 7, 2, -1, -1, 5, 1}),
			sum:  22,
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pathSum2(tt.tree, tt.sum); got != tt.want {
				t.Errorf("pathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
