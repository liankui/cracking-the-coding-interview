package cracking_the_coding_interview

import "testing"

// 时间复杂度：O（n）
// 空间复杂度：O（n）
func isUnique(astr string) bool {
	m := make(map[rune]struct{})
	for _, v := range astr {
		if _, ok := m[v]; ok {
			return false
		} else {
			m[v] = struct{}{}
		}
	}
	return true
}

// 时间复杂度：O（n^2）
// 空间复杂度：O（1）
func isUnique2(astr string) bool {
	for i := 0; i < len(astr)-1; i++ {
		for j := i + 1; j < len(astr); j++ {
			if astr[i] == astr[j] {
				return false
			}
		}
	}
	return true
}

// 时间复杂度：O（n）
// 空间复杂度：O（1）
func isUnique3(astr string) bool {
	bitMask := 0

	for i := 0; i < len(astr); i++ {
		pos := astr[i] - 'a'

		if bitMask&(1<<pos) != 0 {
			return false
		}

		bitMask |= 1 << pos
	}

	return true
}

func Test_isUnique(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "empty string",
			input: "",
			want:  true,
		},
		{
			name:  "all unique characters",
			input: "abc",
			want:  true,
		},
		{
			name:  "non unique characters",
			input: "abca",
			want:  false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := isUnique3(tc.input); got != tc.want {
				t.Errorf("isUnique(%v) = %v; want %v", tc.input, got, tc.want)
			}
		})
	}
}
