package cracking_the_coding_interview

import (
	"fmt"
	"testing"
)

// 使用栈
// 时间复杂度：O（n）
// 空间复杂度：O（n）
func isPalindrome(head *ListNode) bool {
	var stack []int
	cur := head
	for cur != nil {
		stack = append(stack, cur.Val)
		cur = cur.Next
	}

	for i := 0; i < len(stack)/2; i++ {
		if stack[i] != stack[len(stack)-i-1] {
			return false
		}
	}

	return true
}

// 使用快慢指针
// 时间复杂度：O（n）
// 空间复杂度：O（1）
func isPalindrome2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 反转链表的后半部分
	var prev *ListNode
	for slow != nil {
		next := slow.Next
		slow.Next = prev
		prev = slow
		slow = next
	}

	// 比较前半部分和反转后的后半部分
	left, right := head, prev
	for right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}

	return true
}

// 使用递归
// 时间复杂度：O（n）
// 空间复杂度：O（n）
func isPalindrome3(head *ListNode) bool {
	var front *ListNode

	var recursive func(*ListNode) bool
	recursive = func(cur *ListNode) bool {
		fmt.Printf("cur = %v\n", cur)
		if cur == nil {
			return true
		}

		if !recursive(cur.Next) {
			return false
		}

		if front.Val != cur.Val {
			return false
		}

		front = front.Next
		return true
	}

	front = head
	return recursive(front)
}

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		want bool
	}{
		{
			name: "empty node",
			head: nil,
			want: true,
		},
		{
			name: "not palindrome",
			head: NewListNode(1).Add(2).Add(3),
			want: false,
		},
		{
			name: "not palindrome",
			head: NewListNode(1).Add(1).Add(2).Add(1),
			want: false,
		},
		{
			name: "palindrome",
			head: NewListNode(1).Add(2).Add(2).Add(1),
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome3(tt.head); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
