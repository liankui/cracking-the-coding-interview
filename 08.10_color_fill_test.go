package cracking_the_coding_interview

import (
	"reflect"
	"testing"
)

// 深度优先搜索算法
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if len(image) <= sr || len(image) <= sc {
		return image
	}

	originalColor := image[sr][sc]
	if originalColor == newColor {
		return image
	}

	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r >= len(image) || c < 0 || c >= len(image[0]) || image[r][c] != originalColor {
			return
		}

		image[r][c] = newColor
		dfs(r-1, c)
		dfs(r+1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}

	dfs(sr, sc)

	return image
}

// 广度优先搜索算法
func floodFill2(image [][]int, sr int, sc int, newColor int) [][]int {
	if len(image) <= sr || len(image) <= sc {
		return image
	}

	originalColor := image[sr][sc]
	if originalColor == newColor {
		return image
	}

	queue := [][]int{{sr, sc}}
	image[sr][sc] = newColor

	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		r, c := current[0], current[1]
		for _, dir := range directions {
			nr, nc := r+dir[0], c+dir[1]
			if nr >= 0 && nr < len(image) && nc >= 0 && nc < len(image[0]) && image[nr][nc] == originalColor {
				image[nr][nc] = newColor
				queue = append(queue, []int{nr, nc})
			}
		}
	}

	return image
}

func Test_floodFill(t *testing.T) {
	tests := []struct {
		name     string
		image    [][]int
		sr       int
		sc       int
		newColor int
		want     [][]int
	}{
		{
			name:     "floodFill",
			image:    [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}},
			sr:       1,
			sc:       1,
			newColor: 2,
			want:     [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := floodFill2(tt.image, tt.sr, tt.sc, tt.newColor); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("floodFill() = %v, want %v", got, tt.want)
			}
		})
	}
}
