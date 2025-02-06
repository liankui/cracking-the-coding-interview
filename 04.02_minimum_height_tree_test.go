package cracking_the_coding_interview

import (
	"fmt"
	"testing"
)

func sortedArrayToBST(nums []int) *TreeNode {
	return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}

	mid := left + (right-left)/2
	root := &TreeNode{Val: nums[mid]}

	root.Left = helper(nums, left, mid-1)
	root.Right = helper(nums, mid+1, right)

	return root
}

func Test_sortedArrayToBST(t *testing.T) {
	nums := []int{-10, -3, 0, 5, 9}
	root := sortedArrayToBST(nums)

	// Print the tree in pre-order to verify
	fmt.Println("Pre-order traversal of the constructed BST:")
	preorderTraversal(root)
}
