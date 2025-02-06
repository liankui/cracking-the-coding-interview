package cracking_the_coding_interview

import "testing"

/*
•	dp[i] 表示在考虑到第 i 个请求时，按摩师能够获得的最大分钟数。
•	对于第 i 个请求，按摩师有两种选择：
1.	不接受第 i 个请求：此时她的最大分钟数是 dp[i-1]。
2.	接受第 i 个请求：此时，她必须跳过第 i-1 个请求，所以下一个能接受的请求是第 i-2 个请求，所以她的最大分钟数是 dp[i-2] + requests[i]。

// 时间复杂度：O（n）
// 空间复杂度：O（1）
*/
func massage(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	prev2 := 0
	prev1 := nums[0]

	for i := 1; i < len(nums); i++ {
		cur := max(prev1, prev2+nums[i])
		prev2 = prev1
		prev1 = cur
	}

	return prev1
}

func TestMessage(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "case1",
			nums: []int{1, 2, 3, 1},
			want: 4,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := massage(test.nums); got != test.want {
				t.Errorf("got: %v, want: %v", got, test.want)
			}
		})
	}
}
