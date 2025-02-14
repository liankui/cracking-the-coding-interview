package cracking_the_coding_interview

import "testing"

const modNum = 1000000007

// 动态规划
// 时间复杂度：O（n）
// 空间复杂度：O（n）
func waysToStep(n int) int {
	if n == 0 {
		return 1
	}

	dp := make([]int, n+1)
	dp[0] = 1
	if n >= 1 {
		dp[1] = 1
	}
	if n >= 2 {
		dp[2] = 2
	}

	for i := 3; i <= n; i++ {
		dp[i] = (dp[i-1] + dp[i-2] + dp[i-3]) % modNum
	}

	return dp[n]
}

// 动态规划，使用滚动数组优化空间
// 时间复杂度：O（n）
// 空间复杂度：O（1）
func waysToStep2(n int) int {
	if n == 0 {
		return 1
	}

	dp0, dp1, dp2 := 1, 1, 2
	if n == 1 {
		return dp1
	}
	if n == 2 {
		return dp2
	}

	for i := 3; i <= n; i++ {
		dpNext := (dp0 + dp1 + dp2) % modNum
		dp0, dp1, dp2 = dp1, dp2, dpNext
	}

	return dp2
}

func Test_waysToStep(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "0",
			n:    0,
			want: 1,
		},
		{
			name: "1",
			n:    1,
			want: 1,
		},
		{
			name: "3",
			n:    3,
			want: 4,
		},
		{
			name: "5",
			n:    5,
			want: 13,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := waysToStep2(tt.n); got != tt.want {
				t.Errorf("waysToStep() = %v, want %v", got, tt.want)
			}
		})
	}
}
