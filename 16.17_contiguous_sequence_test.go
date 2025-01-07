package cracking_the_coding_interview

import (
	"math"
	"testing"
)

// maxSubArray
// 时间复杂度：O（n^2）
// 空间复杂度：O（1）
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max := math.MinInt64
	for i := 0; i < len(nums); i++ {
		tmp := 0
		for j := i; j < len(nums); j++ {
			tmp += nums[j]
			if tmp > max {
				max = tmp
			}
		}
	}

	return max
}

func maxSubArray2(nums []int) int {
	return 0
}

func Test_maxSubArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "default",
			nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, // 4,-1,2,1
			want: 6,
		},
		{
			name: "empty",
			nums: []int{},
			want: 0,
		},
		{
			name: "all negative",
			nums: []int{-2, -1, -3, -4, -1, -2, -1, -5, -4},
			want: -1,
		},
		{
			name: "last positive",
			nums: []int{-2, -1, -3, -4, -1, -2, -1, -5, 4},
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray(tt.nums); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
