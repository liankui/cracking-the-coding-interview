package cracking_the_coding_interview

import (
	"testing"
)

// 时间复杂度：O（n）
// 空间复杂度：O（1）
func trailingZeroes(n int) int {
	count := 0
	for i := 1; i <= n; i++ {
		if i%5 == 0 {
			count++
			for j := i / 5; j > 0 && j%5 == 0; j /= 5 {
				count++
			}
		}
	}

	return count
}

// 时间复杂度：O（log n）
// 空间复杂度：O（1）
func trailingZeroes2(n int) int {
	res := 0
	for n > 0 {
		n /= 5
		res += n
	}
	return res
}

func Test_trailingZeroes(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "3",
			n:    3,
			want: 0,
		},
		{
			name: "5",
			n:    5,
			want: 1,
		},
		{
			name: "10",
			n:    10,
			want: 2,
		},
		{
			name: "30",
			n:    30,
			want: 7,
		},
		{
			name: "200",
			n:    200,
			want: 49,
		},
		{
			name: "1808548329",
			n:    1808548329,
			want: 452137076,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := trailingZeroes2(test.n); got != test.want {
				t.Errorf("trailingZeroes() = %v, want %v", got, test.want)
			}
		})
	}
}
