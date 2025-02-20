package cracking_the_coding_interview

import (
	"strings"
	"testing"
)

func respace(dictionary []string, sentence string) int {
	diff := 0
	for _, word := range dictionary {
		for _, char := range word {
			diff ^= int(char - 'a')
		}
	}

	for _, char := range sentence {
		diff ^= int(char - 'a')
	}

	count := 0
	for diff != 0 {
		if diff&1 == 1 {
			count++
		}
		diff >>= 1
	}

	return count
}

func respace2(dictionary []string, sentence string) int {
	l := len(sentence)
	if len(dictionary) == 0 {
		return l
	}

	dl := len(dictionary) - 1
	count := 0

	for i := l - 1; i >= 0; {
		if strings.HasSuffix(sentence, dictionary[dl]) {
			wordLen := len(dictionary[dl])
			i -= wordLen
			sentence = sentence[:len(sentence)-wordLen]
			dl--
		} else {
			i--
			count++
		}

		if dl < 0 {
			count += len(sentence)
			break
		}
	}

	return count
}

// TODO
func Test_respace(t *testing.T) {
	tests := []struct {
		name string
		dict []string
		sent string
		want int
	}{
		{
			name: "normal",
			dict: []string{"a", "c"},
			sent: "abac",
			want: 2,
		},
		{
			name: "normal",
			dict: []string{"a", "d"},
			sent: "abcd",
			want: 2,
		},
		{
			name: "case",
			dict: []string{"looked", "just", "like", "her", "brother"},
			sent: "jesslookedjustliketimherbrother",
			want: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := respace2(tt.dict, tt.sent); got != tt.want {
				t.Errorf("respace() = %v, want %v", got, tt.want)
			}
		})
	}
}
