package day10

import (
	"bufio"
	"io"
)

func Part1(input io.Reader) (int, error) {
	matrix := [][]int{}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		row := []int{}

		for _, e := range scanner.Text() {
			row = append(row, int(e-'0'))
		}

		matrix = append(matrix, row)
	}

	directions := [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
	isValid := func(row, col int) bool {
		return 0 <= row && row < len(matrix) && 0 <= col && col < len(matrix[0])
	}

	ans := 0

	var dfs func(int, int, map[*int]bool)
	dfs = func(row, col int, seen map[*int]bool) {
		if matrix[row][col] == 9 {
			if !seen[&matrix[row][col]] {
				seen[&matrix[row][col]] = true
				ans++
			}

			return
		}

		for _, dir := range directions {
			nextRow, nextCol := row+dir[0], col+dir[1]

			if isValid(nextRow, nextCol) && matrix[nextRow][nextCol] == matrix[row][col]+1 {
				dfs(nextRow, nextCol, seen)
			}
		}
	}

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			if matrix[row][col] == 0 {
				dfs(row, col, make(map[*int]bool))
			}
		}
	}

	return ans, nil
}


func Part2(input io.Reader) (int, error) {
	matrix := [][]int{}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		row := []int{}

		for _, e := range scanner.Text() {
			row = append(row, int(e-'0'))
		}

		matrix = append(matrix, row)
	}

	directions := [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
	isValid := func(row, col int) bool {
		return 0 <= row && row < len(matrix) && 0 <= col && col < len(matrix[0])
	}

	ans := 0

	var dfs func(int, int)
	dfs = func(row, col int) {
		if matrix[row][col] == 9 {
			ans++
			return
		}

		for _, dir := range directions {
			nextRow, nextCol := row+dir[0], col+dir[1]

			if isValid(nextRow, nextCol) && matrix[nextRow][nextCol] == matrix[row][col]+1 {
				dfs(nextRow, nextCol)
			}
		}
	}

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			if matrix[row][col] == 0 {
				dfs(row, col)
			}
		}
	}

	return ans, nil
}
