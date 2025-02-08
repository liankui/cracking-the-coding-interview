package cracking_the_coding_interview

import "testing"

// 时间复杂度：O（n）
// 空间复杂度：O（n）
func missingNumber(nums []int) int {
	m := make(map[int]struct{}, len(nums))
	for _, v := range nums {
		m[v] = struct{}{}
	}

	for i := 0; i < len(nums); i++ {
		if _, ok := m[i]; !ok {
			return i
		}
	}

	return len(nums)
}

// 时间复杂度：O（n log n）
// 空间复杂度：O（log n）
func missingNumber2(nums []int) int {
	newNums := quickSortMissingNumber(nums)
	for i := 0; i < len(newNums); i++ {
		if i != newNums[i] {
			return i
		}
	}

	return len(newNums)
}

func quickSortMissingNumber(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}

	qovit := nums[0]
	var left, right []int
	for _, num := range nums[1:] {
		if num < qovit {
			left = append(left, num)
		} else {
			right = append(right, num)
		}
	}

	left = quickSortMissingNumber(left)
	right = quickSortMissingNumber(right)

	return append(append(left, qovit), right...)
}

// 时间复杂度：O（n）
// 空间复杂度：O（1）
func missingNumber3(nums []int) int {
	l := len(nums)
	expectedSum := l * (l + 1) / 2
	actualSum := 0
	for _, num := range nums {
		actualSum += num
	}
	return expectedSum - actualSum
}

// 利用 a^a = 0 和 a^0 = a 的性质。  将 0到n 与数组中的所有数字异或，相同的数字异或为 0，最后的结果就是 missing number
// 时间复杂度：O（n）
// 空间复杂度：O（1）
func missingNumber4(nums []int) int {
	missing := len(nums)
	for i, num := range nums {
		missing ^= i ^ num
	}
	return missing
}

// 时间复杂度：O（n）
// 空间复杂度：O（1）
func missingNumber5(nums []int) int {
	l := len(nums)
	for i := 0; i < l; i++ { // O（n^2） or O（n）
		for nums[i] != i && nums[i] < l {
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}

	for i := 0; i < l; i++ {
		if nums[i] != i {
			return i
		}
	}

	return l
}

func Test_missingNumber(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "empty",
			nums: []int{},
			want: 0,
		},
		{
			name: "miss 1",
			nums: []int{0},
			want: 1,
		},
		{
			name: "miss 2",
			nums: []int{3, 0, 1},
			want: 2,
		},
		{
			name: "miss 8",
			nums: []int{9, 6, 4, 2, 3, 5, 7, 0, 1},
			want: 8,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := missingNumber5(test.nums); got != test.want {
				t.Errorf("%v: got %v, want %v", test.name, got, test.want)
			}
		})
	}
}
