package cracking_the_coding_interview

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

var (
	tests = []struct {
		name  string
		array []int
		want  []int
	}{
		{
			name:  "case1",
			array: []int{12, 8, 38, 65, 2, 93},
			want:  []int{2, 8, 12, 38, 65, 93},
		},
		{
			name:  "case2",
			array: []int{64, 34, 25, 12, 22, 11, 90},
			want:  []int{11, 12, 22, 25, 34, 64, 90},
		},
		{
			name:  "case3",
			array: []int{1, 2, 3, 2, 5, 6},
			want:  []int{1, 2, 2, 3, 5, 6},
		},
	}
)

// SelectionSort 在每一轮中遍历未排序的部分，找到最大值，然后进行交换操作。(拿第i个数字和i后边的所有数字进行比较，小的排在前边)
// 时间复杂度：O（n^2）
// 空间复杂度：O（1）
func SelectionSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				fmt.Printf("swap %2d[%2d] and %2d[%2d], arr:%#v", arr[i], i, arr[j], j, arr)
				arr[i], arr[j] = arr[j], arr[i]
				fmt.Printf(" -> arr:%#v\n", arr)
			}
		}
	}

	return arr
}

func TestSort_SelectionSort(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := SelectionSort(test.array)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

// InsertionSort 将数组分为已排序和未排序两部分，每次从未排序部分取一个元素，插入到已排序部分的正确位置。
// 时间复杂度：O（n^2）
// 空间复杂度：O（1）
func InsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		tmp := arr[i]
		j := i - 1

		// 将大于 key 的元素向后移动
		for j >= 0 && arr[j] > tmp {
			arr[j+1] = arr[j]
			j--
		}

		// 插入元素
		arr[j+1] = tmp
		fmt.Printf(" -> arr:%#v\n", arr)
	}

	return arr
}

func TestSort_InsertionSort(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := InsertionSort(test.array)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

// BubbleSort 通过多次遍历，依次比较相邻的两个元素并进行交换，将较大的元素逐步“冒泡”到序列的末尾。
// 时间复杂度：O（n^2）
// 空间复杂度：O（1）
func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ { // 外层循环控制轮数
		swapped := false

		for j := 0; j < len(arr)-i-1; j++ { // 内层循环比较相邻元素
			if arr[j] > arr[j+1] {
				fmt.Printf("[%d] swap %2d[%2d] and %2d[%2d], arr:%#v", i, arr[j], j, arr[j+1], j+1, arr)
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
				fmt.Printf(" -> arr:%#v\n", arr)
			}
		}

		// 如果在某一轮比较中没有发生交换，说明数组已经是有序的，可以提前退出。
		if !swapped {
			break
		}
	}

	return arr
}

func TestSort_BubbleSort(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := BubbleSort(test.array)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

// QuickSort 快速排序是一种分治法的排序算法，主要思想是通过一个基准值（pivot）将数组划分为两个子数组，然后递归地对两个子数组进行排序
// 时间复杂度：O（n log n）
// 空间复杂度：O（log n）
func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// 1.选择基准值
	pivot := arr[0]

	// 2.分区
	var left, right []int
	for _, val := range arr[1:] {
		if val < pivot {
			left = append(left, val)
		} else {
			right = append(right, val)
		}
	}

	// 3.递归排序并合并
	left = QuickSort(left)
	right = QuickSort(right)

	return append(append(left, pivot), right...)
}

func TestSort_QuickSort(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := QuickSort(test.array)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

// QuickSort2 随机化基准值，避免最差情况（极度不平衡的分区）。
func QuickSort2(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// 1.选择基准值
	// 随机化基准值
	randIdx := rand.Intn(len(arr))
	arr[0], arr[randIdx] = arr[randIdx], arr[0]
	pivot := arr[0]

	// 2.分区
	var left, right []int
	for _, val := range arr[1:] {
		if val < pivot {
			left = append(left, val)
		} else {
			right = append(right, val)
		}
	}

	// 3.递归排序并合并
	left = QuickSort2(left)
	right = QuickSort2(right)

	return append(append(left, pivot), right...)
}

// QuickSort3 三路快排，提升性能，特别是处理大量重复元素的情况。
func QuickSort3(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// 1.选择基准值
	pivot := arr[0]

	// 2.分区
	var left, middle, right []int
	for _, val := range arr[1:] {
		if val < pivot {
			left = append(left, val)
		} else if val == pivot {
			middle = append(middle, val)
		} else if val > pivot {
			right = append(right, val)
		}
	}

	// 3.递归排序并合并
	left = QuickSort3(left)
	right = QuickSort3(right)

	return append(append(left, middle...), right...)
}
