package cracking_the_coding_interview

import (
	"strings"
	"testing"
)

// 时间复杂度：O（n）
// 空间复杂度：O（n）
// 使用 strings.Builder 可以有效避免频繁的字符串拼接操作，适合性能要求较高的场景
func replaceSpaces(S string, length int) string {
	t := strings.Builder{}
	for i := 0; i < length; i++ {
		if i >= len(S) || S[i] == ' ' {
			t.WriteString("%20")
		} else {
			t.WriteByte(S[i])
		}
	}

	return t.String()
}

// 时间复杂度：O（n）
// 空间复杂度：O（n）
// 直接通过字符串的拼接来实现替换（性能问题）
func replaceSpaces2(S string, length int) string {
	result := ""
	for i := 0; i < length; i++ {
		if i >= len(S) || S[i] == ' ' {
			result += "%20"
		} else {
			result += string(S[i])
		}
	}

	return result
}

// 时间复杂度：O（n）
// 空间复杂度：O（n）
func replaceSpaces3(S string, length int) string {
	s1 := strings.TrimRight(S, " ")
	if s1 == " " {
		s1 = "%20"
	}

	s2 := strings.ReplaceAll(s1, " ", "%20")

	l := length - len(s1)
	for l > 0 {
		s2 += "%20"
		l--
	}

	return s2
}

func Test_replaceSpaces(t *testing.T) {
	tests := []struct {
		name string
		s    string
		len  int
		want string
	}{
		{
			name: "only spaces",
			s:    "  ",
			len:  3,
			want: "%20%20%20",
		},
		{
			name: "leading and trailing spaces",
			s:    "  a  ",
			len:  4,
			want: "%20%20a%20",
		},
		{
			name: "all empty",
			s:    "",
			len:  0,
			want: "",
		},
		{
			name: "case1",
			s:    "Mr John Smith ",
			len:  13,
			want: "Mr%20John%20Smith",
		},
		{
			name: "all spaces and too long",
			s:    "               ",
			len:  5,
			want: "%20%20%20%20%20",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceSpaces3(tt.s, tt.len); got != tt.want {
				t.Errorf("replaceSpaces() = %v, want %v", got, tt.want)
			}
		})
	}
}
