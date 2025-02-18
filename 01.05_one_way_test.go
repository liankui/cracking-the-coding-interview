package cracking_the_coding_interview

import (
	"testing"
)

// 计算长度差值添加判断
// 时间复杂度：O（n）
// 空间复杂度：O（1）
func oneEditAway(first, second string) bool {
	if len(first) > len(second) {
		first, second = second, first
	}

	diff := len(second) - len(first)

	if diff > 1 {
		return false
	} else if diff == 1 {
		if len(first) == 0 || len(second) == 0 {
			return true
		}

		for i := 0; i < len(first); i++ {
			if first[i] != second[i] {
				second = second[:i] + second[i+1:]
				break
			}
		}

		return first == second[:len(first)]
	} else {
		count := 0
		for i := 0; i < len(first); i++ {
			if first[i] != second[i] {
				count++
			}
		}

		return count <= 1
	}
}

func Test_oneEditAway(t *testing.T) {
	tests := []struct {
		name   string
		first  string
		second string
		want   bool
	}{
		{
			name:   "all empty",
			first:  "",
			second: "",
			want:   true,
		},
		{
			name:   "first empty",
			first:  "",
			second: "a",
			want:   true,
		},
		{
			name:   "first empty2",
			first:  "",
			second: "ab",
			want:   false,
		},
		{
			name:   "second empty",
			first:  "a",
			second: "",
			want:   true,
		},
		{
			name:   "second empty2",
			first:  "ab",
			second: "",
			want:   false,
		},
		{
			name:   "second removed a character",
			first:  "abc",
			second: "ab",
			want:   true,
		},
		{
			name:   "second removed a character2",
			first:  "abc",
			second: "ac",
			want:   true,
		},
		{
			name:   "second removed a character3",
			first:  "abc",
			second: "bc",
			want:   true,
		},
		{
			name:   "second edit a character",
			first:  "abc",
			second: "avc",
			want:   true,
		},
		{
			name:   "second insert a character",
			first:  "ac",
			second: "abc",
			want:   true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := oneEditAway(test.first, test.second); got != test.want {
				t.Errorf("oneEditAway(%q, %q) = %v, want %v", test.first, test.second, got, test.want)
			}
		})
	}
}
