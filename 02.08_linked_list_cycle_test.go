package cracking_the_coding_interview

// 遍历指针
// 时间复杂度：O（n）
// 空间复杂度：O（n）
func detectCycle(head *ListNode) *ListNode {
	seen := make(map[*ListNode]struct{})
	for head != nil {
		if _, ok := seen[head]; ok {
			return head
		} else {
			seen[head] = struct{}{}
			head = head.Next
		}
	}
	return nil
}

// 快慢指针
// 时间复杂度：O（n）
// 空间复杂度：O（1）
func detectCycle2(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}

		fast = fast.Next.Next
		if fast == slow {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}

			return p
		}
	}

	return nil
}

// func Test_detectCycle(t *testing.T) {
// 	node := NewListNode(1)
// 	node1 := node.Add(2).Add(3)
// 	node1.Next = node
//
// 	tests := []struct {
// 		name string
// 		head *ListNode
// 		want *ListNode
// 	}{
// 		{
// 			name: "case1",
// 			head: node,
// 			want:
// 		},
// 	}
// }
