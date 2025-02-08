package cracking_the_coding_interview

import (
	"math"
	"testing"
)

// func maximum(a int, b int) int {
// 	signBit := (a - b) >> (uint64(^0) >> 1)
//
// 	mask := signBit & 1
//
// 	return a - mask*(a-b)
// }

func maximum2(a int, b int) int {
	return max(a, b)
}

func maximum3(a int, b int) int {
	return (a + b + int(math.Abs(float64(a-b)))) / 2
}

func maximum4(a int, b int) int {
	// (a - b) >> 31 将差值的符号位移到最低位（如果 a > b，结果为 0；如果 a < b，结果为 -1）。
	// & 运算符会确保最终只取较大的那个数
	return a - ((a - b) & ((a - b) >> 63))
}

func Test_maximum(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{
			name: "max 2",
			a:    1,
			b:    2,
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum4(tt.a, tt.b); got != tt.want {
				t.Errorf("maximum() = %v, want %v", got, tt.want)
			}
		})
	}
}
