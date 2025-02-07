package cracking_the_coding_interview

import "testing"

// 时间复杂度：O（n）
// 空间复杂度：O（n）
func CheckPermutation(s1 string, s2 string) bool {
	m1 := make(map[rune]int, len(s1))
	m2 := make(map[rune]int, len(s2))

	for _, v := range s1 {
		m1[v]++
	}

	for _, v := range s2 {
		m2[v]++
	}

	if len(m1) != len(m2) {
		return false
	} else {
		for k, v := range m2 {
			if m1[k] != v {
				return false
			}
		}
	}

	return true
}

// 时间复杂度：O（n log n）
// 空间复杂度：O（log n）
func CheckPermutation2(s1 string, s2 string) bool {
	return quickSortString(s1) == quickSortString(s2)
}

func quickSortString(s string) string {
	if len(s) == 0 {
		return ""
	}

	pivot := s[0]

	var left, right string
	for i := 1; i < len(s); i++ {
		if pivot < s[i] {
			left += string(s[i])
		} else if pivot >= s[i] {
			right += string(s[i])
		}
	}

	left = quickSortString(left)
	right = quickSortString(right)

	return left + string(pivot) + right
}

func Test_CheckPermutation(t *testing.T) {
	tests := []struct {
		name string
		s1   string
		s2   string
		want bool
	}{
		{
			name: "empty s1",
			s1:   "",
			s2:   "aa",
			want: false,
		},
		{
			name: "empty s2",
			s1:   "a",
			s2:   "",
			want: false,
		},
		{
			name: "is true",
			s1:   "abc",
			s2:   "bac",
			want: true,
		},
		{
			name: "is true, duplicated",
			s1:   "abcc",
			s2:   "cbac",
			want: true,
		},
		{
			name: "is false",
			s1:   "abc",
			s2:   "bcd",
			want: false,
		},
		{
			name: "is false, duplicated",
			s1:   "abcc",
			s2:   "bcdc",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckPermutation2(tt.s1, tt.s2); got != tt.want {
				t.Errorf("CheckPermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
