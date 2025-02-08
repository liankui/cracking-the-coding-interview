package cracking_the_coding_interview

import (
	"fmt"
	"testing"
)

// 时间复杂度：O（1）
// 空间复杂度：O（1）
func add(a int, b int) int {
	for b != 0 {
		// 计算无进位部分
		tmp := a ^ b // 异或，不同为1，相同为0
		// 计算进位部分
		b = (a & b) << 1
		a = tmp
		fmt.Printf("a = %4b, b = %4b\n", a, b)
	}
	return a
}

// 时间复杂度：O（1）
// 空间复杂度：O（1）
func add2(a int, b int) int {
	if b == 0 {
		return a
	}
	return add(a^b, (a&b)<<1)
}

func Test_add(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{
			name: "1+1",
			a:    1,
			b:    1,
			want: 2,
		}, {
			name: "1+2",
			a:    1,
			b:    2,
			want: 3,
		}, {
			name: "1-11",
			a:    1,
			b:    -11,
			want: -10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add2(tt.a, tt.b); got != tt.want {
				t.Errorf("add() = %d, want %d", got, tt.want)
			}
		})
	}
}
