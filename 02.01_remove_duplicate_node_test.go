package cracking_the_coding_interview

import (
	"reflect"
	"testing"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 哈希表
// 时间复杂度：O（n）
// 空间复杂度：O（n）
func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	m := make(map[int]struct{})
	dummy := &ListNode{Next: head}
	cur := dummy
	for cur.Next != nil {
		if _, ok := m[cur.Next.Val]; !ok {
			m[cur.Next.Val] = struct{}{}
			cur = cur.Next
		} else {
			cur.Next = cur.Next.Next
		}
	}

	return dummy.Next
}

// 不使用额外空间
// 时间复杂度：O（n^2）
// 空间复杂度：O（1）
func removeDuplicateNodes2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	cur := head

	for cur != nil {
		runner := cur

		for runner.Next != nil {
			if runner.Next.Val == cur.Val {
				runner.Next = runner.Next.Next
			} else {
				runner = runner.Next
			}
		}

		cur = cur.Next
	}

	return head
}

func Test_removeDuplicateNodes(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		want *ListNode
	}{
		{
			name: "empty list",
			head: nil,
			want: nil,
		},
		{
			name: "all node has same value",
			head: NewListNode(1).Add(1),
			want: NewListNode(1),
		},
		{
			name: "two nodes",
			head: NewListNode(1).Add(1).Add(2),
			want: NewListNode(1).Add(2),
		},
		{
			name: "three nodes",
			head: NewListNode(1).Add(2).Add(3).Add(3).Add(2).Add(1),
			want: NewListNode(1).Add(2).Add(3),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicateNodes(tt.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeDuplicateNodes() = %v, want %v", got, tt.want)
				got.Print()
				tt.want.Print()
			}
		})
	}
}
