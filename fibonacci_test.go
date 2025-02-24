package cracking_the_coding_interview

import "testing"

// 递归
// 时间复杂度：O（2^n）
// 空间复杂度：O（n）
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// 动态规划法（自底向上）
// 时间复杂度：O（n）
// 空间复杂度：O（n）
func fibonacci2(n int) int {
	if n <= 1 {
		return n
	}

	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

// 动态规划法（优化空间）
// 时间复杂度：O（n）
// 空间复杂度：O（1）
func fibonacci3(n int) int {
	if n <= 1 {
		return n
	}

	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}

	return b
}

// 矩阵快速幂
/*
通过快速幂方法在 O(log n) 时间复杂度内计算斐波那契数列的第 n 项。原理基于斐波那契数列的矩阵形式：
\begin{bmatrix}
F(n+1) & F(n) \\
F(n) & F(n-1)
\end{bmatrix}

\begin{bmatrix}
1 & 1 \\
1 & 0
\end{bmatrix}^n
*/

// 时间复杂度：O（log n）
// 空间复杂度：O（1）
func fibonacci4(n int) int {
	if n == 0 {
		return 0
	}

	matrix := [2][2]int{{1, 1}, {1, 0}}
	result := matrixPower(matrix, n-1)
	return result[0][0]
}

func matrixMultiply(a, b [2][2]int) [2][2]int {
	return [2][2]int{
		{a[0][0]*b[0][0] + a[0][1]*b[1][0], a[0][0]*b[0][1] + a[0][1]*b[1][1]},
		{a[1][0]*b[0][0] + a[1][1]*b[1][0], a[1][0]*b[0][1] + a[1][1]*b[1][1]},
	}
}

func matrixPower(matrix [2][2]int, n int) [2][2]int {
	result := [2][2]int{{1, 0}, {0, 1}} // 单位矩阵
	base := matrix

	for n > 0 {
		if n%2 == 1 {
			result = matrixMultiply(result, base)
		}
		base = matrixMultiply(base, base)
		n /= 2
	}

	return result
}

func Test_fibonacci(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{n: 0, want: 0},
		{n: 1, want: 1},
		{n: 2, want: 1},
		{n: 3, want: 2},
		{n: 4, want: 3},
		{n: 5, want: 5},
		{n: 6, want: 8},
		{n: 7, want: 13},
	}

	for _, tt := range tests {
		if got := fibonacci4(tt.n); got != tt.want {
			t.Errorf("fibonacci(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}
