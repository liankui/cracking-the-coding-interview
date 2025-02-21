package cracking_the_coding_interview

import (
	"reflect"
	"testing"
)

// 暴力法
// 时间复杂度：O（n^2）
// 空间复杂度：O（n）
func nthPrimeNumber(n int) []int {
	nums := make([]int, 0, n)

	for i := 2; i <= n; i++ {
		prime := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				prime = false
			}
		}

		if prime {
			nums = append(nums, i)
		}
	}

	return nums
}

// 暴力法优化 (只检查前sqrt n的元素）
// 时间复杂度：O（n log n）
// 空间复杂度：O（n）
func nthPrimeNumber2(n int) []int {
	nums := make([]int, 0, n)

	for i := 2; i <= n; i++ {
		prime := true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				prime = false
				break
			}
		}

		if prime {
			nums = append(nums, i)
		}
	}

	return nums
}

// 全部标记法，将小于n的一个数的所有倍数进行标记，没被标记则为质数
// 时间复杂度：O（n log n）
// 空间复杂度：O（n）
func nthPrimeNumber3(n int) []int {
	nums := make([]int, 0, n)

	primes := make([]bool, n+2)
	for i := 2; i <= n; i++ {
		primes[i] = true
	}

	for i := 2; i*i <= n; i++ {
		if primes[i] {
			for j := i * i; j <= n; j += i {
				primes[j] = false
			}
		}
	}

	for i := 2; i <= n; i++ {
		if primes[i] {
			nums = append(nums, i)
		}
	}

	return nums
}

// 暴力法优化, 使用已知质数来判断
// 时间复杂度：O（n log n）
// 空间复杂度：O（n）
func nthPrimeNumber4(n int) []int {
	primes := make([]int, 0, n)
	for i := 2; i <= n; i++ {
		if isPrime(i, primes) {
			primes = append(primes, i)
		}
	}

	return primes
}

func isPrime(num int, primes []int) bool {
	for _, p := range primes {
		if p*p > num {
			break
		}

		if num%p == 0 {
			return false
		}
	}

	return true
}

func Test_nthPrimeNumber(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []int
	}{
		{
			name: "5",
			n:    5,
			want: []int{2, 3, 5},
		},
		{
			name: "11",
			n:    11,
			want: []int{2, 3, 5, 7, 11},
		},
		{
			name: "100",
			n:    100,
			want: []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nthPrimeNumber4(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nthPrimeNumber() = %+v, want %v", got, tt.want)
			}
		})
	}
}
