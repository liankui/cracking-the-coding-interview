package cracking_the_coding_interview

import "testing"

// 时间复杂度：O（n）
// 空间复杂度：O（1）
func findMagicIndex(nums []int) int {
	for i, num := range nums {
		if i == num {
			return i
		}
	}
	return -1
}

// 时间复杂度：O（log n）
// 空间复杂度：O（log n ）
func findMagicIndex2(nums []int) int {
	return findMagicIndexHelper(nums, 0, len(nums)-1)
}

func findMagicIndexHelper(nums []int, left, right int) int {
	if left > right {
		return -1
	}

	mid := left + (right-left)/2

	leftIndex := findMagicIndexHelper(nums, left, mid-1)
	if leftIndex != -1 {
		return leftIndex
	} else if nums[mid] == mid {
		return mid
	} else {
		return findMagicIndexHelper(nums, mid+1, right)
	}
}

func Test_findMagicIndex(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "case1",
			nums: []int{},
			want: -1,
		},
		{
			name: "case2",
			nums: []int{0},
			want: 0,
		},
		{
			name: "case3",
			nums: []int{1, 1, 2},
			want: 1,
		},
		{
			name: "case4",
			nums: []int{1, 2},
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMagicIndex2(tt.nums); got != tt.want {
				t.Errorf("findMagicIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
