package cracking_the_coding_interview

import (
	"reflect"
	"sort"
	"testing"
)

func merge(A []int, m int, B []int, n int) {
	A = A[:m]
	B = B[:n]
	A = append(A, B...)
	sort.Ints(A)
}

func merge2(A []int, m int, B []int, n int) {
	A = A[:m]
	B = B[:n]
	A = append(A, B...)
	SelectionSort(A)
}

func Test_merge(t *testing.T) {
	tests := []struct {
		name string
		A    []int
		m    int
		B    []int
		n    int
		want []int
	}{
		{
			name: "default",
			A:    []int{1, 2, 3, 0, 0, 0},
			m:    3,
			B:    []int{2, 5, 6},
			n:    3,
			want: []int{1, 2, 2, 3, 5, 6},
		},
		{
			name: "A is empty",
			A:    []int{0, 0, 0},
			m:    0,
			B:    []int{2, 5, 6},
			n:    3,
			want: []int{2, 5, 6},
		},
		// {
		// 	name: "B is empty",
		// 	A:    []int{1, 2, 3, 0, 0, 0},
		// 	m:    3,
		// 	B:    []int{},
		// 	n:    0,
		// 	want: []int{1, 2, 3},
		// },
		// {
		// 	name: "A and B is empty",
		// 	A:    []int{0, 0, 0},
		// 	m:    0,
		// 	B:    []int{},
		// 	n:    0,
		// 	want: []int{},
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge2(tt.A, tt.m, tt.B, tt.n)
			if !reflect.DeepEqual(tt.A, tt.want) {
				t.Errorf("merge() = %v, want %v", tt.A, tt.want)
			}
		})
	}
}
