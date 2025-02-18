package cracking_the_coding_interview

import (
	"reflect"
	"testing"
)

// 使用两个单独的链表拆分
// 时间复杂度：O（n）
// 空间复杂度：O（n）
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}

	lessHead := &ListNode{}
	greaterHead := &ListNode{}
	less := lessHead
	greater := greaterHead

	for head != nil {
		if head.Val < x {
			less.Next = head
			less = less.Next
		} else {
			greater.Next = head
			greater = greater.Next
		}

		head = head.Next
	}

	// 防止环形链表，切断 greater 部分的连接
	greater.Next = nil

	less.Next = greaterHead.Next

	return lessHead.Next
}

// 原地修改法
// 时间复杂度：O（n）
// 空间复杂度：O（1）
func partition2(head *ListNode, x int) *ListNode {
	var lessHead, greaterHead, less, greater *ListNode

	for head != nil {
		if head.Val < x {
			if less == nil {
				lessHead = head
				less = lessHead
			} else {
				less.Next = head
				less = less.Next
			}
		} else {
			if greater == nil {
				greaterHead = head
				greater = greaterHead
			} else {
				greater.Next = head
				greater = greater.Next
			}
		}

		head = head.Next
	}

	// 连接两个部分
	if less != nil {
		less.Next = greaterHead
	}
	if greater != nil {
		greater.Next = nil
	}

	if lessHead != nil {
		return lessHead
	}

	return greaterHead
}

func Test_partition(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		x    int
		want *ListNode
	}{
		{
			name: "head",
			head: NewListNode(3).Add(5).Add(8).Add(5).Add(10).Add(2).Add(1),
			x:    5,
			want: NewListNode(3).Add(2).Add(1).Add(5).Add(8).Add(5).Add(10),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition2(tt.head, tt.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}
