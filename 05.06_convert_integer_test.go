package cracking_the_coding_interview

import (
	"math/bits"
	"testing"
)

// 使用异或操作
// 时间复杂度：O（n）
// 空间复杂度：O（1）
func convertInteger(A int, B int) int {
	diff := int32(A ^ B)
	count := 0
	for diff != 0 {
		count++
		diff &= diff - 1
	}
	return count
}

// 逐位比较
// 时间复杂度：O（n）
// 空间复杂度：O（1）
// 效率问题
func convertInteger2(A int, B int) int {
	count := 0
	for A != 0 || B != 0 {
		if A&1 != B&1 {
			count++
		}
		A >>= 1
		B >>= 1
	}
	return count
}

// 位运算 + 内置函数
// 时间复杂度：O（n）
// 空间复杂度：O（1）
func convertInteger3(A int, B int) int {
	return bits.OnesCount32(uint32(A ^ B))
}

func Test_convertInteger(t *testing.T) {
	tests := []struct {
		name string
		A    int
		B    int
		want int
	}{
		{
			name: "A equals B",
			A:    1,
			B:    1,
			want: 0,
		},
		{
			name: "A not equals B",
			A:    1,
			B:    2,
			want: 2,
		},
		{
			name: "A not equals B 2",
			A:    29,
			B:    15,
			want: 2,
		},
		{
			name: "A not equals B 2",
			A:    826966453,
			B:    -729934991,
			want: 46,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertInteger3(tt.A, tt.B); got != tt.want {
				t.Errorf("convertInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}
