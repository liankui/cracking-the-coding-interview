package cracking_the_coding_interview

import (
	"fmt"
	"math"
	"testing"
)

func reverseBits(num int) int {
	// if num == 0 {
	// 	return 1
	// }
	// if num < 0 {
	// 	num = ^num
	// }

	maxLen, prevLen, curLen := 0, 0, 0
	for num != 0 {
		if num&1 == 1 {
			curLen++
		} else {
			prevLen = curLen + 1
			curLen = 0
		}

		num >>= 1
		maxLen = max(maxLen, prevLen+curLen)
	}

	return maxLen
}

// 时间复杂度：O（n）
// 空间复杂度：O（1）
func reverseBits2(num int) int {
	maxLen, prevLen, curLen := 0, 0, 0
	for i := 0; i < 32; i++ {
		if num&1 == 1 {
			curLen++
		} else {
			prevLen = curLen + 1
			curLen = 0
		}

		maxLen = max(maxLen, prevLen+curLen)
		num >>= 1
	}

	return maxLen
}

func Test_reverseBits(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want int
	}{
		{
			name: "0",
			num:  0,
			want: 1,
		},
		{
			name: "7",
			num:  7,
			want: 4,
		},
		{
			name: "1775",
			num:  1775,
			want: 8,
		},
		{
			name: "-1",
			num:  -1,
			want: 32,
		},
		{
			name: "-2",
			num:  -2,
			want: 32,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := reverseBits2(test.num); got != test.want {
				t.Errorf("reverseBits(%d) = %d, want %d", test.num, got, test.want)
			}
		})
	}
}

func Test_negative(t *testing.T) {
	i := -1
	fi := ^i - 1
	fmt.Printf("%32b, ^i: %32b\n", i, fi)

	fi2 := math.MaxInt32
	fmt.Printf("%32b\n", fi2)

	fi3 := math.MinInt32 + i
	fmt.Printf("%32b\n", fi3)
}
