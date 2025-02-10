package cracking_the_coding_interview

import (
	"strconv"
	"strings"
	"testing"
)

// 时间复杂度：O（n）
// 空间复杂度：O（n）
func compressString(S string) string {
	if len(S) == 0 {
		return ""
	}
	if len(S) == 1 {
		return S
	}

	b := strings.Builder{}
	count := 1
	for i := 0; i < len(S)-1; i++ {
		if count == 1 {
			b.WriteByte(S[i])
		}

		if S[i] == S[i+1] {
			count++
			if i == len(S)-2 {
				b.WriteString(strconv.Itoa(count))
			}
		} else {
			b.WriteString(strconv.Itoa(count))
			count = 1
			if i == len(S)-2 {
				b.WriteByte(S[i+1])
				b.WriteString(strconv.Itoa(count))
			}
		}
	}

	if b.Len() >= len(S) {
		return S
	} else {
		return b.String()
	}
}

func Test_compressString(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "empty string",
			s:    "",
			want: "",
		},
		{
			name: "aabbb",
			s:    "aabbb",
			want: "a2b3",
		},
		{
			name: "aabbbc",
			s:    "aabbbc",
			want: "a2b3c1",
		},
		{
			name: "aabcccccaaa",
			s:    "aabcccccaaa",
			want: "a2b1c5a3",
		},
		{
			name: "abbccd",
			s:    "abbccd",
			want: "abbccd",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compressString(tt.s); got != tt.want {
				t.Errorf("compressString() = %v, want %v", got, tt.want)
			}
		})
	}
}
