package cracking_the_coding_interview

import (
	"sort"
	"testing"
)

// 哈希表
// 时间复杂度：O（n）
// 空间复杂度：O（n）
func majorityElement(nums []int) int {
	m := make(map[int]int)
	for _, n := range nums {
		m[n]++
	}

	if len(m) == 1 {
		return nums[0]
	}

	for k, v := range m {
		if v > len(nums)/2 {
			return k
		}
	}

	return -1
}

// 穷举遍历
// 时间复杂度：O（n^2）
// 空间复杂度：O（1）
func majorityElement2(nums []int) int {
	for i := 0; i < len(nums); i++ {
		count := 1
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				count++
			}
		}

		if count > len(nums)/2 {
			return nums[i]
		}
	}

	return -1
}

/*
Boyer-Moore 投票算法通过维护一个候选元素和一个计数器来识别多数元素。算法的核心思想是：
	•	如果计数器为零，则选择当前元素为候选元素，计数器设为1。
	•	如果计数器不为零，则：
	•	如果当前元素与候选元素相同，计数器加1。
	•	如果当前元素与候选元素不同，计数器减1。
*/
// 时间复杂度：O（n）
// 空间复杂度：O（1）
func majorityElement3(nums []int) int {
	candidate, count := 0, 0
	for _, n := range nums {
		if count == 0 {
			candidate = n
			count = 1
		} else if n == candidate {
			count++
		} else {
			count--
		}
	}

	count2 := 0
	for _, n := range nums {
		if n == candidate {
			count2++
		}
	}

	if count2 > len(nums)/2 {
		return candidate
	}

	return -1
}

// 排序
// 时间复杂度：O（n log N）
// 空间复杂度：O（1）
func majorityElement4(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	sort.Ints(nums)
	candidate := nums[len(nums)/2]
	count := 0
	for _, n := range nums {
		if n == candidate {
			count++
		}
	}

	if count > len(nums)/2 {
		return candidate
	}

	return -1
}

func Test_majorityElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "empty",
			nums: []int{},
			want: -1,
		},
		{
			name: "no majority",
			nums: []int{1, 2, 3},
			want: -1,
		},
		{
			name: "only one element",
			nums: []int{1},
			want: 1,
		},
		{
			name: "one majority",
			nums: []int{1, 2, 1},
			want: 1,
		},
		{
			name: "one majority2",
			nums: []int{2, 2, 2, 3, 3, 4, 4},
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := majorityElement(tt.nums); got != tt.want {
				t.Errorf("majorityElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
