package cracking_the_coding_interview

import (
	"strings"
	"testing"
)

// 时间复杂度：O（n）
// 空间复杂度：O（n）
func isFlippedString(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	if len(s1) == 0 && len(s2) == 0 {
		return true
	}

	for i := 0; i < len(s1); i++ {
		rotation := s1[i:] + s1[:i]
		if rotation == s2 {
			return true
		}
	}

	return false
}

// 时间复杂度：O（n）
// 空间复杂度：O（n）
func isFlippedString2(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	return strings.Contains(s2+s2, s1)
}

func Test_isFlippedString(t *testing.T) {
	tests := []struct {
		name string
		s1   string
		s2   string
		want bool
	}{
		{
			name: "all empty",
			s1:   "",
			s2:   "",
			want: true,
		},
		{
			name: "abc",
			s1:   "abc",
			s2:   "abc",
			want: true,
		},
		{
			name: "rotated",
			s1:   "abc",
			s2:   "cab",
			want: true,
		},
		{
			name: "length not equal",
			s1:   "abc",
			s2:   "ab",
			want: false,
		},
		{
			name: "waterbottle",
			s1:   "waterbottle",
			s2:   "erbottlewat",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isFlippedString2(tt.s1, tt.s2); got != tt.want {
				t.Errorf("isFlipedString() = %v, want %v", got, tt.want)
			}
		})
	}
}
