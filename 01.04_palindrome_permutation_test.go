package cracking_the_coding_interview

import "testing"

// 时间复杂度：O（n）
// 空间复杂度：O（n）
func canPermutePalindrome(s string) bool {
	l := len(s)
	if l == 0 {
		return true
	}

	m := make(map[rune]int, len(s))
	for _, char := range s {
		m[char]++
	}

	oddNum := 0
	for _, v := range m {
		if v%2 != 0 {
			oddNum++
		}
	}

	if l%2 == 0 {
		if oddNum > 0 {
			return false
		}
	} else {
		if oddNum > 1 {
			return false
		}
	}

	return true
}

// 时间复杂度：O（n）
// 空间复杂度：O（n）
func canPermutePalindrome2(s string) bool {
	var mask int64
	for _, char := range s {
		mask ^= 1 << char
	}

	return (mask == 0) || ((mask & (mask - 1)) == 0)
}

func Test_canPermutePalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "empty string",
			s:    "",
			want: true,
		},
		{
			name: "not palindrome",
			s:    "abcd",
			want: false,
		},
		{
			name: "not palindrome2",
			s:    "code",
			want: false,
		},
		{
			name: "palindrome",
			s:    "tactcoa",
			want: true,
		},
		{
			name: "palindrome2",
			s:    "tactca",
			want: true,
		},
		{
			name: "palindrome3",
			s:    "AaBb//a",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPermutePalindrome2(tt.s); got != tt.want {
				t.Errorf("canPermutePalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
