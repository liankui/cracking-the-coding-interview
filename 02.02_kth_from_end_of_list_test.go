package cracking_the_coding_interview

import "testing"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// kthToLast
// 时间复杂度：O（n）
// 空间复杂度：O（n）
func kthToLast(head *ListNode, k int) int {
	vals := make([]int, 0)
	for head != nil {
		vals = append(vals, head.Val)
		head = head.Next
	}

	return vals[len(vals)-k]
}

// 时间复杂度：O（n）
// 空间复杂度：O（1）
func kthToLast2(head *ListNode, k int) int {
	var prev *ListNode
	cur := head
	for cur != nil {
		prev, cur, cur.Next = cur, cur.Next, prev
	}

	for i := 1; i <= k; i++ {
		if i == k {
			return prev.Val
		}
		prev = prev.Next
	}

	return 0
}

// 时间复杂度：O（n）
// 空间复杂度：O（1）
func kthToLast3(head *ListNode, k int) int {
	fast, slow := head, head
	for k > 0 {
		fast = fast.Next
		k--
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	return slow.Val
}

func Test_kthToLast(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		k    int
		want int
	}{
		{
			name: "default",
			head: NewListNode(1).Add(2).Add(3).Add(4).Add(5),
			k:    2,
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthToLast3(tt.head, tt.k); got != tt.want {
				t.Errorf("kthToLast() = %v, want %v", got, tt.want)
			}
		})
	}
}
